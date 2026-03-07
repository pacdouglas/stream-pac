package tui

// DashboardModel é a tela [1].
// Recebe bodyHeight do model raiz via baseTab — não conhece header/footer.
// Três painéis: status (esq) | chat scrollável (centro) | eventos (dir).

import (
	"fmt"
	"strings"

	"github.com/pacdouglas/stream-pac/internal/tui/fake"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DashboardModel struct {
	baseTab        // width, height(body), th, s via composição
	chatVP viewport.Model
	ready  bool
}

func newDashboard(width, height int, th Theme, s Strings) DashboardModel {
	d := DashboardModel{baseTab: newBaseTab(width, height, th, s)}
	d.initViewport()
	return d
}

func (d *DashboardModel) initViewport() {
	statusW, eventsW := 24, 24
	vpW := d.width - statusW - eventsW - 6
	// panel border(2) + title(1) + scroll footer(1) = 4
	vpH := d.height - 4
	if vpW < 10 {
		vpW = 10
	}
	if vpH < 1 {
		vpH = 1
	}
	d.chatVP = viewport.New(vpW, vpH)
	d.chatVP.SetContent(d.buildChatContent())
	d.chatVP.GotoBottom()
	d.ready = true
}

func (d DashboardModel) Update(msg tea.Msg) (DashboardModel, tea.Cmd) {
	d.baseTab = d.baseTab.handleBase(msg) // atualiza width/height/th/s
	switch msg.(type) {
	case tea.WindowSizeMsg:
		d.initViewport()
		return d, nil
	case themeMsg:
		d.chatVP.SetContent(d.buildChatContent())
		return d, nil
	case langMsg:
		d.chatVP.SetContent(d.buildChatContent())
		return d, nil
	}
	var cmd tea.Cmd
	d.chatVP, cmd = d.chatVP.Update(msg)
	return d, cmd
}

// FooterHints retorna os atalhos de teclado desta tela para o footer fixo.
func (d DashboardModel) FooterHints() string {
	th, s := d.th, d.s
	return th.FooterKeyStyle().Render("[1-5]") + th.DimStyle().Render(" navegar") + th.Sep() +
		th.Key("/") + th.DimStyle().Render(s.FooterSearch) + th.Sep() +
		th.Key("f") + th.DimStyle().Render(s.FooterFilter) + th.Sep() +
		th.Key("t") + th.DimStyle().Render("tema") + th.Sep() +
		th.Key("l") + th.DimStyle().Render("idioma") + th.Sep() +
		th.Key("?") + th.DimStyle().Render(s.FooterHelp) + th.Sep() +
		th.Key("q") + th.DimStyle().Render(s.FooterQuit)
}

func (d DashboardModel) View() string {
	if !d.ready {
		return "carregando..."
	}
	statusW, eventsW := 24, 24
	chatW := d.width - statusW - eventsW

	return lipgloss.JoinHorizontal(lipgloss.Top,
		d.renderStatus(statusW, d.height),
		d.renderChat(chatW, d.height),
		d.renderEvents(eventsW, d.height),
	)
}

// ── Status ────────────────────────────────────────────────────────────────────

func (d DashboardModel) renderStatus(w, h int) string {
	th, s := d.th, d.s
	v := fake.Viewers
	total := v.Twitch + v.Kick + v.YouTube

	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(th.TitleStyle().Render(s.DashPlatforms) + "\n\n")
	sb.WriteString(lipgloss.NewStyle().Foreground(th.Purple).Render("●") + " " +
		lipgloss.NewStyle().Foreground(th.Purple).Render("Twitch") + "   " +
		lipgloss.NewStyle().Foreground(th.Green).Render("●") + "\n")
	sb.WriteString(lipgloss.NewStyle().Foreground(th.Green).Render("●") + " " +
		lipgloss.NewStyle().Foreground(th.Green).Render("Kick") + "     " +
		lipgloss.NewStyle().Foreground(th.Green).Render("●") + "\n")
	sb.WriteString(th.DimStyle().Render("○  YouTube  ─") + "\n")
	sb.WriteString("\n" + hline(w-4, th) + "\n\n")

	sb.WriteString(th.TitleStyle().Render(s.DashViewers) + "\n\n")
	sb.WriteString(fmt.Sprintf(" %s   %s\n", s.DashTotal,
		lipgloss.NewStyle().Foreground(th.Yellow).Bold(true).Render(fmt.Sprint(total))))
	sb.WriteString(fmt.Sprintf(" Twitch  %s\n", th.TextStyle().Render(fmt.Sprint(v.Twitch))))
	sb.WriteString(fmt.Sprintf(" Kick    %s\n", th.TextStyle().Render(fmt.Sprint(v.Kick))))
	sb.WriteString(fmt.Sprintf(" YouTube %s\n", th.DimStyle().Render(fmt.Sprint(v.YouTube))))
	sb.WriteString("\n" + hline(w-4, th) + "\n\n")

	sb.WriteString(th.TitleStyle().Render(s.DashSession) + "\n\n")
	sb.WriteString(fmt.Sprintf(" %s  %s\n", s.DashUptime, th.TextStyle().Render("02:34:15")))
	sb.WriteString(fmt.Sprintf(" %s    %s\n", s.DashMessages, th.TextStyle().Render("1.432")))
	sb.WriteString(fmt.Sprintf(" %s %s\n", s.DashClients, th.TextStyle().Render("2")))

	return th.PanelStyle(w, h).Padding(0, 1).
		Render(th.TitleStyle().Render("Status") + "\n" + sb.String())
}

// ── Chat ──────────────────────────────────────────────────────────────────────

func (d DashboardModel) buildChatContent() string {
	var sb strings.Builder
	for _, msg := range fake.ChatMessages {
		sb.WriteString(d.formatChatLine(msg) + "\n")
	}
	return sb.String()
}

func (d DashboardModel) formatChatLine(msg fake.ChatMsg) string {
	th := d.th
	switch msg.Platform {
	case "sys":
		return " " + th.DimStyle().Render(msg.Text)
	case "tw":
		col := msg.Color
		if col == "" {
			col = "#9b72ff"
		}
		badge := lipgloss.NewStyle().Foreground(th.Purple).Bold(true).Render("[TW]")
		user := lipgloss.NewStyle().Foreground(lipgloss.Color(col)).Bold(true).Render(msg.User)
		return fmt.Sprintf(" %s %s: %s", badge, user, th.TextStyle().Render(msg.Text))
	case "ki":
		col := msg.Color
		if col == "" {
			col = "#53fc18"
		}
		badge := lipgloss.NewStyle().Foreground(th.Green).Bold(true).Render("[KI]")
		user := lipgloss.NewStyle().Foreground(lipgloss.Color(col)).Bold(true).Render(msg.User)
		return fmt.Sprintf(" %s %s: %s", badge, user, th.TextStyle().Render(msg.Text))
	case "yt":
		badge := lipgloss.NewStyle().Foreground(th.Red).Bold(true).Render("[YT]")
		user := lipgloss.NewStyle().Foreground(lipgloss.Color("#ff6868")).Bold(true).Render(msg.User)
		return fmt.Sprintf(" %s %s: %s", badge, user, th.TextStyle().Render(msg.Text))
	}
	return " " + msg.Text
}

func (d DashboardModel) renderChat(w, h int) string {
	th, s := d.th, d.s
	hint := th.DimStyle().Render("  [/] " + s.DashChatSearch + "  [f] " + s.DashChatFilter)
	title := th.TitleStyle().Render(s.DashChatTitle) + hint
	scrollFooter := d.buildScrollFooter(w - 4)
	return th.FocusedPanelStyle(w, h).Padding(0, 1).
		Render(title + "\n" + d.chatVP.View() + "\n" + scrollFooter)
}

func (d DashboardModel) buildScrollFooter(w int) string {
	th := d.th
	total := len(fake.ChatMessages)
	pct := d.chatVP.ScrollPercent()

	if d.chatVP.AtTop() {
		left := lipgloss.NewStyle().Foreground(th.Yellow).Bold(true).Render("▲ início")
		right := th.DimStyle().Render(fmt.Sprintf("%d msgs", total))
		gap := w - lipgloss.Width(left) - lipgloss.Width(right)
		if gap < 1 {
			gap = 1
		}
		return " " + left + strings.Repeat(" ", gap) + right
	}

	if d.chatVP.AtBottom() {
		left := lipgloss.NewStyle().Foreground(th.Green).Bold(true).Render("✓ ao vivo")
		right := th.DimStyle().Render(fmt.Sprintf("%d msgs", total))
		gap := w - lipgloss.Width(left) - lipgloss.Width(right)
		if gap < 1 {
			gap = 1
		}
		return " " + left + strings.Repeat(" ", gap) + right
	}

	barW := 10
	filled := int(pct * float64(barW))
	bar := lipgloss.NewStyle().Foreground(th.Blue).Render(strings.Repeat("█", filled)) +
		th.DimStyle().Render(strings.Repeat("░", barW-filled))
	approxLine := int(pct * float64(total))
	pctStr := th.DimStyle().Render(fmt.Sprintf("%d%%", int(pct*100)))
	lineStr := th.DimStyle().Render(fmt.Sprintf("~%d/%d", approxLine, total))
	return " " + th.DimStyle().Render("↑↓") + " " + bar + " " + pctStr + "  " + lineStr
}

// ── Eventos ───────────────────────────────────────────────────────────────────

func (d DashboardModel) renderEvents(w, h int) string {
	th, s := d.th, d.s
	var sb strings.Builder
	sb.WriteString("\n")
	for _, ev := range fake.Events {
		icon, col := platformIconColor(ev.Platform)
		evType := lipgloss.NewStyle().Foreground(lipgloss.Color(col)).Bold(true).Render(ev.Type)
		sb.WriteString(fmt.Sprintf(" %s %s\n", icon, evType))
		sb.WriteString(fmt.Sprintf(" %s\n", th.BoldText().Render(ev.User)))
		sb.WriteString(fmt.Sprintf(" %s\n\n", th.DimStyle().Render(ev.Detail+" · "+ev.Age)))
	}
	return th.PanelStyle(w, h).Padding(0, 1).
		Render(th.TitleStyle().Render(s.DashEvents) + "\n" + sb.String())
}

// ── Helper compartilhado ──────────────────────────────────────────────────────

func platformIconColor(platform string) (icon, color string) {
	switch platform {
	case "tw":
		return "●", "#9b72ff"
	case "ki":
		return "●", "#53fc18"
	case "yt":
		return "●", "#ff4444"
	}
	return "○", "#6875A8"
}
