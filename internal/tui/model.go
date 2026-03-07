// Package tui — scaffold raiz do StreamPac (Bubble Tea).
//
// SCAFFOLD FIXO
// ═════════════
//  linha 0   │ logo bar  (logo + toggles tema/idioma)          ← header fixo
//  linha 1   │ tab bar   (todas as abas sempre visíveis)        ← header fixo
//  linhas 2… │ body      (View() do sub-modelo ativo)           ← conteúdo variável
//  linha -1  │ footer    (status plataformas + notif + relógio) ← footer fixo
//
// O body é renderizado com lipgloss.Height(bodyH) que FORÇA exatamente bodyH
// linhas — o sub-modelo não precisa saber nada do scaffold.
// O root model distribui bodyH via bodyMsg() para que os sub-modelos
// dimensionem seus viewports corretamente.
package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/pacdouglas/stream-pac/internal/tui/fake"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ── Constantes de aba ─────────────────────────────────────────────────────────

const (
	TabDashboard = 1
	TabOverlays  = 2
	TabConfig    = 3
	TabAI        = 4
	TabLogs      = 5
)

// headerH e footerH são fixos — alterar aqui ajusta tudo automaticamente.
const (
	headerH = 2 // linha 0: logo bar  |  linha 1: tab bar
	footerH = 1 // linha final: status + notif + clock
)

// ── Mensagens ─────────────────────────────────────────────────────────────────

type switchTabMsg int
type tickMsg time.Time

// notifyMsg permite que qualquer sub-modelo envie uma notificação para o footer.
// Uso: return m, func() tea.Msg { return notifyMsg("URL copiada!") }
type notifyMsg string

type themeMsg struct {
	name  string
	theme Theme
}

type langMsg struct {
	lang    string
	strings Strings
}

// ── Picker: qual popup está aberto ────────────────────────────────────────────

type activePicker int

const (
	pickerNone    activePicker = iota
	pickerTheme                // [t] — seleção de tema
	pickerLang                 // [l] — seleção de idioma
	pickerConfirm              // ctrl+q — confirmação de saída
)

// ── Model raiz ────────────────────────────────────────────────────────────────

type Model struct {
	width        int
	height       int
	activeTab    int
	theme        Theme
	lang         string
	strings      Strings
	showHelp     bool
	clock        string // atualizado pelo ticker a cada segundo
	notification string // último evento/notificação para o footer

	picker       PickerModel  // popup ativo (reutilizado para todos)
	activePickr  activePicker // qual popup está aberto

	tabZones []tabZone

	// Sub-modelos — recebem bodyHeight, nunca conhecem o scaffold.
	dashboard DashboardModel
	overlays  OverlaysModel
	config    ConfigModel
	aigen     AIGenModel
	logs      LogsModel
}

// langsByCode mapeia código de idioma → Strings.
var langsByCode = map[string]Strings{
	"pt": PT,
	"en": EN,
}

type tabZone struct{ from, to, tab int }

func (m Model) bodyHeight() int {
	h := m.height - headerH - footerH
	if h < 1 {
		h = 1
	}
	return h
}

func (m Model) bodyMsg() tea.WindowSizeMsg {
	return tea.WindowSizeMsg{Width: m.width, Height: m.bodyHeight()}
}

// New cria o Model inicial. Dimensões padrão enquanto o terminal não envia WindowSizeMsg.
func New() Model {
	th := Dark
	s := PT
	w, h := 120, 40
	bh := h - headerH - footerH

	// Notificação inicial: último evento fake
	ev := fake.Events[0]
	notif := "⚡ " + ev.Type + " de " + ev.User + " — " + ev.Detail

	m := Model{
		width:        w,
		height:       h,
		activeTab:    TabDashboard,
		theme:        th,
		lang:         "pt",
		strings:      s,
		clock:        time.Now().Format("15:04:05"),
		notification: notif,
		dashboard:    newDashboard(w, bh, th, s),
		overlays:     newOverlays(w, bh, th, s),
		config:       newConfig(w, bh, th, s),
		aigen:        newAIGen(w, bh, th, s),
		logs:         newLogs(w, bh, th, s),
	}
	m.tabZones = m.computeTabZones()
	return m
}

// Init inicia o ticker do relógio.
func (m Model) Init() tea.Cmd {
	return tickCmd()
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// computeTabZones calcula as posições X das abas na linha 1 (tab bar).
func (m Model) computeTabZones() []tabZone {
	th, s := m.theme, m.strings
	defs := []struct{ idx int; label string }{
		{TabDashboard, s.TabDashboard},
		{TabOverlays, s.TabOverlays},
		{TabConfig, s.TabConfig},
		{TabAI, s.TabAI},
		{TabLogs, s.TabLogs},
	}
	pos := 1 // padding esquerdo " " na tab bar
	zones := make([]tabZone, 0, len(defs))
	for _, t := range defs {
		label := fmt.Sprintf("[%d] %s", t.idx, t.label)
		w := lipgloss.Width(th.InactiveTabStyle().Render(label))
		zones = append(zones, tabZone{from: pos, to: pos + w, tab: t.idx})
		pos += w
	}
	return zones
}

func (m Model) tabAtX(x int) int {
	for _, z := range m.tabZones {
		if x >= z.from && x < z.to {
			return z.tab
		}
	}
	return 0
}

// ── Update ────────────────────────────────────────────────────────────────────

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.tabZones = m.computeTabZones()
		bm := m.bodyMsg()
		m.dashboard, _ = m.dashboard.Update(bm)
		m.overlays, _ = m.overlays.Update(bm)
		m.config, _ = m.config.Update(bm)
		m.aigen, _ = m.aigen.Update(bm)
		m.logs, _ = m.logs.Update(bm)
		return m, nil

	case tickMsg:
		m.clock = time.Time(msg).Format("15:04:05")
		return m, tickCmd() // mantém o ticker vivo

	case notifyMsg:
		m.notification = string(msg)
		return m, nil

	case switchTabMsg:
		m.activeTab = int(msg)
		return m, nil

	case themeMsg:
		m.theme = msg.theme
		m.tabZones = m.computeTabZones()
		m.dashboard, _ = m.dashboard.Update(msg)
		m.overlays, _ = m.overlays.Update(msg)
		m.config, _ = m.config.Update(msg)
		m.aigen, _ = m.aigen.Update(msg)
		m.logs, _ = m.logs.Update(msg)
		return m, nil

	case langMsg:
		m.lang = msg.lang
		m.strings = msg.strings
		m.tabZones = m.computeTabZones()
		m.dashboard, _ = m.dashboard.Update(msg)
		m.overlays, _ = m.overlays.Update(msg)
		m.config, _ = m.config.Update(msg)
		m.aigen, _ = m.aigen.Update(msg)
		m.logs, _ = m.logs.Update(msg)
		return m, nil

	case tea.MouseMsg:
		if m.showHelp {
			m.showHelp = false
			return m, nil
		}
		// Y==1: tab bar (linha 1 do header)
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft && msg.Y == 1 {
			if tab := m.tabAtX(msg.X); tab != 0 {
				m.activeTab = tab
				return m, nil
			}
		}
		return m.updateActiveTab(msg)

	case genStepMsg, genDoneMsg:
		var cmd tea.Cmd
		m.aigen, cmd = m.aigen.Update(msg)
		return m, cmd

	case tea.KeyMsg:
		// ctrl+c bloqueado — saída apenas via ctrl+q com confirmação.
		if msg.String() == "ctrl+c" {
			return m, nil
		}

		// Se há um picker aberto, ele consome todos os eventos.
		if m.activePickr != pickerNone {
			var result PickerResult
			m.picker, result = m.picker.Update(msg)
			switch result {
			case PickerConfirmed:
				return m.applyPicker()
			case PickerDismissed:
				m.activePickr = pickerNone
			}
			return m, nil
		}

		// Quando há campo de texto focado, nada é interceptado globalmente.
		if m.isEditingText() {
			return m.updateActiveTab(msg)
		}

		if m.showHelp {
			m.showHelp = false
			return m, nil
		}

		switch {
		case key.Matches(msg, DefaultKeys.Quit):
			m.picker = newPicker("Sair do StreamPac?", quitPickerOptions(), "n", m.theme)
			m.activePickr = pickerConfirm
			return m, nil
		case key.Matches(msg, DefaultKeys.Help):
			m.showHelp = true
			return m, nil
		case key.Matches(msg, DefaultKeys.Tab1):
			m.activeTab = TabDashboard
			return m, nil
		case key.Matches(msg, DefaultKeys.Tab2):
			m.activeTab = TabOverlays
			return m, nil
		case key.Matches(msg, DefaultKeys.Tab3):
			m.activeTab = TabConfig
			return m, nil
		case key.Matches(msg, DefaultKeys.Tab4):
			m.activeTab = TabAI
			return m, nil
		case key.Matches(msg, DefaultKeys.Tab5):
			m.activeTab = TabLogs
			return m, nil
		case msg.String() == "t":
			m.picker = newPicker("Tema", themePickerOptions(), m.theme.Name, m.theme)
			m.activePickr = pickerTheme
			return m, nil
		case msg.String() == "l":
			m.picker = newPicker("Idioma", langPickerOptions(), m.lang, m.theme)
			m.activePickr = pickerLang
			return m, nil
		}
		return m.updateActiveTab(msg)
	}

	return m.updateActiveTab(msg)
}

// applyPicker aplica a seleção do picker ativo e fecha o popup.
func (m Model) applyPicker() (tea.Model, tea.Cmd) {
	sel := m.picker.Selected()
	ap := m.activePickr
	m.activePickr = pickerNone

	switch ap {
	case pickerTheme:
		if th, ok := ThemesByName[sel.Value]; ok {
			return m, func() tea.Msg { return themeMsg{name: sel.Value, theme: th} }
		}
	case pickerLang:
		if s, ok := langsByCode[sel.Value]; ok {
			return m, func() tea.Msg { return langMsg{lang: sel.Value, strings: s} }
		}
	case pickerConfirm:
		if sel.Value == "y" {
			return m, tea.Quit
		}
	}
	return m, nil
}

// themePickerOptions retorna as opções de tema com swatches de cor.
func themePickerOptions() []PickerOption {
	return []PickerOption{
		{Label: "Arcade", Value: "dark", Colors: []lipgloss.Color{Dark.Yellow, Dark.Blue, Dark.Pink}},
		{Label: "Midnight", Value: "midnight", Colors: []lipgloss.Color{Midnight.Blue, Midnight.Purple, Midnight.Green}},
		{Label: "Dracula", Value: "dracula", Colors: []lipgloss.Color{Dracula.Purple, Dracula.Pink, Dracula.Cyan}},
		{Label: "Nord", Value: "nord", Colors: []lipgloss.Color{Nord.Cyan, Nord.Blue, Nord.Purple}},
		{Label: "Tokyo Night", Value: "tokyo", Colors: []lipgloss.Color{Tokyo.Blue, Tokyo.Purple, Tokyo.Pink}},
		{Label: "Gruvbox", Value: "gruvbox", Colors: []lipgloss.Color{Gruvbox.Yellow, Gruvbox.Red, Gruvbox.Green}},
		{Label: "Catppuccin", Value: "catppuccin", Colors: []lipgloss.Color{Catppuccin.Purple, Catppuccin.Pink, Catppuccin.Blue}},
		{Label: "Solarized", Value: "solarized", Colors: []lipgloss.Color{Solarized.Cyan, Solarized.Blue, Solarized.Yellow}},
		{Label: "Monokai", Value: "monokai", Colors: []lipgloss.Color{Monokai.Green, Monokai.Yellow, Monokai.Pink}},
		{Label: "Matrix", Value: "matrix", Colors: []lipgloss.Color{Matrix.Green, Matrix.Cyan, Matrix.Yellow}},
	}
}

// langPickerOptions retorna as opções de idioma disponíveis.
func langPickerOptions() []PickerOption {
	return []PickerOption{
		{Label: "Português", Value: "pt"},
		{Label: "English", Value: "en"},
	}
}

// quitPickerOptions retorna as opções da confirmação de saída.
func quitPickerOptions() []PickerOption {
	return []PickerOption{
		{Label: "Não, continuar", Value: "n"},
		{Label: "Sim, sair", Value: "y"},
	}
}

func (m Model) updateActiveTab(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.activeTab {
	case TabDashboard:
		m.dashboard, cmd = m.dashboard.Update(msg)
	case TabOverlays:
		m.overlays, cmd = m.overlays.Update(msg)
	case TabConfig:
		m.config, cmd = m.config.Update(msg)
	case TabAI:
		m.aigen, cmd = m.aigen.Update(msg)
	case TabLogs:
		m.logs, cmd = m.logs.Update(msg)
	}
	return m, cmd
}

// ── View ──────────────────────────────────────────────────────────────────────
//
// Scaffold garantido:
//   header (2 linhas fixas) +
//   body   (Height(bodyH) força exatamente bodyH linhas) +
//   footer (1 linha fixa)
//
// Independente do que View() do sub-modelo retornar, o body nunca
// vai transbordar nem empurrar o footer para fora da tela.

func (m Model) View() string {
	th := m.theme

	// Picker ativo: renderiza o popup centrado sobre o fundo BG.
	if m.activePickr != pickerNone {
		return m.picker.View(m.width, m.height)
	}

	if m.showHelp {
		return lipgloss.NewStyle().
			Background(th.BG).Width(m.width).Height(m.height).
			Render(m.renderHelp())
	}

	header := m.renderHeader()

	body := lipgloss.NewStyle().
		Width(m.width).
		Height(m.bodyHeight()).
		MaxHeight(m.bodyHeight()). // clipa conteúdo que transborda
		Render(m.renderBody())

	footer := m.renderFooter()

	// Height+MaxHeight garante exatamente m.height linhas — scaffold fixo.
	return lipgloss.NewStyle().
		Background(th.BG).
		Width(m.width).
		Height(m.height).
		MaxHeight(m.height).
		Render(lipgloss.JoinVertical(lipgloss.Left, header, body, footer))
}

// ── Header ────────────────────────────────────────────────────────────────────

// renderHeader retorna SEMPRE exatamente 2 linhas:
//
//	linha 0: logo + versão  (esq)  ·  [t] tema  [l] idioma  (dir)
//	linha 1: tab bar com todas as abas sobre fundo Panel
func (m Model) renderHeader() string {
	th, s := m.theme, m.strings

	// ── linha 0: logo bar ─────────────────────────────────────────────────
	logo := lipgloss.NewStyle().Foreground(th.Yellow).Bold(true).Render("ᗧ···") +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render("StreamPac")
	ver := th.DimStyle().Render("  v0.1")
	logoSide := logo + ver

	swatch := lipgloss.NewStyle().Foreground(th.FocusBorder).Bold(true).Render("■")
	themeToggle := th.DimStyle().Render("[t]") + " " + swatch + " " +
		lipgloss.NewStyle().Foreground(th.Cyan).Bold(true).Render(m.theme.Name)
	langToggle := th.DimStyle().Render("[l]") + " " +
		lipgloss.NewStyle().Foreground(th.Cyan).Bold(true).Render(m.lang)
	right := "  " + themeToggle + "   " + langToggle + " "

	gap := m.width - lipgloss.Width(logoSide) - lipgloss.Width(right)
	if gap < 0 {
		gap = 0
	}
	line0 := logoSide + strings.Repeat(" ", gap) + right

	// ── linha 1: tab bar ──────────────────────────────────────────────────
	defs := []struct{ idx int; label string }{
		{TabDashboard, s.TabDashboard},
		{TabOverlays, s.TabOverlays},
		{TabConfig, s.TabConfig},
		{TabAI, s.TabAI},
		{TabLogs, s.TabLogs},
	}
	var tabParts []string
	for _, t := range defs {
		label := fmt.Sprintf("[%d] %s", t.idx, t.label)
		if t.idx == m.activeTab {
			tabParts = append(tabParts, th.ActiveTabStyle().Render(label))
		} else {
			tabParts = append(tabParts, th.InactiveTabStyle().Render(label))
		}
	}
	tabRow := " " + lipgloss.JoinHorizontal(lipgloss.Center, tabParts...)
	line1 := lipgloss.NewStyle().
		Background(th.Panel).
		Width(m.width).
		Render(tabRow)

	return lipgloss.JoinVertical(lipgloss.Left, line0, line1)
}

// ── Body ──────────────────────────────────────────────────────────────────────

func (m Model) renderBody() string {
	switch m.activeTab {
	case TabDashboard:
		return m.dashboard.View()
	case TabOverlays:
		return m.overlays.View()
	case TabConfig:
		return m.config.View()
	case TabAI:
		return m.aigen.View()
	case TabLogs:
		return m.logs.View()
	}
	return ""
}

// ── Footer ────────────────────────────────────────────────────────────────────

// renderFooter retorna SEMPRE exatamente 1 linha:
//
//	esq: status das plataformas + notificação do último evento
//	dir: atalhos globais rápidos + relógio
func (m Model) renderFooter() string {
	th := m.theme

	// Status das plataformas
	tw := lipgloss.NewStyle().Foreground(th.Purple).Bold(true).Render("●") +
		lipgloss.NewStyle().Foreground(th.Purple).Render(" TW")
	ki := lipgloss.NewStyle().Foreground(th.Green).Bold(true).Render("●") +
		lipgloss.NewStyle().Foreground(th.Green).Render(" KI")
	yt := th.DimStyle().Render("○ YT")
	status := tw + "  " + ki + "  " + yt

	// Notificação
	notif := ""
	if m.notification != "" {
		notif = th.Sep() +
			lipgloss.NewStyle().Foreground(th.Yellow).Render(m.notification)
	}

	// Atalhos globais compactos
	hints := th.Sep() +
		th.FooterKeyStyle().Render("[?]") + th.DimStyle().Render(" ajuda") +
		th.Sep() +
		th.FooterKeyStyle().Render("[ctrl+q]") + th.DimStyle().Render(" sair")

	// Relógio
	clock := lipgloss.NewStyle().Foreground(th.Cyan).Bold(true).Render(m.clock)

	left := " " + status + notif + hints
	right := clock + " "

	gap := m.width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap < 1 {
		gap = 1
	}
	return left + strings.Repeat(" ", gap) + right
}

// ── Help overlay ──────────────────────────────────────────────────────────────

func (m Model) renderHelp() string {
	th, s := m.theme, m.strings
	content := "\n" +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render(s.HelpGlobal) + "\n\n" +
		" " + th.DimStyle().Render("[1-5]") + "      navegar entre telas\n" +
		" " + th.DimStyle().Render("[t]") + "        trocar tema claro/escuro\n" +
		" " + th.DimStyle().Render("[l]") + "        trocar idioma PT/EN\n" +
		" " + th.DimStyle().Render("[q/ctrl+c]") + " " + s.FooterQuit + "\n" +
		" " + th.DimStyle().Render("[?]") + "        " + s.HelpClose + " este painel\n\n" +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render("Dashboard") + "\n\n" +
		" " + th.DimStyle().Render("[↑↓]") + "       scroll no chat\n" +
		" " + th.DimStyle().Render("[/]") + "        buscar no histórico\n\n" +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render("Overlays") + "\n\n" +
		" " + th.DimStyle().Render("[↑↓]") + "       navegar lista\n" +
		" " + th.DimStyle().Render("[enter]") + "    copiar URL\n" +
		" " + th.DimStyle().Render("[o]") + "        abrir no browser\n" +
		" " + th.DimStyle().Render("[/]") + "        filtrar lista\n\n" +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render("Config") + "\n\n" +
		" " + th.DimStyle().Render("[tab]") + "      próximo campo\n" +
		" " + th.DimStyle().Render("[enter]") + "    salvar\n" +
		" " + th.DimStyle().Render("[esc]") + "      descartar\n\n" +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render("IA") + "\n\n" +
		" " + th.DimStyle().Render("[ctrl+enter]") + " gerar overlay\n\n" +
		" " + lipgloss.NewStyle().Foreground(th.Blue).Bold(true).Render("Logs") + "\n\n" +
		" " + th.DimStyle().Render("[0]") + "        todos\n" +
		" " + th.DimStyle().Render("[1-6]") + "      filtrar por fonte\n" +
		" " + th.DimStyle().Render("[↑↓/PgUp]") + "  scroll\n"

	helpW, helpH := 52, 36
	panel := th.FocusedPanelStyle(helpW, helpH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.HelpTitle) + content)

	padLeft := (m.width - helpW) / 2
	padTop := (m.height - helpH) / 2
	if padLeft < 0 {
		padLeft = 0
	}
	if padTop < 0 {
		padTop = 0
	}
	lines := strings.Split(panel, "\n")
	var sb strings.Builder
	sb.WriteString(strings.Repeat("\n", padTop))
	for _, line := range lines {
		sb.WriteString(strings.Repeat(" ", padLeft) + line + "\n")
	}
	return sb.String()
}

// ── Acessores para testes ─────────────────────────────────────────────────────

// isEditingText retorna true quando um campo de texto está realmente focado
// no tab ativo — apenas nesse momento atalhos globais são suspensos.
func (m Model) isEditingText() bool {
	switch m.activeTab {
	case TabConfig:
		return m.config.IsEditing()
	case TabAI:
		return m.aigen.IsEditing()
	}
	return false
}

func (m Model) ActiveTab() int      { return m.activeTab }
func (m Model) ActiveTheme() string { return m.theme.Name }
func (m Model) ActiveLang() string  { return m.lang }
func (m Model) ShowingHelp() bool   { return m.showHelp }
func (m Model) IsPickerOpen() bool  { return m.activePickr != pickerNone }
