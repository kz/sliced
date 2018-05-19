package database

import (
	"encoding/json"
	"log"
	"os"
)

type file struct {
	filename string
}

func NewFileStore(filename string) *file {
	f, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("Init file failed: ", err.Error())
	}

	f.Close()

	return &file{
		filename: filename,
	}
}

func (file *file) AddItem(item map[string]interface{}) {
	f, err := os.OpenFile(file.filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Open file failed: ", err.Error())
	}

	defer f.Close()

	text, err := json.Marshal(item)
	if err != nil {
		log.Fatal("JSON marshal failed: ", err.Error())
	}

	_, err = f.WriteString(string(text) + "\n")
	if err != nil {
		log.Fatal("Write string failed: ", err.Error())
	}
}
