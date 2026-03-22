package tui

import tea "github.com/charmbracelet/bubbletea"

type ConfigModel struct {
	width  int
	height int
}

func newConfig() ConfigModel {
	return ConfigModel{}
}

func (m ConfigModel) Title() string    { return "Config" }
func (m ConfigModel) IsEditing() bool  { return false }
func (m ConfigModel) Init() tea.Cmd    { return nil }
func (m ConfigModel) View() string     { return "  Config tab - em construcao..." }

func (m ConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}
