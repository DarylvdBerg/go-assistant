package lights

import (
	"github.com/DarylvdBerg/go-assistant/shared"
	"github.com/DarylvdBerg/go-assistant/ui/brightness"

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
