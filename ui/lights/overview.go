package lights

import (
	"go-assistant/shared/models"
	"go-assistant/ui/brightness"
	"go-assistant/ui/style"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LightList struct {
	list            list.Model
	keys            *KeyBindings
	brightnessPanel *brightness.Panel
	width           int
	height          int
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

func initializeLightList(lights []models.Light) list.Model {
	items := make([]list.Item, 0)

	for _, s := range lights {
		items = append(items, s)
	}

	return list.New(items, style.NewOverviewStyleDelegate(), 0, 0)
}

func (e LightList) Init() tea.Cmd {
	return nil
}

func (e LightList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if e.brightnessPanel != nil && e.brightnessPanel.IsOpen() {
		var cmd tea.Cmd
		*e.brightnessPanel, cmd = e.brightnessPanel.Update(msg)

		if !e.brightnessPanel.IsOpen() {
			e.brightnessPanel = nil
		}

		return e, cmd
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

// View implements tea.Model.
func (e LightList) View() string {
	view := docStyle.Render(e.list.View())

	if e.brightnessPanel != nil && e.brightnessPanel.IsOpen() {
		brightnessPanelView := e.brightnessPanel.View()
		view = lipgloss.Place(e.width, e.height, lipgloss.Center, lipgloss.Center, brightnessPanelView)
	}

	return view
}

func (e LightList) getSelectedLight() *models.Light {
	selectedItem := e.list.SelectedItem()
	if selectedItem == nil {
		return nil
	}

	light, ok := selectedItem.(models.Light)
	if !ok {
		return nil
	}

	return &light
}
