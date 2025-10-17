package lights

import (
	"go-assistant-cli/internal/homeassistant"
	"go-assistant-cli/shared"
	"go-assistant-cli/shared/models"
	"go-assistant-cli/ui/brightness"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type lightListKeyMap struct {
	toggleLight       key.Binding
	brightnessControl key.Binding
}

const (
	TurnOnAction  = "turn_" + shared.LightStateOn
	TurnOffAction = "turn_" + shared.LightStateOff
)

func NewLightListKeyMap() *lightListKeyMap {
	return &lightListKeyMap{
		toggleLight: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "Toggle light on/off"),
		),
		brightnessControl: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "Open brightness control"),
		),
	}
}

func (l *lightListKeyMap) HandleKeyPress(input tea.KeyMsg, lightList lightList) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(input, l.toggleLight):
		light := lightList.getSelectedLight()
		updatedState := toggleLight(light)
		lightList.updateLightState(light.EntityID, updatedState)
		return lightList, nil
	case key.Matches(input, l.brightnessControl):
		light := lightList.getSelectedLight()
		if light == nil {
			return lightList, nil
		}

		if light.State == shared.LightStateUnavailable {
			return lightList, nil
		}

		dialog := brightness.NewBrightnessPanel(*light)
		lightList.brightnessPanel = dialog
		return lightList, nil
	}

	var cmd tea.Cmd
	lightList.list, cmd = lightList.list.Update(input)
	return lightList, cmd
}

func toggleLight(light *models.Light) string {
	var action string
	var newState string
	if light.State == shared.LightStateOn {
		action = TurnOffAction
		newState = shared.LightStateOff
	} else {
		action = TurnOnAction
		newState = shared.LightStateOn
	}

	homeassistant.GetClient().ToggleLightState(light.EntityID, action)
	return newState
}

func (e *lightList) updateLightState(entityID string, newState string) {
	items := e.list.Items()
	for i, item := range items {
		if light, ok := item.(models.Light); ok {
			if light.EntityID == entityID {
				// Update the light's state
				light.State = newState
				// Replace the item in the list
				e.list.SetItem(i, light)
				break
			}
		}
	}
}
