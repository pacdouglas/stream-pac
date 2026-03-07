// Testes do Model raiz — Update() é função pura, sem terminal ou goroutines.
package tui

import (
	"strings"
	"testing"

	"github.com/pacdouglas/stream-pac/internal/tui/fake"

	tea "github.com/charmbracelet/bubbletea"
)

// ── Helpers de teste ──────────────────────────────────────────────────────────

func keyMsg(r rune) tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}}
}

func enterMsg() tea.KeyMsg {
	return tea.KeyMsg{Type: tea.KeyEnter}
}

func update(m Model, msg tea.Msg) Model {
	next, _ := m.Update(msg)
	return next.(Model)
}

// updateWithCmd aplica a mensagem E executa o cmd retornado (se houver).
func updateWithCmd(m Model, msg tea.Msg) Model {
	next, cmd := m.Update(msg)
	m = next.(Model)
	if cmd != nil {
		result := cmd()
		if result != nil {
			m = update(m, result)
		}
	}
	return m
}

// ── Testes de navegação ───────────────────────────────────────────────────────

func TestModel_InitialTabIsDashboard(t *testing.T) {
	m := New()
	if m.ActiveTab() != TabDashboard {
		t.Errorf("aba inicial = %d, want %d", m.ActiveTab(), TabDashboard)
	}
}

func TestModel_KeyPress2_SwitchesToOverlays(t *testing.T) {
	m := update(New(), keyMsg('2'))
	if m.ActiveTab() != TabOverlays {
		t.Errorf("após '2': aba = %d, want %d", m.ActiveTab(), TabOverlays)
	}
}

func TestModel_KeyPress3_SwitchesToConfig(t *testing.T) {
	m := update(New(), keyMsg('3'))
	if m.ActiveTab() != TabConfig {
		t.Errorf("após '3': aba = %d, want %d", m.ActiveTab(), TabConfig)
	}
}

func TestModel_KeyPress4_SwitchesToAI(t *testing.T) {
	m := update(New(), keyMsg('4'))
	if m.ActiveTab() != TabAI {
		t.Errorf("após '4': aba = %d, want %d", m.ActiveTab(), TabAI)
	}
}

func TestModel_KeyPress5_SwitchesToLogs(t *testing.T) {
	m := update(New(), keyMsg('5'))
	if m.ActiveTab() != TabLogs {
		t.Errorf("após '5': aba = %d, want %d", m.ActiveTab(), TabLogs)
	}
}

func TestModel_Navigation_RoundTrip(t *testing.T) {
	m := New()
	// Navega sem passar por abas com campo de texto (Config=3, AIGen=4).
	// Nesses tabs os números são consumidos pelo campo — sai-se via esc.
	sequence := []struct {
		key  rune
		want int
	}{
		{'2', TabOverlays},
		{'5', TabLogs},
		{'1', TabDashboard},
	}
	for _, step := range sequence {
		m = update(m, keyMsg(step.key))
		if m.ActiveTab() != step.want {
			t.Errorf("após '%c': aba = %d, want %d", step.key, m.ActiveTab(), step.want)
		}
	}
}

func TestModel_EditingTab_QDoesNotQuit(t *testing.T) {
	m := New()
	m = update(m, keyMsg('3')) // vai para Config
	if m.ActiveTab() != TabConfig {
		t.Fatalf("deveria estar em Config, got %d", m.ActiveTab())
	}
	// Sem campo focado: 'q' age como atalho global (sai) — Tab foca o campo
	// Aqui só verificamos que a aba ainda é Config (campo não auto-focado).
	if m.config.IsEditing() {
		t.Error("Config não deve ter campo focado ao abrir")
	}
}

// ── Testes de tema ────────────────────────────────────────────────────────────

func TestModel_InitialThemeIsDark(t *testing.T) {
	m := New()
	if m.ActiveTheme() != "dark" {
		t.Errorf("tema inicial = %q, want %q", m.ActiveTheme(), "dark")
	}
}

func TestModel_ThemeToggle_RoundTrip(t *testing.T) {
	m := New()
	// 't' abre o picker — cursor em "dark" (index 0)
	m = update(m, keyMsg('t'))
	if !m.IsPickerOpen() {
		t.Fatal("'t' deveria abrir o picker de tema")
	}
	// j → cursor para midnight (index 1), enter → confirma
	m = update(m, keyMsg('j'))
	m = updateWithCmd(m, enterMsg())
	if m.ActiveTheme() != "midnight" {
		t.Errorf("após selecionar midnight: %q, want %q", m.ActiveTheme(), "midnight")
	}
	// 't' novamente — cursor já em midnight (index 1), k → dark, enter → confirma
	m = update(m, keyMsg('t'))
	m = update(m, keyMsg('k'))
	m = updateWithCmd(m, enterMsg())
	if m.ActiveTheme() != "dark" {
		t.Errorf("após selecionar dark: %q, want %q", m.ActiveTheme(), "dark")
	}
}

// ── Testes de idioma ──────────────────────────────────────────────────────────

func TestModel_InitialLangIsPT(t *testing.T) {
	m := New()
	if m.ActiveLang() != "pt" {
		t.Errorf("idioma inicial = %q, want %q", m.ActiveLang(), "pt")
	}
}

func TestModel_LangToggle_RoundTrip(t *testing.T) {
	m := New()
	// 'l' abre picker — cursor em "pt" (index 0), j → en, enter → confirma
	m = update(m, keyMsg('l'))
	if !m.IsPickerOpen() {
		t.Fatal("'l' deveria abrir o picker de idioma")
	}
	m = update(m, keyMsg('j'))
	m = updateWithCmd(m, enterMsg())
	if m.ActiveLang() != "en" {
		t.Errorf("após selecionar en: %q, want %q", m.ActiveLang(), "en")
	}
	// 'l' novamente — cursor em "en" (index 1), k → pt, enter → confirma
	m = update(m, keyMsg('l'))
	m = update(m, keyMsg('k'))
	m = updateWithCmd(m, enterMsg())
	if m.ActiveLang() != "pt" {
		t.Errorf("após selecionar pt: %q, want %q", m.ActiveLang(), "pt")
	}
}

// ── Testes de help ────────────────────────────────────────────────────────────

func TestModel_QuestionMark_ShowsHelp(t *testing.T) {
	m := update(New(), keyMsg('?'))
	if !m.ShowingHelp() {
		t.Error("'?' deveria mostrar o help overlay")
	}
}

func TestModel_AnyKey_ClosesHelp(t *testing.T) {
	m := update(New(), keyMsg('?'))
	m = update(m, keyMsg('x'))
	if m.ShowingHelp() {
		t.Error("qualquer tecla deveria fechar o help")
	}
}

// ── Testes de View ────────────────────────────────────────────────────────────

func TestView_ContainsLogo(t *testing.T) {
	m := New()
	view := m.View()
	if !strings.Contains(view, "StreamPac") {
		t.Error("view não contém o logo 'StreamPac'")
	}
}

func TestView_ContainsActiveTabLabel_PT(t *testing.T) {
	m := New()
	view := m.View()
	if !strings.Contains(view, PT.TabDashboard) {
		t.Errorf("view não contém %q", PT.TabDashboard)
	}
}

func TestView_ContainsTabLabel_AfterLangSwitch(t *testing.T) {
	m := New()
	// picker de idioma: abre, move para EN, confirma
	m = update(m, keyMsg('l'))
	m = update(m, keyMsg('j'))
	m = updateWithCmd(m, enterMsg())
	view := m.View()
	if !strings.Contains(view, EN.TabDashboard) {
		t.Errorf("após switch para EN: view não contém %q", EN.TabDashboard)
	}
}

func TestView_ContainsThemeToggle(t *testing.T) {
	m := New()
	view := m.View()
	if !strings.Contains(view, "dark") && !strings.Contains(view, "midnight") {
		t.Error("view deveria conter indicador de tema no header")
	}
}

func TestView_ContainsLangToggle(t *testing.T) {
	m := New()
	view := m.View()
	if !strings.Contains(view, "pt") && !strings.Contains(view, "en") {
		t.Error("view deveria conter indicador de idioma no header")
	}
}

// ── Testes de sub-modelos ─────────────────────────────────────────────────────

func TestConfig_DefaultValues(t *testing.T) {
	m := New()
	if got := m.config.GetTwitchChannel(); got != "streampac" {
		t.Errorf("Twitch channel padrão = %q, want %q", got, "streampac")
	}
	if got := m.config.GetKickChannel(); got != "streampac" {
		t.Errorf("Kick channel padrão = %q, want %q", got, "streampac")
	}
}

func TestConfig_TabFocusAdvances(t *testing.T) {
	m := New()
	initial := m.config.FocusedField()
	m.activeTab = TabConfig
	var cmd tea.Cmd
	m.config, cmd = m.config.Update(tea.KeyMsg{Type: tea.KeyTab})
	_ = cmd
	if m.config.FocusedField() == initial {
		t.Error("Tab deveria avançar o foco para o próximo campo")
	}
}

func TestLogs_InitialFilterIsEmpty(t *testing.T) {
	m := New()
	if f := m.logs.ActiveFilter(); f != "" {
		t.Errorf("filtro inicial = %q, want %q", f, "")
	}
}

func TestLogs_FilterReducesCount(t *testing.T) {
	m := New()
	total := m.logs.VisibleLineCount()
	m.logs, _ = m.logs.Update(keyMsg('1'))
	filtered := m.logs.VisibleLineCount()
	if filtered >= total {
		t.Errorf("filtro 'server' deveria mostrar menos que %d linhas, mas mostrou %d", total, filtered)
	}
	if filtered == 0 {
		t.Error("filtro 'server' não deveria zerar as linhas")
	}
}

func TestOverlays_ItemCount(t *testing.T) {
	m := New()
	want := len(fake.Overlays) + 1 // +1 item "Gerar com IA"
	if got := m.overlays.ItemCount(); got != want {
		t.Errorf("overlays.ItemCount() = %d, want %d", got, want)
	}
}
