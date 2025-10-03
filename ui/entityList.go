package ui

import (
	"fmt"
	"go-assistant-cli/internal/homeassistant"
	"io"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


var docStyle = lipgloss.NewStyle().Margin(1, 2)

var client *homeassistant.Client;

type lightDelegate struct {
	onStyle 			lipgloss.Style
	offStyle 			lipgloss.Style
	unavailableStyle	lipgloss.Style
}

func newLightDelegate() lightDelegate {
	return lightDelegate {
		onStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true),
		offStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
		unavailableStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
	}
}

func (d lightDelegate) Height() int { return 1 }
func (d lightDelegate) Spacing() int { return 1 }
func (d lightDelegate) ShowDescription() bool { return false }
func (d lightDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d lightDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
    light, ok := listItem.(homeassistant.Light)
    if !ok {
        return
    }

    var style lipgloss.Style
    switch light.State {
		case "on":
			style = d.onStyle
		case "unavailable":
			style = d.unavailableStyle
		default:
			style = d.offStyle
	}

    cursor := " "
    if index == m.Index() {
        cursor = ">"
        style = style.Underline(true)
    }

    fmt.Fprintf(w, "%s %s", cursor, style.Render(light.FilterValue()))
}


type entityList struct {
	list list.Model
}

func InitData(data []homeassistant.Light) entityList {	
	items := make([]list.Item, 0);

	for _, s := range data {
		items = append(items, s)
	}

	itemStyle := newLightDelegate()

	el := entityList{list: list.New(items, itemStyle, 0, 0)}
	el.list.Title = "Available lights"

	config, err := homeassistant.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %w", err)
	}
	client = homeassistant.CreateNewClient(config.BaseUrl, config.Token)

	return el;
} 

func (e entityList) Init() tea.Cmd {
	return nil;
}

func (e entityList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:

			switch msg.String() {
				case "q" :
					return e, tea.Quit
				case "p":
					return e.toggleLight()
			}

		case tea.WindowSizeMsg: 
			h, v := docStyle.GetFrameSize()
			e.list.SetSize(msg.Width - h, msg.Height - v)
	}

	var cmd tea.Cmd
	e.list, cmd = e.list.Update(msg)
	return e, cmd;
}

func (e entityList) View() string {
	return docStyle.Render(e.list.View())
}

func (e entityList) GetSelectedLight() *homeassistant.Light {
	if selectedItem := e.list.SelectedItem(); selectedItem != nil {
		if light, ok := selectedItem.(homeassistant.Light); ok {
			return &light
		}
	}

	return nil
}

func (e entityList) toggleLight() (tea.Model, tea.Cmd){
	light := e.GetSelectedLight();
	if light.State == "unavailable" {
		return e, nil
	}
	
	var newState string
	if light.State == "on" {
		client.TurnOffLight(light.EntityID)
		newState = "off"
	} else {
		client.TurnOnLight(light.EntityID)
		newState = "on"
	}

	e.updateLightState(light.EntityID, newState)
	return e, nil
}

func (e *entityList) updateLightState(entityID string, newState string) {
    items := e.list.Items()
    for i, item := range items {
        if light, ok := item.(homeassistant.Light); ok {
            if light.EntityID == entityID {
                // Update the light's state
                light.State = newState
                // Replace the item in the list
                e.list.SetItem(i, light)
                break
            }
        }
    }
}