package tui

// AIGenModel é a tela [4] — gerador de overlays via IA.
// Recebe bodyHeight do model raiz via baseTab — não conhece header/footer.

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AIGenModel struct {
	baseTab            // width, height(body), th, s via composição
	prompt     textarea.Model
	progress   string
	result     string
	generating bool
}

// genStepMsg e genDoneMsg simulam streaming de geração.
type genStepMsg struct{ text string }
type genDoneMsg struct{ result string }

func newAIGen(width, height int, th Theme, s Strings) AIGenModel {
	ta := textarea.New()
	ta.Placeholder = "Chat lateral direita, fundo escuro semi-transparente..."
	ta.SetValue("Chat lateral direita, fundo escuro semi-transparente, usernames em neon amarelo brilhante, fonte monospace, animação slide-left na entrada das mensagens, máximo 8 visíveis, emotes grandes")
	// Sem auto-focus — usuário usa Tab ou clique para entrar no campo.
	ta.CharLimit = 0
	ta.SetWidth(width/2 - 4)
	ta.SetHeight(6)

	ta.FocusedStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.FocusBorder)
	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.Border)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle().Background(th.Panel)

	return AIGenModel{baseTab: newBaseTab(width, height, th, s), prompt: ta}
}

func (ai AIGenModel) Update(msg tea.Msg) (AIGenModel, tea.Cmd) {
	ai.baseTab = ai.baseTab.handleBase(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		ai.prompt.SetWidth(ai.width/2 - 4)
		return ai, nil
	case themeMsg:
		ai.prompt.FocusedStyle.Base = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).BorderForeground(ai.th.FocusBorder)
		ai.prompt.BlurredStyle.Base = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).BorderForeground(ai.th.Border)
		return ai, nil
	case langMsg:
		ai.prompt.Placeholder = ai.s.AIWaiting
		return ai, nil
	case genStepMsg:
		ai.progress += msg.text + "\n"
		return ai, nil
	case genDoneMsg:
		ai.result = msg.result
		ai.generating = false
		return ai, nil
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlJ || (msg.Type == tea.KeyEnter && msg.Alt) {
			if !ai.generating {
				ai.generating = true
				ai.progress = ""
				ai.result = ai.s.AIGenerating
				return ai, ai.startGeneration()
			}
			return ai, nil
		}
	}
	var cmd tea.Cmd
	ai.prompt, cmd = ai.prompt.Update(msg)
	return ai, cmd
}

// startGeneration simula geração de overlay em background.
func (ai AIGenModel) startGeneration() tea.Cmd {
	return func() tea.Msg {
		steps := []struct{ delay time.Duration }{
			{0},
			{500 * time.Millisecond},
			{800 * time.Millisecond},
			{600 * time.Millisecond},
			{600 * time.Millisecond},
			{500 * time.Millisecond},
		}
		for _, step := range steps {
			time.Sleep(step.delay)
		}
		return genDoneMsg{result: "✓ arcade_chat.html  (4.2 kb)\n\nURL\nhttp://localhost:7777/arcade_chat.html\n\nConecta em /events automaticamente.\nCompatível com o hub SSE.\n\n[s] salvar  [2] ver em Overlays"}
	}
}

// FooterHints retorna os atalhos de teclado desta tela para o footer fixo.
// IsEditing retorna true quando o textarea está realmente focado.
func (ai AIGenModel) IsEditing() bool { return ai.prompt.Focused() }

func (ai AIGenModel) FooterHints() string {
	th, s := ai.th, ai.s
	return th.Key("ctrl+enter") + th.DimStyle().Render("gerar") + th.Sep() +
		th.Key("tab") + th.DimStyle().Render("próximo campo") + th.Sep() +
		th.Key("q") + th.DimStyle().Render(s.FooterQuit)
}

func (ai AIGenModel) View() string {
	th, s := ai.th, ai.s
	halfW := ai.width / 2
	optH := ai.height - 12

	promptTitle := th.TitleStyle().Render(s.AIPromptTitle) +
		"  " + th.DimStyle().Render(s.AIPromptHint)

	typeContent := "\n" +
		" " + th.DimStyle().Render("☐") + " " + th.TextStyle().Render(s.AIChat) + "\n" +
		" " + lipgloss.NewStyle().Foreground(th.Yellow).Render("☑") + " " + th.TextStyle().Render(s.AIWebcam) + "\n" +
		" " + th.DimStyle().Render("☐") + " " + th.TextStyle().Render(s.AIAlert) + "\n" +
		" " + th.DimStyle().Render("☐") + " " + th.TextStyle().Render(s.AIViewers) + "\n" +
		" " + th.DimStyle().Render("☐") + " " + th.TextStyle().Render(s.AIHypeTrain)

	promptPanel := th.FocusedPanelStyle(halfW, 12).Padding(0, 1).
		Render(promptTitle + "\n\n" + ai.prompt.View())
	typePanel := th.PanelStyle(halfW, 12).Padding(0, 1).
		Render(th.TitleStyle().Render(s.AIType) + typeContent)

	topRow := lipgloss.JoinHorizontal(lipgloss.Top, promptPanel, typePanel)

	genBtn := lipgloss.NewStyle().
		Foreground(th.BG).Background(th.Yellow).Bold(true).Padding(0, 1).
		Render(" " + s.AIGenerate + " ")

	optContent := "\n" +
		fmt.Sprintf(" %s  %s\n\n", th.DimStyle().Render(s.AITheme),
			lipgloss.NewStyle().Foreground(th.Text).Render("arcade")) +
		fmt.Sprintf(" %s  %s\n\n", th.DimStyle().Render(s.AIPalette),
			lipgloss.NewStyle().Foreground(th.Text).Render("amarelo/azul")) +
		fmt.Sprintf(" %s  %s\n", th.DimStyle().Render(s.AIWidth),
			th.TextStyle().Render("400")) +
		fmt.Sprintf(" %s %s\n\n", th.DimStyle().Render(s.AIHeight),
			th.TextStyle().Render("1080")) +
		" " + genBtn

	progressContent := "\n"
	if ai.progress == "" && !ai.generating {
		progressContent += " " + th.DimStyle().Render(s.AIWaiting) + "\n" +
			" " + th.DimStyle().Render(s.AIWaitingHint)
	} else {
		progressContent += lipgloss.NewStyle().Foreground(th.Cyan).Render(ai.progress)
	}

	resultContent := "\n"
	if ai.result == "" {
		resultContent += " " + th.DimStyle().Render("Nenhum overlay gerado ainda.")
	} else if ai.generating {
		resultContent += " " + th.DimStyle().Render(ai.result)
	} else {
		resultContent += toastStyle(th).Render(ai.result)
	}

	if optH < 4 {
		optH = 4
	}
	optPanel := th.PanelStyle(halfW/2, optH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.AIOptions) + optContent)
	statusPanel := th.PanelStyle(halfW/2, optH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.AIStatus) + progressContent)
	resultPanel := th.PanelStyle(ai.width-halfW, optH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.AIResult) + resultContent)

	bottomRow := lipgloss.JoinHorizontal(lipgloss.Top, optPanel, statusPanel, resultPanel)

	return lipgloss.JoinVertical(lipgloss.Left, topRow, bottomRow)
}
