package tabs

import tea "github.com/charmbracelet/bubbletea"

// Dashboard shows connected platforms, viewer counts, and live message rates.
// TODO: implement using real SSE data from the hub.
type Dashboard struct {
	client Client
	width  int
	height int
}

func NewDashboard(client Client) *Dashboard { return &Dashboard{client: client} }

func (d *Dashboard) Title() string                           { return "Dashboard" }
func (d *Dashboard) Init() tea.Cmd                           { return nil } // TODO
func (d *Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return d, nil } // TODO
func (d *Dashboard) View() string                            { return "" } // TODO

var _ Tab = (*Dashboard)(nil)
