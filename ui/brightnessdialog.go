package ui

import (
	"fmt"
	"go-assistant-cli/internal/homeassistant"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type brightnessDialog struct {
    light      homeassistant.Light
    brightness int
    active     bool
}

var (
    dialogStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("62")).
        Padding(1, 2).
        Width(50).
        Align(lipgloss.Center)

    progressStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("205"))

    progressBgStyle = lipgloss.NewStyle().
        Foreground(lipgloss.Color("240"))
)

func newBrightnessDialog(light homeassistant.Light) brightnessDialog {
    // Start with current brightness or 50% if unknown
    brightness := 50
    if light.Brightness > 0 {
        brightness = int(light.Brightness * 100 / 255) // Convert from 0-255 to 0-100
    }
    
    return brightnessDialog{
        light:      light,
        brightness: brightness,
        active:     true,
    }
}

func (bd brightnessDialog) Init() tea.Cmd {
    return nil
}

func (bd brightnessDialog) Update(msg tea.Msg) (brightnessDialog, tea.Cmd) {
    if !bd.active {
        return bd, nil
    }

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "left", "h":
            if bd.brightness > 1 {
                bd.brightness--
            }
        case "right", "l":
            if bd.brightness < 100 {
                bd.brightness++
            }
        case "down", "j":
            if bd.brightness > 10 {
                bd.brightness -= 10
            } else {
                bd.brightness = 1
            }
        case "up", "k":
            if bd.brightness < 90 {
                bd.brightness += 10
            } else {
                bd.brightness = 100
            }
        case "enter":
            // Apply brightness
            brightnessValue := uint8(float64(bd.brightness) * 255 / 100) // Convert to 0-255 and cast to uint8
            client.ChangeBrightness(bd.light.EntityID, brightnessValue)
            bd.active = false
            return bd, nil
        case "esc":
            bd.active = false
            return bd, nil
        }
    }

    return bd, nil
}

func (bd brightnessDialog) View() string {
    if !bd.active {
        return ""
    }

    // Create progress bar
    progressWidth := 40
    filled := int(float64(progressWidth) * float64(bd.brightness) / 100)
    
    progressBar := ""
    for i := 0; i < progressWidth; i++ {
        if i < filled {
            progressBar += progressStyle.Render("█")
        } else {
            progressBar += progressBgStyle.Render("░")
        }
    }

    content := fmt.Sprintf(
        "Set Brightness for: %s\n\n%s\n\nBrightness: %d%%\n\n"+
        "Controls:\n"+
        "← → or h l: ±1    ↑ ↓ or k j: ±10\n"+
        "Enter: Apply    Esc: Cancel",
        bd.light.FilterValue(), // Using FilterValue() instead of Name
        progressBar,
        bd.brightness,
    )

    return lipgloss.Place(40, 15, lipgloss.Center, lipgloss.Center,
        dialogStyle.Render(content))
}

func (bd brightnessDialog) IsActive() bool {
    return bd.active
}