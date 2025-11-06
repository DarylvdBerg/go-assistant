package temperature

import (
	"log"

	"github.com/DarylvdBerg/go-assistant/internal/homeassistant"
	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyBindings struct {
	applyTemperature key.Binding
	cancel           key.Binding
	increaseBy100K   key.Binding
	decreaseBy100K   key.Binding
	increaseByOneK   key.Binding
	decreaseByOneK   key.Binding
}

func NewTemperatureKeyBindings() *KeyBindings {
	return &KeyBindings{
		applyTemperature: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Apply temperature"),
		),
		cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Quit panel"),
		),
		increaseBy100K: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "Increase temperature by 100K"),
		),
		decreaseBy100K: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "Decrease temperature by 100K"),
		),
		increaseByOneK: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "Increase temperature by 1000K"),
		),
		decreaseByOneK: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "Decrease temperature by 1000K"),
		),
	}
}

func (k KeyBindings) HandleKeyPress(input tea.KeyMsg, panel Panel) (Panel, tea.Cmd) {
	if !panel.IsOpen {
		return panel, nil
	}

	switch {
	case key.Matches(input, k.cancel):
		panel.IsOpen = false
		return panel, nil
	case key.Matches(input, k.applyTemperature):
		err := homeassistant.GetClient().ChangeColorTemp(panel.Light.EntityID, panel.Light.ColorTemp.Temp)
		if err != nil {
			log.Printf("error applying color temperature change: %v", err)
		}
		panel.Light.State = light_state.On
		if panel.OnApply != nil {
			panel.OnApply(panel.Light)
		}
		panel.IsOpen = false
		return panel, nil
	case key.Matches(input, k.increaseBy100K):
		panel.Light.ColorTemp.Temp += 100

		if panel.Light.ColorTemp.Temp > panel.Light.ColorTemp.MaxTemp {
			panel.Light.ColorTemp.Temp = panel.Light.ColorTemp.MaxTemp
		}

		return panel, nil
	case key.Matches(input, k.decreaseBy100K):
		panel.Light.ColorTemp.Temp -= 100

		if panel.Light.ColorTemp.Temp < panel.Light.ColorTemp.MinTemp {
			panel.Light.ColorTemp.Temp = panel.Light.ColorTemp.MinTemp
		}

		return panel, nil
	case key.Matches(input, k.increaseByOneK):
		panel.Light.ColorTemp.Temp += 1000
		if panel.Light.ColorTemp.Temp > panel.Light.ColorTemp.MaxTemp {
			panel.Light.ColorTemp.Temp = panel.Light.ColorTemp.MaxTemp
		}

		return panel, nil
	case key.Matches(input, k.decreaseByOneK):
		panel.Light.ColorTemp.Temp -= 1000
		if panel.Light.ColorTemp.Temp < panel.Light.ColorTemp.MinTemp {
			panel.Light.ColorTemp.Temp = panel.Light.ColorTemp.MinTemp
		}

		return panel, nil
	}

	var cmd tea.Cmd
	return panel, cmd
}
