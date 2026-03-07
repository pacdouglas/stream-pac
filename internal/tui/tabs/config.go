package tabs

import tea "github.com/charmbracelet/bubbletea"

// Config allows editing platform credentials, server settings, and UI preferences.
// TODO: implement using real config store.
type Config struct {
	client Client
	width  int
	height int
}

func NewConfig(client Client) *Config { return &Config{client: client} }

func (c *Config) Title() string                           { return "Config" }
func (c *Config) Init() tea.Cmd                           { return nil } // TODO
func (c *Config) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return c, nil } // TODO
func (c *Config) View() string                            { return "" } // TODO

var _ Tab = (*Config)(nil)
