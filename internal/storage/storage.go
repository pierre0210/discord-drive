package storage

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var DBPath string = "./db.json"

type FileTable struct {
	Files   map[string]string `json:"files"`
	IdChain map[string]string `json:"chain"`
}

func (f *FileTable) AddFile(names string, id string) {
	f.Files[names] = id
}

func (f *FileTable) AddToChain(parentId string, childId string) {
	f.IdChain[parentId] = childId
}

func InitTable() {
	if _, err := os.Stat(DBPath); errors.Is(err, os.ErrNotExist) {
		table := new(FileTable)
		table.Files = make(map[string]string)
		table.IdChain = make(map[string]string)

		file, err := os.Create(DBPath)
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer file.Close()

		jsonBytes, _ := json.Marshal(table)
		file.Write(jsonBytes)
	}
}
