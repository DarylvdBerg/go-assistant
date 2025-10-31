package base

import tea "github.com/charmbracelet/bubbletea"

type KeyBindings interface {
	HandleKeyPress(input tea.KeyMsg, panel Panel) (Panel, tea.Cmd)
	InitializeKeyBindings() *KeyBindings
}
