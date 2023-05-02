package handler

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/pierre0210/discord-drive/internal/discordutil"
	"github.com/pierre0210/discord-drive/internal/storage"
)

func PostUpload(ctx *gin.Context) {
	var table storage.FileTable
	tableBytes, _ := os.ReadFile(storage.DBPath)
	json.Unmarshal(tableBytes, &table)
	bot, _ := ctx.MustGet("bot").(*discordgo.Session)
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println(err.Error())
	}
	files := form.File["files"]

	for _, file := range files {
		var prevId string
		chunkSize, _ := strconv.Atoi(os.Getenv("CHUNKSIZE"))
		content, _ := file.Open()

		reader := bufio.NewReaderSize(content, chunkSize)
		for i := 0; i < int(math.Ceil(float64(file.Size)/float64(chunkSize))); i++ {
			splitBuff := make([]byte, chunkSize)
			size, _ := reader.Read(splitBuff)
			chunkSum := fmt.Sprintf("%x", md5.Sum(splitBuff[:size]))
			log.Printf("%d %s", size, chunkSum)
			message := discordutil.UploadFileToChannel(bot, chunkSum, bytes.NewBuffer(splitBuff[:size]))

			if i == 0 {
				table.AddFile(file.Filename, message.ID)
				prevId = message.ID
			} else {
				table.AddToChain(prevId, message.ID)
				prevId = message.ID
			}
		}
		content.Close()
	}
	tableBytes, _ = json.Marshal(table)
	os.WriteFile(storage.DBPath, tableBytes, 0666)
}
