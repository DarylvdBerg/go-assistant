package base

import (
	"github.com/DarylvdBerg/go-assistant/shared/models"
	tea "github.com/charmbracelet/bubbletea"
)

type Panel interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Panel, tea.Cmd)
	NewPanel(light models.Light) *Panel
	View() string
}
