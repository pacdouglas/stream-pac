package tabs

import tea "github.com/charmbracelet/bubbletea"

// Logs streams application logs from GET /sse/logs and displays them
// in a scrollable viewport with level-based coloring.
// TODO: implement real SSE log streaming.
type Logs struct {
	client Client
	width  int
	height int
	lines  []string
}

func NewLogs(client Client) *Logs { return &Logs{client: client} }

func (l *Logs) Title() string                           { return "Logs" }
func (l *Logs) Init() tea.Cmd                           { return nil }    // TODO
func (l *Logs) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return l, nil } // TODO
func (l *Logs) View() string                            { return "" }     // TODO

var _ Tab = (*Logs)(nil)
