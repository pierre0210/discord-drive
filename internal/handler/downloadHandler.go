package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/pierre0210/discord-drive/internal/discordutil"
	"github.com/pierre0210/discord-drive/internal/storage"
)

func GetFile(ctx *gin.Context) {
	bot, _ := ctx.MustGet("bot").(*discordgo.Session)
	fileName := ctx.Query("file")
	if fileName == "" {
		ctx.Status(http.StatusBadRequest)
		return
	}
	var table storage.FileTable
	dbBytes, _ := os.ReadFile(storage.DBPath)
	json.Unmarshal(dbBytes, &table)
	indexId, ok := table.Files[fileName]
	if !ok {
		ctx.Status(http.StatusNotFound)
		return
	}

	header := ctx.Writer.Header()
	header.Set("Transfer-Encoding", "chunked")
	header.Set("Content-Type", "application/octet-stream")
	header.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	ctx.Writer.WriteHeader(http.StatusOK)

	for {
		fileBytes, err := discordutil.DownloadFileFromChannel(bot, indexId)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return
		}
		ctx.Writer.Write(fileBytes)
		ctx.Writer.(http.Flusher).Flush()
		indexId, ok = table.IdChain[indexId]
		if !ok {
			break
		}
	}
}
