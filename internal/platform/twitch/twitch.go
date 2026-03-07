package twitch

import (
	"context"

	"github.com/pacdouglas/stream-pac/internal/platform"
)

func init() {
	platform.Register("twitch", func(cfg map[string]any, creds platform.Credentials) (platform.Platform, error) {
		c := Config{}
		if ch, ok := cfg["channel"].(string); ok {
			c.Channel = ch
		}
		return New(c, creds)
	})
}

// OAuthEndpoints holds the Twitch OAuth2 URLs.
// Required scopes: chat:read, user:read:email
var OAuthEndpoints = struct {
	AuthURL  string
	TokenURL string
}{
	AuthURL:  "https://id.twitch.tv/oauth2/authorize",
	TokenURL: "https://id.twitch.tv/oauth2/token",
}

// Config holds Twitch-specific configuration.
type Config struct {
	Channel string
}

// Twitch connects to Twitch IRC over WebSocket (irc.chat.twitch.tv:6697).
type Twitch struct {
	cfg    Config
	creds  platform.Credentials
	msgs   chan platform.Message
	status platform.Status
}

// New creates a new Twitch platform instance.
func New(cfg Config, creds platform.Credentials) (*Twitch, error) {
	panic("not implemented")
}

func (t *Twitch) Name() string                      { return "twitch" }
func (t *Twitch) Connect(ctx context.Context) error { panic("not implemented") }
func (t *Twitch) Disconnect() error                 { panic("not implemented") }
func (t *Twitch) Messages() <-chan platform.Message  { panic("not implemented") }
func (t *Twitch) Status() platform.Status           { panic("not implemented") }

var _ platform.Platform = (*Twitch)(nil)
