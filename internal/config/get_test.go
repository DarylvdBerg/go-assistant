package config

import (
	"fmt"
	"os"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	userConfigPath, err := os.UserConfigDir()
	if err != nil {
		t.Fatalf("Failed to get user config directory: %v", err)
	}

	configPath := fmt.Sprintf("%s/go-assistant/config.json", userConfigPath)

	if GetConfigPath() != configPath {
		t.Errorf("GetConfigPath returned wrong config path")
	}
}
