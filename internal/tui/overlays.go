package tui

// OverlaysModel é a tela [2].
// Recebe bodyHeight do model raiz via baseTab — não conhece header/footer.

import (
	"fmt"
	"io"

	"github.com/pacdouglas/stream-pac/internal/tui/fake"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type overlayItem struct {
	overlay fake.Overlay
	isAI    bool
}

func (i overlayItem) FilterValue() string { return i.overlay.Name }

type overlayDelegate struct {
	th Theme
	s  Strings
}

func (d overlayDelegate) Height() int                              { return 2 }
func (d overlayDelegate) Spacing() int                            { return 1 }
func (d overlayDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d overlayDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	ov, ok := item.(overlayItem)
	if !ok {
		return
	}
	th, s := d.th, d.s
	isSelected := index == m.Index()

	var title, desc string
	if ov.isAI {
		title = lipgloss.NewStyle().Foreground(th.Cyan).Bold(true).Render(s.OvAIItem)
		desc = th.DimStyle().Render("   " + s.OvAIDesc)
	} else {
		status := lipgloss.NewStyle().Foreground(th.Green).Render(s.OvActive)
		if !ov.overlay.Active {
			status = th.DimStyle().Render(s.OvInactive)
		}
		title = fmt.Sprintf("%s  [%s]", ov.overlay.Name, status)
		desc = th.DimStyle().Render(fmt.Sprintf("   %d×%d  ·  localhost:7777/%s",
			ov.overlay.Width, ov.overlay.Height, ov.overlay.File))
	}

	if isSelected {
		title = lipgloss.NewStyle().
			Foreground(th.Text).Bold(true).
			BorderLeft(true).
			BorderStyle(lipgloss.ThickBorder()).
			BorderForeground(th.FocusBorder).
			PaddingLeft(1).
			Render(title)
		desc = lipgloss.NewStyle().Foreground(th.Muted).PaddingLeft(3).Render(desc)
	} else {
		title = lipgloss.NewStyle().Foreground(th.Text).PaddingLeft(2).Render(title)
		desc = lipgloss.NewStyle().Foreground(th.Dim).PaddingLeft(2).Render(desc)
	}

	fmt.Fprintf(w, "%s\n%s", title, desc)
}

// ── OverlaysModel ─────────────────────────────────────────────────────────────

type OverlaysModel struct {
	baseTab       // width, height(body), th, s via composição
	list  list.Model
	toast string
}

func newOverlays(width, height int, th Theme, s Strings) OverlaysModel {
	items := buildOverlayItems()
	delegate := overlayDelegate{th: th, s: s}
	listW, listH := width/2-2, height-2
	l := list.New(items, delegate, listW, listH)
	l.Title = s.OvTitle
	l.Styles.Title = lipgloss.NewStyle().Foreground(th.Title).Bold(true)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.FilterPrompt = lipgloss.NewStyle().Foreground(th.Blue)
	l.Styles.FilterCursor = lipgloss.NewStyle().Foreground(th.Yellow)

	return OverlaysModel{baseTab: newBaseTab(width, height, th, s), list: l}
}

func buildOverlayItems() []list.Item {
	items := make([]list.Item, 0, len(fake.Overlays)+1)
	for _, ov := range fake.Overlays {
		items = append(items, overlayItem{overlay: ov})
	}
	items = append(items, overlayItem{isAI: true})
	return items
}

func (o OverlaysModel) Update(msg tea.Msg) (OverlaysModel, tea.Cmd) {
	o.baseTab = o.baseTab.handleBase(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		o.list.SetSize(o.width/2-2, o.height-2)
		return o, nil
	case themeMsg:
		o.list.Styles.Title = lipgloss.NewStyle().Foreground(o.th.Title).Bold(true)
		o.list.SetDelegate(overlayDelegate{th: o.th, s: o.s})
		return o, nil
	case langMsg:
		o.list.Title = o.s.OvTitle
		o.list.SetDelegate(overlayDelegate{th: o.th, s: o.s})
		return o, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if sel, ok := o.list.SelectedItem().(overlayItem); ok {
				if sel.isAI {
					return o, func() tea.Msg { return switchTabMsg(TabAI) }
				}
				o.toast = o.s.OvCopied
				return o, func() tea.Msg { return notifyMsg(o.s.OvCopied) }
			}
			return o, nil
		case "o":
			o.toast = o.s.OvOpening
			return o, func() tea.Msg { return notifyMsg(o.s.OvOpening) }
		case "delete", "backspace":
			o.toast = o.s.OvRemoved
			return o, func() tea.Msg { return notifyMsg(o.s.OvRemoved) }
		}
	}
	var cmd tea.Cmd
	o.list, cmd = o.list.Update(msg)
	return o, cmd
}

// FooterHints retorna os atalhos de teclado desta tela para o footer fixo.
func (o OverlaysModel) FooterHints() string {
	th, s := o.th, o.s
	return th.Key("enter") + th.DimStyle().Render("copiar URL") + th.Sep() +
		th.Key("o") + th.DimStyle().Render("browser") + th.Sep() +
		th.Key("del") + th.DimStyle().Render("remover") + th.Sep() +
		th.Key("/") + th.DimStyle().Render(s.FooterFilter) + th.Sep() +
		th.Key("q") + th.DimStyle().Render(s.FooterQuit)
}

func (o OverlaysModel) View() string {
	halfW := o.width / 2
	listPanel := o.th.PanelStyle(halfW, o.height).Render(o.list.View())
	rightPanel := o.renderRight(halfW, o.height)
	return lipgloss.JoinHorizontal(lipgloss.Top, listPanel, rightPanel)
}

func (o OverlaysModel) renderRight(w, h int) string {
	th, s := o.th, o.s
	sel, ok := o.list.SelectedItem().(overlayItem)
	previewH := h - 10
	if previewH < 4 {
		previewH = 4
	}
	obsH := 8

	var previewContent, obsContent string

	if !ok || sel.isAI {
		previewContent = "\n\n" +
			"  " + lipgloss.NewStyle().Foreground(th.Cyan).Bold(true).Render("✦") +
			" " + th.BoldText().Render(s.OvAIItem) + "\n\n" +
			"  " + th.TextStyle().Render("Descreva o overlay que você quer") + "\n" +
			"  " + th.TextStyle().Render("e o Claude vai gerar o HTML") + "\n" +
			"  " + th.TextStyle().Render("pronto para usar no OBS.") + "\n\n" +
			"  " + th.DimStyle().Render(s.OvAIHint)
	} else {
		ov := sel.overlay
		previewContent = "\n " + th.DimStyle().Render(s.OvSimulation) + "\n\n" +
			"  ┌────────────────────────────┐\n" +
			"  │ " + lipgloss.NewStyle().Foreground(th.Yellow).Bold(true).Render("STREAMPAC") +
			"  " + lipgloss.NewStyle().Foreground(th.Purple).Render("●") +
			lipgloss.NewStyle().Foreground(th.Green).Render("●") + "          │\n" +
			"  │ " + th.DimStyle().Render("viewers: 142") + "                  │\n" +
			"  │ " + th.DimStyle().Render("──────────────────────────────") + "│\n" +
			"  │ " + lipgloss.NewStyle().Foreground(th.Purple).Bold(true).Render("[TW]") +
			" Pac: " + th.TextStyle().Render("salve galera!!") + "    │\n" +
			"  │ " + lipgloss.NewStyle().Foreground(th.Green).Bold(true).Render("[KI]") +
			" fan_top: " + th.TextStyle().Render("oi pessoal") + "    │\n" +
			"  └────────────────────────────┘"

		obsContent = "\n" +
			" " + th.DimStyle().Render(s.OvURL) + "\n" +
			" " + lipgloss.NewStyle().Foreground(th.Blue).Render("http://localhost:7777/"+ov.File) + "\n\n" +
			fmt.Sprintf(" %s  %s    %s  %s\n",
				th.DimStyle().Render(s.OvWidth), th.TextStyle().Render(fmt.Sprint(ov.Width)),
				th.DimStyle().Render(s.OvHeight), th.TextStyle().Render(fmt.Sprint(ov.Height)),
			) +
			"\n " + th.DimStyle().Render(s.OvCopyHint)
	}

	if o.toast != "" {
		obsContent = "\n " + toastStyle(th).Render("✓ "+o.toast)
	}

	previewPanel := th.PanelStyle(w, previewH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.OvPreview) + "\n" + previewContent)
	obsPanel := th.PanelStyle(w, obsH).Padding(0, 1).
		Render(th.TitleStyle().Render(s.OvOBSSource) + "\n" + obsContent)

	return lipgloss.JoinVertical(lipgloss.Left, previewPanel, obsPanel)
}

func (o OverlaysModel) ItemCount() int { return len(o.list.Items()) }

func (o OverlaysModel) SelectedOverlayFile() string {
	if sel, ok := o.list.SelectedItem().(overlayItem); ok && !sel.isAI {
		return sel.overlay.File
	}
	return ""
}
