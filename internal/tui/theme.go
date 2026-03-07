package tui

import "github.com/charmbracelet/lipgloss"

// Theme agrupa todas as cores da aplicação.
// Trocar de tema = trocar de Theme struct — sem nenhuma lógica condicional espalhada.
type Theme struct {
	Name string

	// Superfícies
	BG    lipgloss.Color // fundo da janela
	Panel lipgloss.Color // fundo dos painéis

	// Bordas
	Border      lipgloss.Color // borda padrão
	FocusBorder lipgloss.Color // borda do painel com foco

	// Texto
	Text  lipgloss.Color // texto principal
	Dim   lipgloss.Color // texto secundário / dica
	Muted lipgloss.Color // texto ainda mais apagado

	// Identidade / acento
	Title  lipgloss.Color // títulos dos painéis
	Blue   lipgloss.Color // cor de ação primária (maze blue)
	Purple lipgloss.Color // Twitch / IA
	Green  lipgloss.Color // Kick / sucesso
	Red    lipgloss.Color // YouTube / erro
	Yellow lipgloss.Color // Pac-Man / destaque
	Cyan   lipgloss.Color // power pellet / info
	Pink   lipgloss.Color // Pinky ghost / log ia
}

// Dark é o tema padrão — arcade night (StreamPac).
var Dark = Theme{
	Name:        "dark",
	BG:          lipgloss.Color("#0A0E27"),
	Panel:       lipgloss.Color("#0D1440"),
	Border:      lipgloss.Color("#1E3A5F"),
	FocusBorder: lipgloss.Color("#FFD700"),
	Text:        lipgloss.Color("#E6EDF3"),
	Dim:         lipgloss.Color("#4A5080"),
	Muted:       lipgloss.Color("#6875A8"),
	Title:       lipgloss.Color("#FFD700"),
	Blue:        lipgloss.Color("#1E90FF"),
	Purple:      lipgloss.Color("#9b72ff"),
	Green:       lipgloss.Color("#53fc18"),
	Red:         lipgloss.Color("#ff4444"),
	Yellow:      lipgloss.Color("#FFD700"),
	Cyan:        lipgloss.Color("#39d0d8"),
	Pink:        lipgloss.Color("#FF69B4"),
}

// ── Temas adicionais ──────────────────────────────────────────────────────────

// Midnight — GitHub Dark: azul frio, quase preto.
var Midnight = Theme{
	Name:        "midnight",
	BG:          lipgloss.Color("#0D1117"),
	Panel:       lipgloss.Color("#161B22"),
	Border:      lipgloss.Color("#30363D"),
	FocusBorder: lipgloss.Color("#58A6FF"),
	Text:        lipgloss.Color("#E6EDF3"),
	Dim:         lipgloss.Color("#484F58"),
	Muted:       lipgloss.Color("#6E7681"),
	Title:       lipgloss.Color("#58A6FF"),
	Blue:        lipgloss.Color("#58A6FF"),
	Purple:      lipgloss.Color("#BC8CFF"),
	Green:       lipgloss.Color("#3FB950"),
	Red:         lipgloss.Color("#F85149"),
	Yellow:      lipgloss.Color("#E3B341"),
	Cyan:        lipgloss.Color("#39C5CF"),
	Pink:        lipgloss.Color("#F778BA"),
}

var Dracula = Theme{
	Name:        "dracula",
	BG:          lipgloss.Color("#282a36"),
	Panel:       lipgloss.Color("#313444"),
	Border:      lipgloss.Color("#44475a"),
	FocusBorder: lipgloss.Color("#ff79c6"),
	Text:        lipgloss.Color("#f8f8f2"),
	Dim:         lipgloss.Color("#6272a4"),
	Muted:       lipgloss.Color("#4a4d65"),
	Title:       lipgloss.Color("#bd93f9"),
	Blue:        lipgloss.Color("#6272a4"),
	Purple:      lipgloss.Color("#bd93f9"),
	Green:       lipgloss.Color("#50fa7b"),
	Red:         lipgloss.Color("#ff5555"),
	Yellow:      lipgloss.Color("#f1fa8c"),
	Cyan:        lipgloss.Color("#8be9fd"),
	Pink:        lipgloss.Color("#ff79c6"),
}

var Nord = Theme{
	Name:        "nord",
	BG:          lipgloss.Color("#2e3440"),
	Panel:       lipgloss.Color("#3b4252"),
	Border:      lipgloss.Color("#434c5e"),
	FocusBorder: lipgloss.Color("#88c0d0"),
	Text:        lipgloss.Color("#eceff4"),
	Dim:         lipgloss.Color("#4c566a"),
	Muted:       lipgloss.Color("#616e88"),
	Title:       lipgloss.Color("#88c0d0"),
	Blue:        lipgloss.Color("#81a1c1"),
	Purple:      lipgloss.Color("#b48ead"),
	Green:       lipgloss.Color("#a3be8c"),
	Red:         lipgloss.Color("#bf616a"),
	Yellow:      lipgloss.Color("#ebcb8b"),
	Cyan:        lipgloss.Color("#88c0d0"),
	Pink:        lipgloss.Color("#b48ead"),
}

var Tokyo = Theme{
	Name:        "tokyo",
	BG:          lipgloss.Color("#1a1b26"),
	Panel:       lipgloss.Color("#24283b"),
	Border:      lipgloss.Color("#414868"),
	FocusBorder: lipgloss.Color("#7aa2f7"),
	Text:        lipgloss.Color("#c0caf5"),
	Dim:         lipgloss.Color("#414868"),
	Muted:       lipgloss.Color("#565f89"),
	Title:       lipgloss.Color("#7aa2f7"),
	Blue:        lipgloss.Color("#7aa2f7"),
	Purple:      lipgloss.Color("#9d7cd8"),
	Green:       lipgloss.Color("#9ece6a"),
	Red:         lipgloss.Color("#f7768e"),
	Yellow:      lipgloss.Color("#e0af68"),
	Cyan:        lipgloss.Color("#7dcfff"),
	Pink:        lipgloss.Color("#bb9af7"),
}

var Gruvbox = Theme{
	Name:        "gruvbox",
	BG:          lipgloss.Color("#282828"),
	Panel:       lipgloss.Color("#3c3836"),
	Border:      lipgloss.Color("#504945"),
	FocusBorder: lipgloss.Color("#d79921"),
	Text:        lipgloss.Color("#ebdbb2"),
	Dim:         lipgloss.Color("#665c54"),
	Muted:       lipgloss.Color("#7c6f64"),
	Title:       lipgloss.Color("#d79921"),
	Blue:        lipgloss.Color("#458588"),
	Purple:      lipgloss.Color("#b16286"),
	Green:       lipgloss.Color("#98971a"),
	Red:         lipgloss.Color("#cc241d"),
	Yellow:      lipgloss.Color("#d79921"),
	Cyan:        lipgloss.Color("#689d6a"),
	Pink:        lipgloss.Color("#d3869b"),
}

var Catppuccin = Theme{
	Name:        "catppuccin",
	BG:          lipgloss.Color("#1e1e2e"),
	Panel:       lipgloss.Color("#181825"),
	Border:      lipgloss.Color("#313244"),
	FocusBorder: lipgloss.Color("#cba6f7"),
	Text:        lipgloss.Color("#cdd6f4"),
	Dim:         lipgloss.Color("#45475a"),
	Muted:       lipgloss.Color("#585b70"),
	Title:       lipgloss.Color("#cba6f7"),
	Blue:        lipgloss.Color("#89b4fa"),
	Purple:      lipgloss.Color("#cba6f7"),
	Green:       lipgloss.Color("#a6e3a1"),
	Red:         lipgloss.Color("#f38ba8"),
	Yellow:      lipgloss.Color("#f9e2af"),
	Cyan:        lipgloss.Color("#89dceb"),
	Pink:        lipgloss.Color("#f5c2e7"),
}

var Solarized = Theme{
	Name:        "solarized",
	BG:          lipgloss.Color("#002b36"),
	Panel:       lipgloss.Color("#073642"),
	Border:      lipgloss.Color("#586e75"),
	FocusBorder: lipgloss.Color("#268bd2"),
	Text:        lipgloss.Color("#93a1a1"),
	Dim:         lipgloss.Color("#586e75"),
	Muted:       lipgloss.Color("#657b83"),
	Title:       lipgloss.Color("#268bd2"),
	Blue:        lipgloss.Color("#268bd2"),
	Purple:      lipgloss.Color("#6c71c4"),
	Green:       lipgloss.Color("#859900"),
	Red:         lipgloss.Color("#dc322f"),
	Yellow:      lipgloss.Color("#b58900"),
	Cyan:        lipgloss.Color("#2aa198"),
	Pink:        lipgloss.Color("#d33682"),
}

var Monokai = Theme{
	Name:        "monokai",
	BG:          lipgloss.Color("#272822"),
	Panel:       lipgloss.Color("#3e3d32"),
	Border:      lipgloss.Color("#75715e"),
	FocusBorder: lipgloss.Color("#a6e22e"),
	Text:        lipgloss.Color("#f8f8f2"),
	Dim:         lipgloss.Color("#75715e"),
	Muted:       lipgloss.Color("#5a5952"),
	Title:       lipgloss.Color("#a6e22e"),
	Blue:        lipgloss.Color("#66d9e8"),
	Purple:      lipgloss.Color("#ae81ff"),
	Green:       lipgloss.Color("#a6e22e"),
	Red:         lipgloss.Color("#f92672"),
	Yellow:      lipgloss.Color("#e6db74"),
	Cyan:        lipgloss.Color("#66d9e8"),
	Pink:        lipgloss.Color("#f92672"),
}

var Matrix = Theme{
	Name:        "matrix",
	BG:          lipgloss.Color("#000000"),
	Panel:       lipgloss.Color("#001200"),
	Border:      lipgloss.Color("#003000"),
	FocusBorder: lipgloss.Color("#00ff41"),
	Text:        lipgloss.Color("#00cc33"),
	Dim:         lipgloss.Color("#005500"),
	Muted:       lipgloss.Color("#007700"),
	Title:       lipgloss.Color("#00ff41"),
	Blue:        lipgloss.Color("#00aaff"),
	Purple:      lipgloss.Color("#aa00ff"),
	Green:       lipgloss.Color("#00ff41"),
	Red:         lipgloss.Color("#ff3300"),
	Yellow:      lipgloss.Color("#ccff00"),
	Cyan:        lipgloss.Color("#00ffcc"),
	Pink:        lipgloss.Color("#ff00aa"),
}

// ThemesByName permite buscar qualquer tema pelo seu Name.
var ThemesByName = map[string]Theme{
	"dark":       Dark,
	"midnight":   Midnight,
	"dracula":    Dracula,
	"nord":       Nord,
	"tokyo":      Tokyo,
	"gruvbox":    Gruvbox,
	"catppuccin": Catppuccin,
	"solarized":  Solarized,
	"monokai":    Monokai,
	"matrix":     Matrix,
}

// ── Helpers de estilo baseados no tema ────────────────────────────────────────

func (th Theme) TitleStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Title).Bold(true)
}

func (th Theme) DimStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Dim)
}

func (th Theme) MutedStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Muted)
}

func (th Theme) TextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Text)
}

func (th Theme) BoldText() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Text).Bold(true)
}

func (th Theme) PanelStyle(w, h int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.Border).
		Width(w - 2).
		Height(h - 2)
}

func (th Theme) FocusedPanelStyle(w, h int) lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(th.FocusBorder).
		Width(w - 2).
		Height(h - 2)
}

func (th Theme) ActiveTabStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(th.BG).
		Background(th.Yellow).
		Bold(true).
		Padding(0, 2)
}

func (th Theme) InactiveTabStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(th.Dim).
		Padding(0, 2)
}

func (th Theme) FooterKeyStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(th.Yellow).Bold(true)
}

func (th Theme) Key(k string) string {
	return th.FooterKeyStyle().Render("["+k+"]") + " "
}

func (th Theme) Sep() string {
	return th.DimStyle().Render("  ·  ")
}
