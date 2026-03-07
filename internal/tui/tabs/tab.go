package tabs

import (
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

// Tab is a Bubble Tea model with a display title.
// TODO: implement when integrating real HTTP client logic.
type Tab interface {
	tea.Model
	Title() string
}

// Client is a shared HTTP client passed to all tabs.
// TODO: wire up to real HTTP server when backend is implemented.
type Client struct {
	HTTP    *http.Client
	BaseURL string // e.g. "http://127.0.0.1:7777"
}
