package brightness

import (
	"fmt"

	"github.com/DarylvdBerg/go-assistant/shared/models"
	"github.com/DarylvdBerg/go-assistant/ui/lights/base"
	"github.com/DarylvdBerg/go-assistant/ui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Panel struct {
	base.Panel
	keys *KeyBindings
}

func InitializeNewBrightnessPanel(light models.Light) *Panel {
	return &Panel{
		Panel: base.InitializeNewBasePanel(light),
		keys:  NewBrightnessKeyBindings(),
	}
}

func (b Panel) Init() tea.Cmd {
	return nil
}

func (b Panel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	if !b.IsOpen {
		return b, nil
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		return b.keys.HandleKeyPress(msg, b)
	case tea.WindowSizeMsg:
		b.UpdateWindowSize(msg)
		return b, nil
	}

	return b, nil
}

func (b Panel) View() string {
	if !b.IsOpen {
		return ""
	}

	// Create progress bar
	progressWidth := 40
	filled := int(float64(progressWidth) * float64(b.Light.Brightness) / 100)

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
		b.Light.FilterValue(),
		progressBar,
		b.Light.Brightness,
	)

	return lipgloss.Place(b.Width, b.Height, lipgloss.Center, lipgloss.Center,
		style.DefaultPanelStyle().Render(content))
}
