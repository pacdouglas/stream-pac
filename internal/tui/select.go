package tui

// SelectModel é um combobox/dropdown interativo feito do zero para Bubble Tea.
//
// Comportamento:
//   - Fechado: mostra a opção selecionada com indicador ▼
//   - Aberto: mostra todas as opções inline, com cursor ▶
//   - Enter/Espaço: abre/fecha e confirma seleção
//   - j/k/↑/↓: navega entre opções quando aberto
//   - Esc: fecha sem confirmar

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SelectModel mantém o estado completo do combobox.
type SelectModel struct {
	options  []string
	cursor   int
	selected int
	open     bool
	focused  bool
	th       Theme
	width    int
}

// NewSelect cria um SelectModel com as opções e largura fornecidas.
func NewSelect(options []string, initial int, width int, th Theme) SelectModel {
	return SelectModel{
		options:  options,
		cursor:   initial,
		selected: initial,
		width:    width,
		th:       th,
	}
}

// Value retorna o texto da opção selecionada.
func (s SelectModel) Value() string {
	if s.selected >= 0 && s.selected < len(s.options) {
		return s.options[s.selected]
	}
	return ""
}

// IsOpen indica se o dropdown está expandido.
func (s SelectModel) IsOpen() bool { return s.open }

// IsFocused indica se o widget está com foco.
func (s SelectModel) IsFocused() bool { return s.focused }

// Focus coloca o widget em modo focado.
func (s *SelectModel) Focus() { s.focused = true }

// Blur remove o foco (fecha o dropdown também).
func (s *SelectModel) Blur() {
	s.focused = false
	s.open = false
}

// Update processa mensagens. Só reage a teclas quando focused == true.
func (s SelectModel) Update(msg tea.Msg) (SelectModel, tea.Cmd) {
	switch msg := msg.(type) {
	case themeMsg:
		s.th = msg.theme
		return s, nil
	case tea.KeyMsg:
		if !s.focused {
			return s, nil
		}
		switch msg.String() {
		case "enter", " ":
			if s.open {
				s.selected = s.cursor
				s.open = false
			} else {
				s.cursor = s.selected
				s.open = true
			}
		case "esc":
			if s.open {
				s.cursor = s.selected
				s.open = false
			}
		case "up", "k":
			if s.open && s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.open && s.cursor < len(s.options)-1 {
				s.cursor++
			}
		}
	}
	return s, nil
}

// View renderiza o combobox.
func (s SelectModel) View() string {
	th := s.th

	borderColor := th.Border
	valueColor := th.Dim
	arrowColor := th.Dim
	if s.focused {
		borderColor = th.FocusBorder
		valueColor = th.Text
		arrowColor = th.Yellow
	}

	value := lipgloss.NewStyle().Foreground(valueColor).Render(s.Value())
	arrow := lipgloss.NewStyle().Foreground(arrowColor).Bold(true).Render("▼")
	gap := s.width - lipgloss.Width(value) - lipgloss.Width(arrow)
	if gap < 1 {
		gap = 1
	}
	inner := value + strings.Repeat(" ", gap) + arrow

	field := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		Width(s.width).
		Render(inner)

	if !s.open {
		return field
	}

	rows := make([]string, len(s.options))
	for i, opt := range s.options {
		if i == s.cursor {
			rows[i] = lipgloss.NewStyle().
				Foreground(th.BG).Background(th.Yellow).Bold(true).
				Render("▶ " + opt + strings.Repeat(" ", s.width-2-len([]rune(opt))))
		} else {
			rows[i] = lipgloss.NewStyle().Foreground(th.Text).
				Render("  " + opt)
		}
	}

	dropdown := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.FocusBorder).
		Width(s.width).
		Render(strings.Join(rows, "\n"))

	return lipgloss.JoinVertical(lipgloss.Left, field, dropdown)
}
