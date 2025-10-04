package brightness

import (
	"fmt"
	"go-assistant-cli/shared/models"
	"go-assistant-cli/ui/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type BrightnessPanel struct {
	light      *models.Light
	brightness int
	isOpen     bool
}

func NewBrightnessPanel(light models.Light) *BrightnessPanel {
	return &BrightnessPanel{
		light:      &light,
		brightness: 50,
		isOpen:     true,
	}
}

func (b BrightnessPanel) Init() tea.Cmd {
	return nil
}

func (b BrightnessPanel) Update(msg tea.Msg) (BrightnessPanel, tea.Cmd) {
	if !b.isOpen {
		return b, nil
	}

	return b, nil
}

func (b BrightnessPanel) View() string {
	if !b.isOpen {
		return ""
	}

	// Create progress bar
	progressWidth := 40
	filled := int(float64(progressWidth) * float64(b.brightness) / 100)

	progressBar := ""
	for i := 0; i < progressWidth; i++ {
		if i < filled {
			progressBar += style.DefaultProgressStyle().Render("█")
		} else {
			progressBar += style.DefaultPanelStyle().Render("░")
		}
	}

	content := fmt.Sprintf(
		"Set Brightness for: %s\n\n%s\n\nBrightness: %d%%\n\n"+
			"Controls:\n"+
			"← → or h l: ±1    ↑ ↓ or k j: ±10\n"+
			"Enter: Apply    Esc: Cancel",
		b.light.FilterValue(), // Using FilterValue() instead of Name
		progressBar,
		b.brightness,
	)

	return lipgloss.Place(40, 15, lipgloss.Center, lipgloss.Center,
		style.DefaultPanelStyle().Render(content))
}

func (b BrightnessPanel) IsOpen() bool {
	return b.isOpen
}
