package cmd

import (
	"fmt"
	"go-assistant-cli/ui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var lightsCmd = &cobra.Command {
	Use: "lights",
	Short: "Get a list of configured lights in your Home Assistant",
	Run: func(command *cobra.Command, args []string) {
		lightsWorkload()
	},
}

func init() {
	rootCmd.AddCommand(lightsCmd)
}

func lightsWorkload() {
	lights, err := GetClient().ListLights()
	if err != nil {
		fmt.Println("error fetch configured lights: ", err)
		return
	}
	
	p := tea.NewProgram(ui.InitData(lights), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program:", err)
		os.Exit(1)
	}
}