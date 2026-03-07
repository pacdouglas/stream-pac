package tabs

import tea "github.com/charmbracelet/bubbletea"

// AIGen is the AI overlay generator screen.
// TODO: implement real Anthropic API integration.
type AIGen struct {
	client Client
	width  int
	height int
	prompt string
}

func NewAIGen(client Client) *AIGen { return &AIGen{client: client} }

func (a *AIGen) Title() string                           { return "IA" }
func (a *AIGen) Init() tea.Cmd                           { return nil }          // TODO
func (a *AIGen) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return a, nil }       // TODO
func (a *AIGen) View() string                            { return "" }            // TODO

var _ Tab = (*AIGen)(nil)
