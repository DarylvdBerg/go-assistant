package lights

import (
	"slices"

	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/models"
	"github.com/DarylvdBerg/go-assistant/shared/supported_color_modes"
	"github.com/DarylvdBerg/go-assistant/ui/lights/brightness"
	"github.com/DarylvdBerg/go-assistant/ui/lights/temperature"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyBindings struct {
	toggleLight        key.Binding
	brightnessControl  key.Binding
	temperatureControl key.Binding
}

var TurnOnAction = "turn_" + light_state.On.StringValue()
var TurnOffAction = "turn_" + light_state.Off.StringValue()

func NewLightListKeyMap() *KeyBindings {
	return &KeyBindings{
		toggleLight: key.NewBinding(
			key.WithKeys("p"),
			key.WithHelp("p", "Toggle light on/off"),
		),
		brightnessControl: key.NewBinding(
			key.WithKeys("b"),
			key.WithHelp("b", "Open brightness control"),
		),
		temperatureControl: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "Open temperature control"),
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
		light, ok := l.getLight(lightList)
		if !ok {
			return lightList, nil
		}

		dialog := brightness.InitializeNewBrightnessPanel(*light)
		dialog.OnApply = lightList.updateLightState

		lightList.brightnessPanel = dialog
		return lightList, nil
	case key.Matches(input, l.temperatureControl):
		light, ok := l.getLight(lightList)
		if !ok {
			return lightList, nil
		}

		if !slices.Contains(light.SupportedColorModes, supported_color_modes.ColorTemp) {
			return lightList, nil
		}

		dialog := temperature.InitializeNewTemperaturePanel(*light)
		dialog.OnApply = lightList.updateLightState

		lightList.temperaturePanel = dialog
	}

	var cmd tea.Cmd
	lightList.list, cmd = lightList.list.Update(input)
	return lightList, cmd
}

func (l *KeyBindings) getLight(lightList LightList) (*models.Light, bool) {
	light := lightList.getSelectedLight()
	if light == nil {
		return nil, false
	}

	if light.State == light_state.Unavailable {
		return nil, false
	}
	return light, true
}
