package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	headerH = 2
	footerH = 1
)

type Model struct {
	width     int
	height    int
	activeTab int
	strings   Strings
	tabs      []Tab
}

func New() Model {
	return Model{
		width:   120,
		height:  40,
		strings: PT,
		tabs: []Tab{
			newConfig(),
			newAggregatedChat(),
		},
	}
}

func (m Model) bodyHeight() int {
	h := m.height - headerH - footerH
	if h < 1 {
		h = 1
	}
	return h
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		var cmds []tea.Cmd
		for i, t := range m.tabs {
			updated, cmd := t.Update(msg)
			m.tabs[i] = updated.(Tab)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)

	case tea.KeyMsg:
		if !m.isEditing() {
			if key.Matches(msg, DefaultKeys.Quit) {
				return m, tea.Quit
			}

			if msg.String() >= "1" && msg.String() <= "9" {
				idx := int(msg.String()[0]-'0') - 1
				if idx < len(m.tabs) {
					m.activeTab = idx
				}
				return m, nil
			}
		}
		return m.updateActiveTab(msg)
	}

	return m.updateActiveTab(msg)
}

func (m Model) updateActiveTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.tabs) == 0 {
		return m, nil
	}
	updated, cmd := m.tabs[m.activeTab].Update(msg)
	m.tabs[m.activeTab] = updated.(Tab)
	return m, cmd
}

func (m Model) isEditing() bool {
	if len(m.tabs) == 0 {
		return false
	}
	return m.tabs[m.activeTab].IsEditing()
}

func (m Model) View() string {
	header := m.renderHeader()

	body := lipgloss.NewStyle().
		Width(m.width).
		Height(m.bodyHeight()).
		MaxHeight(m.bodyHeight()).
		Render(m.renderBody())

	footer := m.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m Model) renderHeader() string {
	logo := "StreamPac v0.1"

	var tabParts []string
	for i, t := range m.tabs {
		label := fmt.Sprintf(" [%d] %s ", i+1, t.Title())
		if i == m.activeTab {
			tabParts = append(tabParts, lipgloss.NewStyle().Bold(true).Reverse(true).Render(label))
		} else {
			tabParts = append(tabParts, label)
		}
	}
	tabBar := strings.Join(tabParts, "")

	return logo + "\n" + tabBar
}

func (m Model) renderBody() string {
	if len(m.tabs) == 0 {
		return ""
	}
	return m.tabs[m.activeTab].View()
}

func (m Model) renderFooter() string {
	s := m.strings
	hints := s.Tabs(len(m.tabs)) + "  [" + HelpKey(DefaultKeys.Quit) + "] " + s.Quit
	gap := m.width - lipgloss.Width(hints)
	if gap < 0 {
		gap = 0
	}
	return hints + strings.Repeat(" ", gap)
}
