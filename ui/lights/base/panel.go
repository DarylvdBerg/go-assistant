package base

import (
	"github.com/DarylvdBerg/go-assistant/shared/models"
	tea "github.com/charmbracelet/bubbletea"
)

type Panel struct {
	Light   *models.Light
	IsOpen  bool
	Width   int
	Height  int
	OnApply func(*models.Light)
}

func InitializeNewBasePanel(light models.Light) Panel {
	return Panel{
		Light:  &light,
		IsOpen: true,
	}
}

func (p Panel) UpdateWindowSize(msg tea.WindowSizeMsg) {
	p.Height = msg.Height
	p.Width = msg.Width
}
