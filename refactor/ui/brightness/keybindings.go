package brightness

import "github.com/charmbracelet/bubbles/key"

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
			key.WithKeys("l"),
			key.WithHelp("l", "Increase brightness by 5%"),
		),
		decreaseByFive: key.NewBinding(
			key.WithKeys("h"),
			key.WithHelp("h", "Decrease brightness by 5%"),
		),
	}
}
