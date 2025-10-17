package cmd

import (
	"fmt"
	"go-assistant/internal/initialize"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "🔧 Initialize the configuration file in the user configuration folder.",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.CreateConfigIfNotExists()
		fmt.Println("Configuration file initialized.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
