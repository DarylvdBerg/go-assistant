package mappers

import (
	"log"
	"slices"

	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/models"
	"github.com/DarylvdBerg/go-assistant/shared/supported_color_modes"
)

const (
	LightsPart = "light."
)

func MapToLight(entity map[string]any) *models.Light {
	id, ok := entity["entity_id"].(string)
	if !ok || len(id) <= 6 || id[:6] != LightsPart {
		return nil
	}

	stateAttr, ok := entity["state"].(string)
	if !ok {
		return nil
	}

	state := light_state.EnumValue(stateAttr)

	attrs, ok := entity["attributes"].(map[string]any)
	if !ok {
		return nil
	}

	name, ok := attrs["friendly_name"].(string)
	if !ok {
		return nil
	}

	light := &models.Light{
		EntityID:     id,
		State:        state,
		FriendlyName: name,
	}

	if supportedModes, ok := attrs["supported_color_modes"].([]any); ok {
		light.SupportedColorModes = mapSupportedColorModes(supportedModes)
	}

	if brightness, ok := attrs["brightness"]; ok {
		light.Brightness = mapBrightness(brightness)
	}

	if slices.Contains(light.SupportedColorModes, supported_color_modes.ColorTemp) {
		light.ColorTemp = mapColorTemp(attrs)
	}

	return light
}

// mapSupportedColorModes maps a list of supported color modes from any type to the SupportedColorModes enum
func mapSupportedColorModes(entries []any) []supported_color_modes.SupportedColorModes {
	modes := make([]supported_color_modes.SupportedColorModes, 0)
	for _, value := range entries {
		mode, ok := value.(string)

		if !ok {
			log.Println("Invalid type assertion for mode from supported color modes.")
			continue
		}

		colorMode := supported_color_modes.EnumValue(mode)
		if colorMode == supported_color_modes.None {
			continue
		}

		modes = append(modes, colorMode)
	}

	return modes
}

// mapBrightness maps a brightness value from 0-255 to a percentage value from 0-100
func mapBrightness(brightnessValue any) int {
	switch v := brightnessValue.(type) {
	case float64:
		return int(v / 255 * 100)
	case int:
		return int(float64(v) / 255 * 100)
	default:
		return 0
	}
}

// mapColorTemp maps color temperature attributes to a ColorTemp model
func mapColorTemp(entries map[string]any) *models.ColorTemp {
	minColorTempKelvin, ok := entries["min_color_temp_kelvin"].(float64)
	if !ok {
		log.Println("Unable to determine min_color_temp_kelvin")
		return nil
	}

	maxColorTempKelvin, ok := entries["max_color_temp_kelvin"].(float64)
	if !ok {
		log.Println("Unable to determine max_color_temp_kelvin")
		return nil
	}

	temp, ok := entries["color_temp"].(float64)
	if !ok {
		temp = minColorTempKelvin
	}

	return &models.ColorTemp{
		MinTemp: int(minColorTempKelvin),
		MaxTemp: int(maxColorTempKelvin),
		Temp:    int(temp),
	}
}
