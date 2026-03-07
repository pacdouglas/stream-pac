package tui

// PickerModel é um popup genérico centralizado na tela.
// Usado para: seleção de tema, idioma, confirmações, e qualquer lista futura.
//
// Uso:
//   p := newPicker("Título", options, currentValue, theme)
//   // no Update: p, result = p.Update(msg)
//   // no View:   return p.View(totalW, totalH)

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// PickerOption representa uma entrada no popup.
type PickerOption struct {
	Label  string
	Value  string
	Colors []lipgloss.Color // swatches de cor opcionais (ex: preview de tema)
}

// PickerResult indica o que aconteceu após um Update.
type PickerResult int

const (
	PickerPending   PickerResult = iota // nada aconteceu
	PickerConfirmed                     // enter → opção confirmada
	PickerDismissed                     // esc   → fechado sem seleção
)

// PickerModel mantém o estado do popup.
type PickerModel struct {
	title   string
	options []PickerOption
	cursor  int
	th      Theme
}

// newPicker cria um picker posicionando o cursor na opção com valor == currentValue.
func newPicker(title string, options []PickerOption, currentValue string, th Theme) PickerModel {
	cursor := 0
	for i, o := range options {
		if o.Value == currentValue {
			cursor = i
			break
		}
	}
	return PickerModel{title: title, options: options, cursor: cursor, th: th}
}

// Selected retorna a opção atualmente sob o cursor.
func (p PickerModel) Selected() PickerOption {
	if p.cursor >= 0 && p.cursor < len(p.options) {
		return p.options[p.cursor]
	}
	return PickerOption{}
}

// Update processa eventos de teclado do picker.
func (p PickerModel) Update(msg tea.Msg) (PickerModel, PickerResult) {
	km, ok := msg.(tea.KeyMsg)
	if !ok {
		return p, PickerPending
	}
	switch km.String() {
	case "up", "k":
		if p.cursor > 0 {
			p.cursor--
		}
	case "down", "j":
		if p.cursor < len(p.options)-1 {
			p.cursor++
		}
	case "enter", " ":
		return p, PickerConfirmed
	case "esc":
		return p, PickerDismissed
	}
	return p, PickerPending
}

// View renderiza o popup centralizado dentro de totalW × totalH.
// O espaço ao redor é preenchido com a cor BG do tema, criando o efeito de overlay.
func (p PickerModel) View(totalW, totalH int) string {
	th := p.th
	rowW := 32

	var rows []string
	for i, opt := range p.options {
		swatches := ""
		for _, c := range opt.Colors {
			swatches += lipgloss.NewStyle().Foreground(c).Render("■")
		}
		if swatches != "" {
			swatches += " "
		}
		label := swatches + opt.Label

		if i == p.cursor {
			rows = append(rows, lipgloss.NewStyle().
				Foreground(th.BG).
				Background(th.FocusBorder).
				Bold(true).
				Width(rowW).
				Render("▶ "+label))
		} else {
			rows = append(rows, lipgloss.NewStyle().
				Foreground(th.Text).
				Width(rowW).
				Render("  "+label))
		}
	}

	hint := th.DimStyle().Render("↑↓ navegar  ·  enter confirmar  ·  esc fechar")

	content := th.TitleStyle().Render(p.title) + "\n\n" +
		strings.Join(rows, "\n") +
		"\n\n" + hint

	popup := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.FocusBorder).
		Background(th.Panel).
		Padding(1, 2).
		Render(content)

	return lipgloss.Place(
		totalW, totalH,
		lipgloss.Center, lipgloss.Center,
		popup,
		lipgloss.WithWhitespaceBackground(th.BG),
	)
}
