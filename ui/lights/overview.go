package lights

import (
	"github.com/DarylvdBerg/go-assistant/shared/models"
	"github.com/DarylvdBerg/go-assistant/ui/lights/brightness"
	"github.com/DarylvdBerg/go-assistant/ui/lights/temperature"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LightList struct {
	list             list.Model
	keys             *KeyBindings
	brightnessPanel  *brightness.Panel
	temperaturePanel *temperature.Panel
	width            int
	height           int
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func InitLightOverview(lights []models.Light) LightList {
	lightData := initializeLightList(lights)
	el := LightList{list: lightData, keys: NewLightListKeyMap()}

	el.list.Title = "Available lights"
	el.list.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			el.keys.toggleLight,
			el.keys.brightnessControl,
		}
	}

	return el
}

func (e LightList) Init() tea.Cmd {
	return nil
}

func (e LightList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	model, t, done := e.isBrightnessPanelInView(msg)
	if done {
		return model, t
	}

	model, t, done = e.isTemperaturePanelInView(msg)
	if done {
		return model, t
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		// Don't handle any keys if we're filtering.
		if e.list.FilterState() == list.Filtering {
			break
		}
		// Handle key binding inputs
		return e.keys.HandleKeyPress(msg, e)

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		e.list.SetSize(msg.Width-h, msg.Height-v)
		e.width = msg.Width
		e.height = msg.Height
	}

	var cmd tea.Cmd
	e.list, cmd = e.list.Update(msg)
	return e, cmd
}

func (e LightList) isBrightnessPanelInView(msg tea.Msg) (tea.Model, tea.Cmd, bool) {
	if e.brightnessPanel != nil && e.brightnessPanel.IsOpen {
		var cmd tea.Cmd
		*e.brightnessPanel, cmd = e.brightnessPanel.Update(msg)

		if !e.brightnessPanel.IsOpen {
			e.brightnessPanel = nil
		}

		return e, cmd, true
	}
	return nil, nil, false
}

func (e LightList) isTemperaturePanelInView(msg tea.Msg) (tea.Model, tea.Cmd, bool) {
	if e.temperaturePanel != nil && e.temperaturePanel.IsOpen {
		var cmd tea.Cmd
		*e.temperaturePanel, cmd = e.temperaturePanel.Update(msg)

		if !e.temperaturePanel.IsOpen {
			e.temperaturePanel = nil
		}

		return e, cmd, true
	}
	return nil, nil, false
}

// View implements tea.Model.
func (e LightList) View() string {
	view := docStyle.Render(e.list.View())

	if e.brightnessPanel != nil && e.brightnessPanel.IsOpen {
		brightnessPanelView := e.brightnessPanel.View()
		view = lipgloss.Place(e.width, e.height, lipgloss.Center, lipgloss.Center, brightnessPanelView)
	}

	if e.temperaturePanel != nil && e.temperaturePanel.IsOpen {
		temperaturePanelView := e.temperaturePanel.View()
		view = lipgloss.Place(e.width, e.height, lipgloss.Center, lipgloss.Center, temperaturePanelView)
	}

	return view
}
