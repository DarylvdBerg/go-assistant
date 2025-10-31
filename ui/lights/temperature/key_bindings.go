package temperature

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type KeyBindings struct {
	applyTemperature     key.Binding
	cancel               key.Binding
	increaseByTwentyFive key.Binding
	decreaseByTwentyFive key.Binding
	increaseByHundred    key.Binding
	decreaseByHundred    key.Binding
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
		increaseByTwentyFive: key.NewBinding(
			key.WithKeys("right"),
			key.WithHelp("→", "Increase temperature by 25K"),
		),
		decreaseByTwentyFive: key.NewBinding(
			key.WithKeys("left"),
			key.WithHelp("←", "Decrease temperature by 25K"),
		),
		increaseByHundred: key.NewBinding(
			key.WithKeys("k"),
			key.WithHelp("k", "Increase temperature by 100K"),
		),
		decreaseByHundred: key.NewBinding(
			key.WithKeys("j"),
			key.WithHelp("j", "Decrease temperature by 100K"),
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
	}

	var cmd tea.Cmd
	return panel, cmd
}
