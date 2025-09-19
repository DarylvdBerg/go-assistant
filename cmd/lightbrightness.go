package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

var lightChangeBrightnessCmd = &cobra.Command {
	Use: "brightness",
	Short: "Change the brightness by passing the entity id and brightness percentage",
	Args: cobra.ExactArgs(2),
	RunE: func(command *cobra.Command, args []string) error {
		entityId := args[0]
		brightnessUint, err := strconv.ParseUint(args[1], 10, 8)

		if err != nil {
			log.Fatal("Unable to convert to unint")
			return nil;
		}

		brightness := uint8(brightnessUint)

		if brightness > 100 {
			log.Fatal("brightness should be between 0 and 100")
			return nil;
		}

		GetClient().ChangeBrightness(entityId, brightness)
		return nil;
	},
}

func init() {
	rootCmd.AddCommand(lightChangeBrightnessCmd)
}