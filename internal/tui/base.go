package tui

// baseTab contém os campos e comportamentos comuns a todas as telas.
// Toda tela embeds baseTab — equivale a herança de campos em linguagens OO.
//
// COMPOSIÇÃO GO vs HERANÇA
// ════════════════════════
// Go não tem herança, mas tem embedding: os campos e métodos de baseTab
// ficam disponíveis diretamente na struct que a embeds.
// handleBase() centraliza o tratamento de WindowSize/Theme/Lang —
// cada sub-modelo chama d.baseTab = d.baseTab.handleBase(msg) para atualizar
// os campos base, e depois faz seu próprio tratamento adicional.
//
// O model raiz distribui bodyHeight (não terminal height) para os sub-modelos,
// garantindo que header e footer sejam sempre fixos e visíveis.

import tea "github.com/charmbracelet/bubbletea"

type baseTab struct {
	width  int
	height int // bodyHeight — já descontado header+footer pelo model raiz
	th     Theme
	s      Strings
}

func newBaseTab(width, height int, th Theme, s Strings) baseTab {
	return baseTab{width: width, height: height, th: th, s: s}
}

// handleBase processa as mensagens que toda tela precisa tratar:
// resize, troca de tema e troca de idioma.
// Retorna o baseTab atualizado (value semantics, imutável).
// Sub-modelos chamam isso primeiro e depois fazem seu tratamento adicional.
func (b baseTab) handleBase(msg tea.Msg) baseTab {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.width, b.height = msg.Width, msg.Height
	case themeMsg:
		b.th = msg.theme
	case langMsg:
		b.s = msg.strings
	}
	return b
}
