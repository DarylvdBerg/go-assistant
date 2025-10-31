package client

import (
	"encoding/json"
	"os"

	"github.com/DarylvdBerg/go-assistant/internal/config"
	"github.com/charmbracelet/log"
)

func LoadConfig() (*config.Config, error) {
	config.CreateConfigIfNotExists()

	file, err := os.Open(config.GetConfigPath())
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}(file)

	var config config.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
