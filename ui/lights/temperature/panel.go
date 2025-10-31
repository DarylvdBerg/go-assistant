package temperature

import (
	"github.com/DarylvdBerg/go-assistant/shared/models"
	"github.com/DarylvdBerg/go-assistant/ui/lights/base"
	tea "github.com/charmbracelet/bubbletea"
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

	return ""
}
