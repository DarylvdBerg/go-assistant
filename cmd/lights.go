package cmd

import (
	"fmt"
	"go-assistant-cli/internal/homeassistant"
	"go-assistant-cli/ui/lights"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var lightsCmd = &cobra.Command{
	Use:   "lights",
	Short: "Get a list of configured lights in your Home Assistant",
	Run: func(command *cobra.Command, args []string) {
		lightsWorkload()
	},
}

func init() {
	rootCmd.AddCommand(lightsCmd)
}

func lightsWorkload() {
	lightsData, err := homeassistant.GetClient().ListLights()
	if err != nil {
		fmt.Println("error fetch configured lights: ", err)
		return
	}

	p := tea.NewProgram(lights.InitLightOverview(lightsData), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Print("Error running program:", err)
		os.Exit(1)
	}
}
