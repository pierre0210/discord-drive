package handler

import (
	"bytes"
	"io"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

func PostUpload(ctx *gin.Context) {
	bot, _ := ctx.MustGet("bot").(*discordgo.Session)
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println(err.Error())
	}
	files := form.File["files"]

	for _, file := range files {
		var buff bytes.Buffer
		content, _ := file.Open()
		io.Copy(&buff, content)
	}
}
