package main

import (
	"go-assistant/cmd"
	"go-assistant/internal/client"
)

func main() {
	client.NewClient()
	cmd.Execute()
}
