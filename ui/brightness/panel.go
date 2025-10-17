package brightness

import (
	"fmt"
	"go-assistant/shared/models"
	"go-assistant/ui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Panel struct {
	light   *models.Light
	isOpen  bool
	keys    *KeyBindings
	width   int
	height  int
	OnApply func(*models.Light)
}

func NewBrightnessPanel(light models.Light) *Panel {
	return &Panel{
		light:  &light,
		isOpen: true,
		keys:   NewBrightnessKeyBindings(),
	}
}

func (b Panel) Init() tea.Cmd {
	return nil
}

func (b Panel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	if !b.isOpen {
		return b, nil
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		return b.keys.HandleKeyPress(msg, b)

	case tea.WindowSizeMsg:
		b.height = msg.Height
		b.width = msg.Width
		return b, nil
	}

	return b, nil
}

func (b Panel) View() string {
	if !b.isOpen {
		return ""
	}

	// Create progress bar
	progressWidth := 40
	filled := int(float64(progressWidth) * float64(b.light.Brightness) / 100)

	progressBar := ""
	for i := 0; i < progressWidth; i++ {
		if i < filled {
			progressBar += style.DefaultProgressStyle().Render("█")
		} else {
			progressBar += style.DefaultProgressBackgroundStyle().Render("░")
		}
	}

	content := fmt.Sprintf(
		"Set Brightness for: %s\n\n%s\n\nBrightness: %d%%\n\n"+
			"Controls:\n"+
			"← → ±10    k j: ±5\n"+
			"Enter: Apply    Esc: Cancel",
		b.light.FilterValue(),
		progressBar,
		b.light.Brightness,
	)

	return lipgloss.Place(b.width, b.height, lipgloss.Center, lipgloss.Center,
		style.DefaultPanelStyle().Render(content))
}

func (b Panel) IsOpen() bool {
	return b.isOpen
}
