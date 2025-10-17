package lights

import (
	"go-assistant/internal/homeassistant"
	"go-assistant/shared"
	"go-assistant/shared/models"
	"go-assistant/ui/brightness"
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyBindings struct {
	toggleLight       key.Binding
	brightnessControl key.Binding
}

const (
	TurnOnAction  = "turn_" + shared.LightStateOn
	TurnOffAction = "turn_" + shared.LightStateOff
)

func NewLightListKeyMap() *KeyBindings {
	return &KeyBindings{
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

func (l *KeyBindings) HandleKeyPress(input tea.KeyMsg, lightList LightList) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(input, l.toggleLight):
		light := lightList.getSelectedLight()
		toggleLight(light)
		lightList.updateLightState(light)
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
		dialog.OnApply = lightList.updateLightState

		lightList.brightnessPanel = dialog
		return lightList, nil
	}

	var cmd tea.Cmd
	lightList.list, cmd = lightList.list.Update(input)
	return lightList, cmd
}

func toggleLight(light *models.Light) {
	var action string
	if light.State == shared.LightStateOn {
		action = TurnOffAction
		light.State = shared.LightStateOff
	} else {
		action = TurnOnAction
		light.State = shared.LightStateOn
	}

	err := homeassistant.GetClient().ToggleLightState(light.EntityID, action)
	if err != nil {
		log.Fatal("failed to toggle light state: ", err)
	}
}

func (e *LightList) updateLightState(updatedLight *models.Light) {
	items := e.list.Items()
	for i, item := range items {
		if light, ok := item.(models.Light); ok {
			if light.EntityID == updatedLight.EntityID {
				// Update the light's state
				light.State = updatedLight.State
				light.Brightness = updatedLight.Brightness
				// Replace the item in the list
				e.list.SetItem(i, light)
				break
			}
		}
	}
}
