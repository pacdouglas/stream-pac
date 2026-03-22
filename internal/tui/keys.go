package tui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit key.Binding
}

var DefaultKeys = KeyMap{
	Quit: key.NewBinding(key.WithKeys("ctrl+q"), key.WithHelp("ctrl+q", "quit")),
}

func HelpKey(b key.Binding) string {
	return b.Help().Key
}
