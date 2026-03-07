// Testes de nível 1: lógica pura, sem terminal.
package tui

import (
	"testing"

	"github.com/pacdouglas/stream-pac/internal/tui/fake"
)

// ── Dados fake ────────────────────────────────────────────────────────────────

func TestFakeChatMessages_NotEmpty(t *testing.T) {
	if len(fake.ChatMessages) == 0 {
		t.Fatal("ChatMessages está vazio")
	}
}

func TestFakeChatMessages_PlatformsAreValid(t *testing.T) {
	valid := map[string]bool{"tw": true, "ki": true, "yt": true, "sys": true}
	for i, msg := range fake.ChatMessages {
		if !valid[msg.Platform] {
			t.Errorf("mensagem[%d]: plataforma inválida %q", i, msg.Platform)
		}
	}
}

func TestFakeOverlays_HaveValidDimensions(t *testing.T) {
	for i, ov := range fake.Overlays {
		if ov.Width <= 0 {
			t.Errorf("overlay[%d] %q: Width inválido (%d)", i, ov.Name, ov.Width)
		}
		if ov.Height <= 0 {
			t.Errorf("overlay[%d] %q: Height inválido (%d)", i, ov.Name, ov.Height)
		}
	}
}

func TestFakeLogEntries_SourcesAreValid(t *testing.T) {
	valid := map[string]bool{
		"server": true, "twitch": true, "kick": true,
		"youtube": true, "hub": true, "ia": true,
	}
	for i, e := range fake.LogEntries {
		if !valid[e.Source] {
			t.Errorf("log[%d]: fonte inválida %q", i, e.Source)
		}
	}
}

func TestFakeViewers_TotalIsPositive(t *testing.T) {
	v := fake.Viewers
	if v.Twitch+v.Kick+v.YouTube <= 0 {
		t.Error("total de viewers deve ser positivo")
	}
}

// ── Constantes de aba ─────────────────────────────────────────────────────────

func TestTabConstants_AreUnique(t *testing.T) {
	seen := map[int]bool{}
	for _, tab := range []int{TabDashboard, TabOverlays, TabConfig, TabAI, TabLogs} {
		if seen[tab] {
			t.Errorf("constante de aba duplicada: %d", tab)
		}
		seen[tab] = true
	}
}

func TestTabConstants_AreSequential(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	actual := []int{TabDashboard, TabOverlays, TabConfig, TabAI, TabLogs}
	for i, tab := range actual {
		if tab != expected[i] {
			t.Errorf("aba[%d]: esperado %d, obtido %d", i, expected[i], tab)
		}
	}
}

// ── Tema ──────────────────────────────────────────────────────────────────────

func TestTheme_DarkHasDistinctColors(t *testing.T) {
	th := Dark
	if th.Text == th.Dim {
		t.Error("Dark: Text e Dim não devem ser iguais")
	}
	if th.FocusBorder == th.Border {
		t.Error("Dark: FocusBorder e Border não devem ser iguais")
	}
}

func TestTheme_MidnightHasDistinctColors(t *testing.T) {
	th := Midnight
	if th.Text == th.Dim {
		t.Error("Midnight: Text e Dim não devem ser iguais")
	}
	if th.FocusBorder == th.Border {
		t.Error("Midnight: FocusBorder e Border não devem ser iguais")
	}
}

func TestTheme_DarkAndMidnightAreDifferent(t *testing.T) {
	if Dark.BG == Midnight.BG {
		t.Error("Dark e Midnight não devem ter o mesmo BG")
	}
	if Dark.FocusBorder == Midnight.FocusBorder {
		t.Error("Dark e Midnight não devem ter o mesmo FocusBorder")
	}
}

// ── i18n ──────────────────────────────────────────────────────────────────────

func TestStrings_PTHasAllTabs(t *testing.T) {
	s := PT
	tabs := []string{s.TabDashboard, s.TabOverlays, s.TabConfig, s.TabAI, s.TabLogs}
	for i, tab := range tabs {
		if tab == "" {
			t.Errorf("PT: tab[%d] está vazio", i)
		}
	}
}

func TestStrings_ENHasAllTabs(t *testing.T) {
	s := EN
	tabs := []string{s.TabDashboard, s.TabOverlays, s.TabConfig, s.TabAI, s.TabLogs}
	for i, tab := range tabs {
		if tab == "" {
			t.Errorf("EN: tab[%d] está vazio", i)
		}
	}
}

func TestStrings_PTAndENAreDifferent(t *testing.T) {
	if PT.TabDashboard == EN.TabDashboard {
		t.Errorf("PT.TabDashboard == EN.TabDashboard: %q — deveriam ser diferentes", PT.TabDashboard)
	}
	if PT.CfgSave == EN.CfgSave {
		t.Errorf("PT.CfgSave == EN.CfgSave: %q", PT.CfgSave)
	}
}

func TestStrings_NoCriticalFieldIsEmpty(t *testing.T) {
	for _, pair := range []struct {
		lang string
		s    Strings
	}{{"PT", PT}, {"EN", EN}} {
		s := pair.s
		fields := map[string]string{
			"OvCopied":   s.OvCopied,
			"CfgSaved":   s.CfgSaved,
			"AIGenerate": s.AIGenerate,
			"LogAll":     s.LogAll,
		}
		for name, val := range fields {
			if val == "" {
				t.Errorf("%s.%s está vazio", pair.lang, name)
			}
		}
	}
}

// ── platformIconColor ─────────────────────────────────────────────────────────

func TestPlatformIconColor_KnownPlatforms(t *testing.T) {
	cases := []struct{ platform, wantColor string }{
		{"tw", "#9b72ff"},
		{"ki", "#53fc18"},
		{"yt", "#ff4444"},
	}
	for _, c := range cases {
		_, color := platformIconColor(c.platform)
		if color != c.wantColor {
			t.Errorf("platformIconColor(%q) color = %q, want %q", c.platform, color, c.wantColor)
		}
	}
}
