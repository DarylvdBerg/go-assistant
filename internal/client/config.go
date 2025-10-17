package client

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	BaseUrl string `json:"base_url"`
	Token   string `json:"token"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}(file)

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
