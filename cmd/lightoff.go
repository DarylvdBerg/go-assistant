package cmd

import "github.com/spf13/cobra"

var lightsOffCmd = &cobra.Command {
	Use: "turn-off",
	Short: "Turn off light",
	Args: cobra.ExactArgs(1),
	RunE: func(command *cobra.Command, args []string) error{
		entityId := args[0]
		GetClient().TurnOffLight(entityId)
		return nil;
	},
}

func init() {
	rootCmd.AddCommand(lightsOffCmd)
}
