package cmd

import "github.com/spf13/cobra"

var lightsOnCmd = &cobra.Command {
	Use: "turn-on",
	Short: "Turn on the light in Home Assistant",
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		entityId := args[0]
		GetClient().TurnOnLight(entityId);
		return nil;
	},
}

func init() {
	rootCmd.AddCommand(lightsOnCmd)
}

func LightsOnWorkLoad(args []string) {
	entityId := args[0]
	GetClient().TurnOnLight(entityId);
} 