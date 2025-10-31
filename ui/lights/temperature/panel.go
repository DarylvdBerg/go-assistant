package temperature

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

func InitializeNewTemperaturePanel(light models.Light) *Panel {
	return &Panel{
		Panel: base.InitializeNewBasePanel(light),
		keys:  NewTemperatureKeyBindings(),
	}
}

func (t Panel) Init() tea.Cmd { return nil }

func (t Panel) Update(msg tea.Msg) (Panel, tea.Cmd) {
	if !t.IsOpen {
		return t, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return t.keys.HandleKeyPress(msg, t)
	case tea.WindowSizeMsg:
		t.UpdateWindowSize(msg)
		return t, nil
	}

	return t, nil
}

func (t Panel) View() string {
	if !t.IsOpen {
		return ""
	}

	// Create progress bar
	progressWidth := 40.0
	var filled float64

	if t.Light.ColorTemp.Temp == nil {
		filled = 0.0
	} else {
		filled = *t.Light.ColorTemp.Temp
	}

	progressBar := ""
	for i := 0.0; i < progressWidth; i++ {
		if i < filled {
			progressBar += style.DefaultProgressStyle().Render("█")
		} else {
			progressBar += style.DefaultProgressBackgroundStyle().Render("░")
		}
	}

	content := fmt.Sprintf(
		"Set Temperature for: %s\n\n%s\n\nTemperature: %d\n\n"+
			"Controls:\n"+
			"← → ±250    k j: ±1000\n"+
			"Enter: Apply    Esc: Cancel",
		t.Light.FilterValue(),
		progressBar,
		t.Light.ColorTemp.Temp,
	)

	return lipgloss.Place(t.Width, t.Height, lipgloss.Center, lipgloss.Center,
		style.DefaultPanelStyle().Render(content))
}
