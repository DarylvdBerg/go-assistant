package lights

import (
	"go-assistant-cli/shared/models"
	"go-assistant-cli/ui/brightness"
	"go-assistant-cli/ui/style"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type lightList struct {
	list            list.Model
	keys            *lightListKeyMap
	brightnessPanel *brightness.BrightnessPanel
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func InitLightOverview(lights []models.Light) lightList {
	list := initializeLightList(lights)
	el := lightList{list: list, keys: NewLightListKeyMap()}

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

func (e lightList) Init() tea.Cmd {
	return nil
}

func (e lightList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
	}

	var cmd tea.Cmd
	e.list, cmd = e.list.Update(msg)
	return e, cmd
}

// View implements tea.Model.
func (e lightList) View() string {
	view := docStyle.Render(e.list.View())

	if e.brightnessPanel != nil && e.brightnessPanel.IsOpen() {
		brightnessPanelView := e.brightnessPanel.View()
		view = lipgloss.Place(40, 15, lipgloss.Center, lipgloss.Center, brightnessPanelView)
	}

	return view
}

func (e lightList) getSelectedLight() *models.Light {
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
