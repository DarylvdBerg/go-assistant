package client

import (
	"encoding/json"
	"go-assistant/internal/initialize"
	"go-assistant/internal/shared"
	"os"

	"github.com/charmbracelet/log"
)

func LoadConfig() (*shared.Config, error) {
	initialize.CreateConfigIfNotExists()
	
	file, err := os.Open(initialize.GetConfigPath())
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}(file)

	var config shared.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
