package server

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/pierre0210/discord-drive/internal/handler"
	"github.com/pierre0210/discord-drive/internal/middleware"
)

func Init(bot *discordgo.Session, port int) {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.GET("/", handler.GetIndex)
	router.GET("/files", handler.GetFileList)
	router.GET("/download", middleware.AddSession(bot), handler.GetFile)
	router.POST("/upload", middleware.AddSession(bot), handler.PostUpload)
	router.Run(fmt.Sprintf(":%d", port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
