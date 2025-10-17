package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "go-assistant [pattern]",
	Short:   "🚀 Control your Home Assistant from the terminal 🚀",
	Example: "go-assistant lights",
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true   // Remove the default completion command since we're not using it.
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true}) // Remove the default help command since we're not using it.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
