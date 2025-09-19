package main

import (
	"go-assistant-cli/cmd"
	"go-assistant-cli/internal/homeassistant"
	"log"
)

func main() {

	// Setup the client
	config, err := homeassistant.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %w", err)
	}
	client := homeassistant.CreateNewClient(config.BaseUrl, config.Token)
	cmd.SetClient(client);
	
	// Execute
	cmd.Execute()
}