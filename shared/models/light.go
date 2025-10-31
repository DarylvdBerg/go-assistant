package models

import (
	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/supported_color_modes"
)

type Light struct {
	EntityID            string
	State               light_state.State
	FriendlyName        string
	Brightness          int
	SupportedColorModes supported_color_modes.SupportedColorModes
}

func (l Light) Title() string       { return l.FriendlyName }
func (l Light) Description() string { return "" }
func (l Light) FilterValue() string { return l.FriendlyName }
