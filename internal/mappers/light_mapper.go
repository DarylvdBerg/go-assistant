package mappers

import (
	"log"

	"github.com/DarylvdBerg/go-assistant/shared/supported_color_modes"
)

// MapSupportedColorModes maps a list of supported color modes from any type to the SupportedColorModes enum
func MapSupportedColorModes(entries []any) []supported_color_modes.SupportedColorModes {
	modes := make([]supported_color_modes.SupportedColorModes, 0)
	for _, value := range entries {
		mode, ok := value.(string)

		if !ok {
			log.Println("Invalid type assertion for mode from supported color modes.")
			continue
		}

		colorMode, err := supported_color_modes.EnumValue(mode)
		if err != nil {
			log.Println("Unable to parse Enum value, it is either invalid or not supported")
			continue
		}

		modes = append(modes, colorMode)
	}

	return modes
}

// MapBrightness maps a brightness value from 0-255 to a percentage value from 0-100
func MapBrightness(brightnessValue any) int {
	switch v := brightnessValue.(type) {
	case float64:
		return int(v / 255 * 100)
	case int:
		return int(float64(v) / 255 * 100)
	default:
		return 0
	}
}
