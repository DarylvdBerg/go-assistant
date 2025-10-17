package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "configure",
	Short: "Edit the configuration for go-assistant",
	Run: func(command *cobra.Command, args []string) {
		configDir, err := os.UserConfigDir()
		if err != nil {
			log.Fatalf("failed to get user config dir: %s", err)
		}

		configPath := filepath.Join(configDir, "go-assistant", "config.json")

		// Ensure the config file exists (optional)
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			fmt.Printf("Config file does not exist at %s\n", configPath)
			return
		}

		nanoPath, err := exec.LookPath("nano")
		if err != nil {
			log.Fatalf("nano not found in PATH: %s", err)
		}

		cmd := exec.Command(nanoPath, configPath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("Unable to open config file: %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
