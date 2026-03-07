package tui

import "github.com/charmbracelet/bubbles/key"

// KeyMap define todos os atalhos globais da aplicação.
type KeyMap struct {
	Tab1  key.Binding
	Tab2  key.Binding
	Tab3  key.Binding
	Tab4  key.Binding
	Tab5  key.Binding
	Help  key.Binding
	Quit  key.Binding
	Up    key.Binding
	Down  key.Binding
	Enter key.Binding
}

// DefaultKeys é a instância global dos atalhos.
var DefaultKeys = KeyMap{
	Tab1:  key.NewBinding(key.WithKeys("1"), key.WithHelp("1", "painel")),
	Tab2:  key.NewBinding(key.WithKeys("2"), key.WithHelp("2", "overlays")),
	Tab3:  key.NewBinding(key.WithKeys("3"), key.WithHelp("3", "config")),
	Tab4:  key.NewBinding(key.WithKeys("4"), key.WithHelp("4", "ia")),
	Tab5:  key.NewBinding(key.WithKeys("5"), key.WithHelp("5", "logs")),
	Help:  key.NewBinding(key.WithKeys("?"), key.WithHelp("?", "ajuda")),
	Quit:  key.NewBinding(key.WithKeys("ctrl+q"), key.WithHelp("ctrl+q", "sair")),
	Up:    key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("↑/k", "subir")),
	Down:  key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("↓/j", "descer")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "confirmar")),
}
