package discordutil

import (
	"bytes"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func UploadFileToChannel(bot *discordgo.Session, fileName string, fileBuff *bytes.Buffer) *discordgo.Message {
	message, err := bot.ChannelFileSend(os.Getenv("CHANNELID"), fileName, fileBuff)
	if err != nil {
		log.Println(err.Error())
	}

	return message
}
