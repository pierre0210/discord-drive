package discordutil

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func DownloadFileFromChannel(bot *discordgo.Session, messageId string) ([]byte, error) {
	var fileBuff bytes.Buffer
	msg, err := bot.ChannelMessage(os.Getenv("CHANNELID"), messageId)
	if err != nil {
		return []byte{}, err
	} else if len(msg.Attachments) != 1 {
		return []byte{}, errors.New("wrong message")
	}
	fileURL := msg.Attachments[0].URL
	res, _ := http.Get(fileURL)
	if res.StatusCode != http.StatusOK {
		return []byte{}, errors.New("request failed")
	}
	defer res.Body.Close()
	io.Copy(&fileBuff, res.Body)

	return fileBuff.Bytes(), nil
}
