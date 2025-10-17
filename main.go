package main

import (
	"go-assistant-cli/cmd"
	"go-assistant-cli/internal/client"
)

func main() {
	client.NewClient()
	cmd.Execute()
}
