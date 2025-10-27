package initialize

import (
	"fmt"
	"log"
	"os"
)

func GetConfigPath() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Unable to get user config dir: %s", err)
	}

	applicationConfigPath := fmt.Sprintf("%s/github.com/DarylvdBerg/go-assistant/config.json", userConfigDir)

	return applicationConfigPath
}
