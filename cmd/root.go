package cmd

import (
	"fmt"
	"go-assistant-cli/internal/homeassistant"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "go-assistant",
	Short: "Control Home Assistant from the CLI",
}

var client *homeassistant.Client;

func SetClient(c *homeassistant.Client) {
	client = c;
}

func GetClient() *homeassistant.Client {
	return client;
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}