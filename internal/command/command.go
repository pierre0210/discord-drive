package command

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/bwmarrin/discordgo"
	"github.com/pierre0210/discord-drive/internal/discordutil"
	"github.com/pierre0210/discord-drive/internal/storage"
)

func Upload(bot *discordgo.Session, filePath string) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		var table storage.FileTable
		var prevId string
		content, _ := os.Open(filePath)
		chunkSize, _ := strconv.Atoi(os.Getenv("CHUNKSIZE"))
		reader := bufio.NewReaderSize(content, chunkSize)
		info, _ := content.Stat()
		size := info.Size()
		for i := 0; i < int(math.Ceil(float64(size)/float64(chunkSize))); i++ {
			splitBuff := make([]byte, chunkSize)
			size, _ := reader.Read(splitBuff)
			log.Printf("%d bytes", size)
			message := discordutil.UploadFileToChannel(bot, "chunk", bytes.NewBuffer(splitBuff[:size]))

			if i == 0 {
				table.AddFile(info.Name(), message.ID)
				prevId = message.ID
			} else {
				table.AddToChain(prevId, message.ID)
				prevId = message.ID
			}
		}
		content.Close()
	} else {
		log.Println("file not exist")
	}
}

func Download(bot *discordgo.Session, fileName string) {
	var table storage.FileTable
	dbBytes, _ := os.ReadFile(storage.DBPath)
	json.Unmarshal(dbBytes, &table)
	indexId, ok := table.Files[fileName]
	if !ok {
		log.Println("file not exist")
		return
	}

	writer, err := os.Create(fileName)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for {
		fileBytes, err := discordutil.DownloadFileFromChannel(bot, indexId)
		if err != nil {
			log.Println("chunk not found")
		}
		_, err = writer.Write(fileBytes)
		if err != nil {
			log.Println(err.Error())
		}
		indexId, ok = table.IdChain[indexId]
		if !ok {
			break
		}
	}
	writer.Close()
}
