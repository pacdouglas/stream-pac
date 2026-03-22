package tui

import tea "github.com/charmbracelet/bubbletea"

type AggregatedChatModel struct {
	width  int
	height int
}

func newAggregatedChat() AggregatedChatModel {
	return AggregatedChatModel{}
}

func (m AggregatedChatModel) Title() string   { return "AggregatedChat" }
func (m AggregatedChatModel) IsEditing() bool { return false }
func (m AggregatedChatModel) Init() tea.Cmd   { return nil }
func (m AggregatedChatModel) View() string    { return "  AggregatedChat tab - em construcao..." }

func (m AggregatedChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}
