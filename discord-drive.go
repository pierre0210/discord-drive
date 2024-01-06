package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/pierre0210/discord-drive/internal/command"
	"github.com/pierre0210/discord-drive/internal/server"
	"github.com/pierre0210/discord-drive/internal/storage"
)

func main() {
	var cmdMode bool
	var listFileName bool
	var filePath string
	var fileName string
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Fail to load .env file.")
	}
	storage.InitTable()

	token := os.Getenv("TOKEN")
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	flag.BoolVar(&cmdMode, "c", false, "command line mode")
	flag.BoolVar(&listFileName, "l", false, "list all files")
	flag.StringVar(&filePath, "u", "", "upload file")
	flag.StringVar(&fileName, "d", "", "download file")
	flag.Parse()

	bot, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = bot.Open()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer bot.Close()
	log.Println("Bot logged in.")

	if cmdMode {
		if filePath != "" {
			command.Upload(bot, filePath)
		} else if fileName != "" {
			command.Download(bot, fileName)
		}
	} else {
		server.Init(bot, port)
	}
}
