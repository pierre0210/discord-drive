package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pierre0210/discord-drive/internal/storage"
)

func GetFiles(ctx *gin.Context) {
	var table storage.FileTable
	tableBytes, _ := os.ReadFile(storage.DBPath)
	json.Unmarshal(tableBytes, &table)
	ctx.JSON(http.StatusOK, table.Files)
}
