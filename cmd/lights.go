package cmd

import (
	"fmt"
	"go-assistant/internal/homeassistant"
	"go-assistant/ui/lights"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var lightsCmd = &cobra.Command{
	Use:   "lights",
	Short: "💡 Control your home-assistant configured lights. You'll be able to turn them on/off and adjust the brightness.",
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
