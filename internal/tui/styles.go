package tui

import "github.com/charmbracelet/lipgloss"

// hline retorna uma linha horizontal de largura w colorida pelo tema.
func hline(w int, th Theme) string {
	line := ""
	for i := 0; i < w; i++ {
		line += "─"
	}
	return th.DimStyle().Render(line)
}

// joinKey formata um atalho no footer: "[k] texto".
func joinKey(k, label string, th Theme) string {
	return th.FooterKeyStyle().Render("["+k+"]") + th.DimStyle().Render(" "+label)
}

// toastStyle retorna o estilo de um toast de sucesso.
func toastStyle(th Theme) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Green)
}
