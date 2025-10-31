package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func CreateConfigIfNotExists() {
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("failed to get user config dir: %s", err)
	}

	applicationConfigPath := fmt.Sprintf("%s/go-assistant", userConfigPath)

	// Config directory does not exist yet, so we'll create it.
	if _, err := os.Stat(applicationConfigPath); os.IsNotExist(err) {
		log.Debug("config directory does not exist, creating it...")
		err = os.MkdirAll(applicationConfigPath, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to create config directory: %s", err)
		}
	}

	// Now validate if we have a config file, otherwise create it.
	configFilePath := fmt.Sprintf("%s/config.json", applicationConfigPath)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		log.Debug("config file does not exist, creating it...")

		_, err := os.Create(configFilePath)
		if err != nil {
			log.Fatalf("failed to create config file: %s", err)
		}

		log.Debug("Creating empty config object...")
		emptyConfigData, err := json.MarshalIndent(Config{}, "", "  ")

		if err != nil {
			log.Fatalf("failed to marshal empty config data: %s", err)
		}

		log.Debug("Writing empty config object to file...")
		err = os.WriteFile(configFilePath, emptyConfigData, os.ModePerm)
		if err != nil {
			log.Fatalf("failed to write empty config data: %s", err)
		}
	}
}
