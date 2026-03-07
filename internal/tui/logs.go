package tui

// LogsModel é a tela [5] — log técnico com filtro por subsistema.
// Recebe bodyHeight do model raiz via baseTab — não conhece header/footer.

import (
	"strings"

	"github.com/pacdouglas/stream-pac/internal/tui/fake"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type LogsModel struct {
	baseTab         // width, height(body), th, s via composição
	vp      viewport.Model
	filter  string
	sources []string
}

var logSources = []string{"server", "twitch", "kick", "youtube", "hub", "ia"}

func newLogs(width, height int, th Theme, s Strings) LogsModel {
	l := LogsModel{
		baseTab: newBaseTab(width, height, th, s),
		sources: logSources,
	}
	l.initViewport()
	return l
}

func (l *LogsModel) initViewport() {
	vpW := l.width - 4
	// filter bar(1) + panel border top(1) + title(1) + panel border bot(1) + hint(1) = 5
	vpH := l.height - 5
	if vpW < 10 {
		vpW = 10
	}
	if vpH < 1 {
		vpH = 1
	}
	l.vp = viewport.New(vpW, vpH)
	l.vp.SetContent(l.buildContent())
	l.vp.GotoBottom()
}

func (l LogsModel) Update(msg tea.Msg) (LogsModel, tea.Cmd) {
	l.baseTab = l.baseTab.handleBase(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		l.initViewport()
		return l, nil
	case themeMsg:
		l.vp.SetContent(l.buildContent())
		return l, nil
	case langMsg:
		return l, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "0":
			l.filter = ""
			l.vp.SetContent(l.buildContent())
			l.vp.GotoBottom()
			return l, nil
		}
		for i, src := range logSources {
			if msg.String() == string(rune('1'+i)) {
				l.filter = src
				l.vp.SetContent(l.buildContent())
				l.vp.GotoBottom()
				return l, nil
			}
		}
	}
	var cmd tea.Cmd
	l.vp, cmd = l.vp.Update(msg)
	return l, cmd
}

// FooterHints retorna os atalhos de teclado desta tela para o footer fixo.
func (l LogsModel) FooterHints() string {
	th, s := l.th, l.s
	return th.Key("↑↓") + th.DimStyle().Render(s.FooterScroll) + th.Sep() +
		th.Key("0-6") + th.DimStyle().Render("filtrar") + th.Sep() +
		th.Key("q") + th.DimStyle().Render(s.FooterQuit)
}

func (l LogsModel) View() string {
	th, s := l.th, l.s

	filterBar := l.renderFilterBar()
	// panel height = body height - filter bar(1) - hint(1)
	panelH := l.height - 2
	logPanel := th.FocusedPanelStyle(l.width, panelH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.LogTitle) + "\n" + l.vp.View())
	hint := " " + th.DimStyle().Render("[0] todos  [1-6] filtrar por fonte  "+s.LogFilterHint)

	return lipgloss.JoinVertical(lipgloss.Left, filterBar, logPanel, hint)
}

func (l LogsModel) renderFilterBar() string {
	th, s := l.th, l.s
	var parts []string

	allLabel := "[0] " + s.LogAll
	if l.filter == "" {
		parts = append(parts, lipgloss.NewStyle().
			Foreground(th.BG).Background(th.Yellow).Bold(true).Padding(0, 1).Render(allLabel))
	} else {
		parts = append(parts, lipgloss.NewStyle().
			Foreground(th.Dim).Padding(0, 1).Render(allLabel))
	}

	for i, src := range logSources {
		label := "[" + string(rune('1'+i)) + "] " + src
		if l.filter == src {
			parts = append(parts, lipgloss.NewStyle().
				Foreground(th.BG).Background(l.sourceColor(src)).Bold(true).Padding(0, 1).
				Render(label))
		} else {
			parts = append(parts, lipgloss.NewStyle().
				Foreground(th.Dim).Padding(0, 1).Render(label))
		}
	}

	bar := lipgloss.JoinHorizontal(lipgloss.Center, parts...)
	return " " + bar
}

func (l LogsModel) buildContent() string {
	th := l.th
	var sb strings.Builder

	colorOf := func(src string) lipgloss.Color {
		switch src {
		case "server":
			return th.Blue
		case "twitch":
			return th.Purple
		case "kick":
			return th.Green
		case "youtube":
			return th.Red
		case "hub":
			return th.Yellow
		case "ia":
			return th.Pink
		}
		return th.Muted
	}

	for _, entry := range fake.LogEntries {
		if l.filter != "" && !strings.EqualFold(entry.Source, l.filter) {
			continue
		}
		col := colorOf(entry.Source)
		ts := th.DimStyle().Render(entry.Timestamp)
		src := lipgloss.NewStyle().Foreground(col).Bold(true).
			Render("[" + padRight(entry.Source, 7) + "]")
		msg := th.MutedStyle().Render(entry.Message)
		sb.WriteString(" " + ts + "  " + src + "  " + msg + "\n")
	}
	return sb.String()
}

func (l LogsModel) sourceColor(src string) lipgloss.Color {
	th := l.th
	switch src {
	case "server":
		return th.Blue
	case "twitch":
		return th.Purple
	case "kick":
		return th.Green
	case "youtube":
		return th.Red
	case "hub":
		return th.Yellow
	case "ia":
		return th.Pink
	}
	return th.Muted
}

func (l LogsModel) ActiveFilter() string { return l.filter }

func (l LogsModel) VisibleLineCount() int {
	if l.filter == "" {
		return len(fake.LogEntries)
	}
	count := 0
	for _, e := range fake.LogEntries {
		if strings.EqualFold(e.Source, l.filter) {
			count++
		}
	}
	return count
}

func padRight(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
}
