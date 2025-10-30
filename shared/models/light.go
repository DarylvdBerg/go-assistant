package models

import (
	"github.com/DarylvdBerg/go-assistant/shared/light_state"
)

type Light struct {
	EntityID     string            `json:"entity_id"`
	State        light_state.State `json:"state"`
	FriendlyName string            `json:"friendly_name"`
	Brightness   int
}

func (l Light) Title() string       { return l.FriendlyName }
func (l Light) Description() string { return "" }
func (l Light) FilterValue() string { return l.FriendlyName }
