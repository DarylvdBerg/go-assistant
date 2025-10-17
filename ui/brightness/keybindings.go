package brightness

import (
	"go-assistant/internal/homeassistant"
	"go-assistant/shared"
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type brightnessKeyBindings struct {
	applyBrightness key.Binding
	cancel          key.Binding
	increaseByTen   key.Binding
	decreaseByTen   key.Binding
	increaseByFive  key.Binding
	decreaseByFive  key.Binding
}

func NewBrightnessKeyBindings() *brightnessKeyBindings {
	return &brightnessKeyBindings{
		applyBrightness: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Apply brightness"),
		),
		cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Quit panel"),
		),
		increaseByTen: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "Increase brightness by 10%"),
		),
		decreaseByTen: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "Decrease brightness by 10%"),
		),
		increaseByFive: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "Increase brightness by 5%"),
		),
		decreaseByFive: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "Decrease brightness by 5%"),
		),
	}
}
func (b brightnessKeyBindings) HandleKeyPress(input tea.KeyMsg, panel BrightnessPanel) (BrightnessPanel, tea.Cmd) {
	if !panel.isOpen {
		return panel, nil
	}

	switch {
	case key.Matches(input, b.cancel):
		panel.isOpen = false
		return panel, nil
	case key.Matches(input, b.applyBrightness):
		err := homeassistant.GetClient().ChangeBrightness(panel.light.EntityID, uint8(panel.light.Brightness))
		if err != nil {
			log.Printf("failed to change brightness: %v", err)
		}

		panel.light.State = shared.LightStateOn
		if panel.OnApply != nil {
			panel.OnApply(panel.light)
		}
		panel.isOpen = false
		return panel, nil
	case key.Matches(input, b.increaseByTen):
		if panel.light.Brightness >= 100 {
			panel.light.Brightness = 100
			return panel, nil
		}
		panel.light.Brightness += 10
		return panel, nil
	case key.Matches(input, b.decreaseByTen):
		if panel.light.Brightness <= 0 {
			panel.light.Brightness = 0
			return panel, nil
		}
		panel.light.Brightness -= 10
		return panel, nil

	case key.Matches(input, b.increaseByFive):
		if panel.light.Brightness >= 100 {
			panel.light.Brightness = 100
			return panel, nil
		}
		panel.light.Brightness += 5
		return panel, nil
	case key.Matches(input, b.decreaseByFive):
		if panel.light.Brightness <= 0 {
			panel.light.Brightness = 0
			return panel, nil
		}
		panel.light.Brightness -= 5
		return panel, nil
	}

	var cmd tea.Cmd
	return panel, cmd
}
