package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"workspace/practice/tcp-dialer/model"
)

const configFileName = "config.json"

// GetConfig build application config form file on disk
func GetConfig() *model.Config {
	file, fileError := os.Open(fmt.Sprintf("config/%s", configFileName))
	defer file.Close()

	if fileError != nil {
		log.Println("GetConfig error occurred while reading config file", fileError)
		panic(fileError)
	}

	decoder := json.NewDecoder(file)
	config := model.Config{}
	decodeError := decoder.Decode(&config)

	if decodeError != nil {
		log.Println("GetConfig error occurred while reading config file", fileError)
		panic(fileError)
	}

	return &config
}
