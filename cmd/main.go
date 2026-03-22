package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pacdouglas/stream-pac/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.New(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
