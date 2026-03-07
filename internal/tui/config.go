package tui

// ConfigModel é a tela [3] — formulário de configuração.
// Recebe bodyHeight do model raiz via baseTab — não conhece header/footer.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ConfigModel struct {
	baseTab       // width, height(body), th, s via composição
	focused int
	inputs  []textinput.Model
	sel     SelectModel
	saved   bool
}

const (
	cfgTwitchCh = iota
	cfgKickCh
	cfgKickRoom
	cfgYouTubeID
	cfgPort
	cfgHost
	cfgAPIKey
	cfgFieldCount // = 7
)

const cfgProvider    = cfgFieldCount     // = 7
const cfgTotalFields = cfgFieldCount + 1 // = 8

func newConfig(width, height int, th Theme, s Strings) ConfigModel {
	inputs := make([]textinput.Model, cfgFieldCount)

	setup := []struct {
		placeholder, value string
		width              int
		password           bool
	}{
		{"streampac", "streampac", 22, false},
		{"streampac", "streampac", 22, false},
		{"auto-detect", "", 22, false},
		{"Fpfdw0iXuv8", "", 22, false},
		{"7777", "7777", 8, false},
		{"localhost", "localhost", 16, false},
		{"sk-ant-...", "", 28, true},
	}

	for i, cfg := range setup {
		t := textinput.New()
		t.Placeholder = cfg.placeholder
		t.SetValue(cfg.value)
		t.Width = cfg.width
		if cfg.password {
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'
		}
		t.PromptStyle = lipgloss.NewStyle().Foreground(th.Blue)
		t.TextStyle = lipgloss.NewStyle().Foreground(th.Text)
		inputs[i] = t
	}
	// Sem auto-focus — usuário usa Tab ou clique para entrar nos campos.

	providers := []string{"Ollama (local)", "Anthropic Claude", "OpenAI GPT"}
	sel := NewSelect(providers, 1, 20, th)

	return ConfigModel{baseTab: newBaseTab(width, height, th, s), inputs: inputs, sel: sel}
}

func (c ConfigModel) Update(msg tea.Msg) (ConfigModel, tea.Cmd) {
	c.baseTab = c.baseTab.handleBase(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return c, nil
	case themeMsg:
		for i := range c.inputs {
			c.inputs[i].PromptStyle = lipgloss.NewStyle().Foreground(c.th.Blue)
			c.inputs[i].TextStyle = lipgloss.NewStyle().Foreground(c.th.Text)
		}
		c.sel, _ = c.sel.Update(msg)
		return c, nil
	case langMsg:
		return c, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if c.sel.IsOpen() {
				break
			}
			c.blurCurrent()
			c.focused = (c.focused + 1) % cfgTotalFields
			c.focusCurrent()
			return c, nil
		case "shift+tab":
			if c.sel.IsOpen() {
				break
			}
			c.blurCurrent()
			c.focused = (c.focused - 1 + cfgTotalFields) % cfgTotalFields
			c.focusCurrent()
			return c, nil
		case "enter":
			if c.focused == cfgProvider {
				break
			}
			c.saved = true
			return c, nil
		case "esc":
			if c.sel.IsOpen() {
				break // sel fecha o dropdown, não sai do campo
			}
			c.blurCurrent()
			return c, nil
		}
	}
	var cmd tea.Cmd
	if c.focused == cfgProvider {
		c.sel, cmd = c.sel.Update(msg)
	} else {
		c.inputs[c.focused], cmd = c.inputs[c.focused].Update(msg)
	}
	return c, cmd
}

// FooterHints retorna os atalhos de teclado desta tela para o footer fixo.
func (c ConfigModel) FooterHints() string {
	th, s := c.th, c.s
	return th.Key("tab") + th.DimStyle().Render("próximo campo") + th.Sep() +
		th.Key("enter") + th.DimStyle().Render("salvar") + th.Sep() +
		th.Key("esc") + th.DimStyle().Render("descartar") + th.Sep() +
		th.Key("q") + th.DimStyle().Render(s.FooterQuit)
}

func (c ConfigModel) View() string {
	halfW := c.width / 2
	return lipgloss.JoinHorizontal(lipgloss.Top,
		c.renderLeft(halfW, c.height),
		c.renderRight(halfW, c.height),
	)
}

func (c ConfigModel) renderLeft(w, h int) string {
	th, s := c.th, c.s
	fields := "\n" +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgTwitchCh, cfgTwitchCh), c.inputs[cfgTwitchCh].View()) +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgKickCh, cfgKickCh), c.inputs[cfgKickCh].View()) +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgKickRoom, cfgKickRoom), c.inputs[cfgKickRoom].View()) +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgYouTubeID, cfgYouTubeID), c.inputs[cfgYouTubeID].View()) +
		" " + th.DimStyle().Render(s.CfgRoomHint)

	platformPanel := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.FocusBorder).
		Width(w-2).Height(h-3).
		Padding(0, 1).
		Render(th.TitleStyle().Render(s.CfgPlatforms) + fields)

	hint := " " + th.DimStyle().Render(s.CfgYouTubeHint)
	return lipgloss.JoinVertical(lipgloss.Left, platformPanel, hint)
}

func (c ConfigModel) renderRight(w, h int) string {
	th, s := c.th, c.s
	serverH, aiH := h/2, h-h/2

	serverFields := "\n" +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgPort, cfgPort), c.inputs[cfgPort].View()) +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgHost, cfgHost), c.inputs[cfgHost].View())
	serverPanel := th.PanelStyle(w, serverH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.CfgServer) + serverFields)

	saveBtn := lipgloss.NewStyle().
		Foreground(th.BG).Background(th.Yellow).Bold(true).Padding(0, 1).
		Render(s.CfgSave)
	discardBtn := lipgloss.NewStyle().
		Foreground(th.Dim).Border(lipgloss.NormalBorder()).BorderForeground(th.Border).Padding(0, 1).
		Render(s.CfgDiscard)

	savedMsg := ""
	if c.saved {
		savedMsg = "\n " + toastStyle(th).Render(s.CfgSaved)
	}

	providerLabel := c.fieldLabel(s.CfgProvider, cfgProvider)

	aiFields := "\n" +
		fmt.Sprintf(" %s\n %s\n\n", c.fieldLabel(s.CfgAPIKey, cfgAPIKey), c.inputs[cfgAPIKey].View()) +
		fmt.Sprintf(" %s\n %s\n\n", providerLabel, c.sel.View()) +
		fmt.Sprintf(" %s  %s%s", saveBtn, discardBtn, savedMsg)

	aiPanel := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.Cyan).
		Width(w-2).Height(aiH-2).
		Padding(0, 1).
		Render(th.TitleStyle().Render(s.CfgAI) + aiFields)

	return lipgloss.JoinVertical(lipgloss.Left, serverPanel, aiPanel)
}

func (c *ConfigModel) blurCurrent() {
	if c.focused < cfgFieldCount {
		c.inputs[c.focused].Blur()
	} else {
		c.sel.Blur()
	}
}

func (c *ConfigModel) focusCurrent() {
	if c.focused < cfgFieldCount {
		c.inputs[c.focused].Focus()
	} else {
		c.sel.Focus()
	}
}

func (c ConfigModel) fieldLabel(text string, idx int) string {
	if c.focused == idx {
		return lipgloss.NewStyle().Foreground(c.th.FocusBorder).Bold(true).Render(text)
	}
	return c.th.DimStyle().Render(text)
}

// IsEditing retorna true quando um campo de texto está realmente focado.
func (c ConfigModel) IsEditing() bool {
	if c.focused < cfgFieldCount {
		return c.inputs[c.focused].Focused()
	}
	return c.sel.IsFocused()
}

func (c ConfigModel) GetTwitchChannel() string { return c.inputs[cfgTwitchCh].Value() }
func (c ConfigModel) GetKickChannel() string   { return c.inputs[cfgKickCh].Value() }
func (c ConfigModel) FocusedField() int        { return c.focused }
