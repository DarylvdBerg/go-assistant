package client

import (
	"encoding/json"
	"os"

	"github.com/DarylvdBerg/go-assistant/internal/initialize"
	"github.com/DarylvdBerg/go-assistant/internal/shared"

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
