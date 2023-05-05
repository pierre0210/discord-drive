package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pierre0210/discord-drive/internal/handler"
	"github.com/pierre0210/discord-drive/internal/middleware"
	"github.com/pierre0210/discord-drive/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Fail to load .env file.")
	}
	storage.InitTable()

	token := os.Getenv("TOKEN")
	port := 5000

	bot, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatalln(err.Error())
	}
	router := gin.Default()
	router.LoadHTMLGlob("view/*")
	router.GET("/", handler.GetIndex)
	router.GET("/files", handler.GetFileList)
	router.GET("/download", middleware.AddSession(bot), handler.GetFile)
	router.POST("/upload", middleware.AddSession(bot), handler.PostUpload)

	err = bot.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer bot.Close()
	log.Println("Bot logged in.")
	router.Run(fmt.Sprintf(":%d", port))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
