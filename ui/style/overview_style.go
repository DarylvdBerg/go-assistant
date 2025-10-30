package style

import (
	"fmt"
	"io"
	"log"

	"github.com/DarylvdBerg/go-assistant/shared/light_state"
	"github.com/DarylvdBerg/go-assistant/shared/models"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type OverviewStyleDelegate struct {
	enabledStyle     lipgloss.Style
	disabledStyle    lipgloss.Style
	unavailableStyle lipgloss.Style
}

func NewOverviewStyleDelegate() OverviewStyleDelegate {
	return OverviewStyleDelegate{
		enabledStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true),
		disabledStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("9")),
		unavailableStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
	}
}

func (d OverviewStyleDelegate) Height() int                               { return 1 }
func (d OverviewStyleDelegate) Spacing() int                              { return 1 }
func (d OverviewStyleDelegate) ShowDescription() bool                     { return false }
func (d OverviewStyleDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d OverviewStyleDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	light, ok := listItem.(models.Light)
	if !ok {
		return
	}

	var style lipgloss.Style
	switch light.State {
	case light_state.On:
		style = d.enabledStyle
	case light_state.Unavailable:
		style = d.unavailableStyle
	default:
		style = d.disabledStyle
	}

	cursor := " "
	if index == m.Index() {
		cursor = ">"
		style = style.Underline(true)
	}

	_, err := fmt.Fprintf(w, "%s %s", cursor, style.Render(light.FilterValue()))
	if err != nil {
		log.Fatalf("failed to style item %d: %v", index, err)
	}
}
