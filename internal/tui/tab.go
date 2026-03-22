package tui

import tea "github.com/charmbracelet/bubbletea"

type Tab interface {
	tea.Model
	Title() string
	IsEditing() bool
}
