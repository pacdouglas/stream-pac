package tabs

import tea "github.com/charmbracelet/bubbletea"

// Overlays lists active overlays, allows creating/deleting, and triggers hot-reload.
// TODO: implement using real overlay manager.
type Overlays struct {
	client Client
	width  int
	height int
}

func NewOverlays(client Client) *Overlays { return &Overlays{client: client} }

func (o *Overlays) Title() string                           { return "Overlays" }
func (o *Overlays) Init() tea.Cmd                           { return nil } // TODO
func (o *Overlays) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return o, nil } // TODO
func (o *Overlays) View() string                            { return "" } // TODO

var _ Tab = (*Overlays)(nil)
