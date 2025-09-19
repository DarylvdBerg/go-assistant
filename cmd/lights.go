package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lightsCmd = &cobra.Command {
	Use: "lights",
	Short: "Get a list of configured lights in your Home Assistant",
	Run: func(command *cobra.Command, args []string) {
		lights, err := GetClient().ListLights()
		if err != nil {
			fmt.Println("error fetch configured lights: ", err)
			return
		}
		for _, light := range lights {
			fmt.Printf("%s (%s): %s\n", light.FriendlyName, light.EntityID, light.State)
		}
	},
}

func init() {
	rootCmd.AddCommand(lightsCmd)
}