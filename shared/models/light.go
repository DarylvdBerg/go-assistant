package models

import (
	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/supported_color_modes"
)

type Light struct {
	// EntityID is the unique identifier of the light
	EntityID string
	// State indicates the current state of the light (on/off/unavailable)
	State light_state.State
	// FriendlyName is the human-readable name of the light
	FriendlyName string
	// Brightness indicates the brightness of the light from 0-255
	Brightness int
	// SupportedColorModes indicates which color modes the light supports
	SupportedColorModes []supported_color_modes.SupportedColorModes
	// ColorTemp holds information about the color temperature of the light, nil if the light does not support color temperature
	ColorTemp *ColorTemp
}

type ColorTemp struct {
	// MinTemp indicates the minimum temperature the light supports in kelvin
	MinTemp int
	// Temp indicates the current temperature the light is set to in kelvin, pointer value because the api can return nil when the light is in a State.Off state.
	Temp int
	// MaxTemp indicates the maximum temperature the light supports in kelvin
	MaxTemp int
}

func (l Light) Title() string       { return l.FriendlyName }
func (l Light) Description() string { return "" }
func (l Light) FilterValue() string { return l.FriendlyName }
