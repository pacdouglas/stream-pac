package kick

import (
	"context"

	"github.com/pacdouglas/stream-pac/internal/platform"
)

func init() {
	platform.Register("kick", func(cfg map[string]any, creds platform.Credentials) (platform.Platform, error) {
		c := Config{}
		if id, ok := cfg["chatroom_id"].(string); ok {
			c.ChatroomID = id
		}
		return New(c, creds)
	})
}

// PusherConfig holds the Kick Pusher WebSocket connection settings.
var PusherConfig = struct {
	Key     string
	Cluster string
}{
	Key:     "32cbd69e4b950bf97679",
	Cluster: "us2",
}

// Config holds Kick-specific configuration.
// ChatroomID must be obtained via browser (Cloudflare blocks direct Go requests).
type Config struct {
	ChatroomID string // Pusher channel: "chatrooms.{ChatroomID}.v2"
}

// Kick connects to Kick chat via Pusher WebSocket.
// Auth uses BrowserSession (cookies) due to Cloudflare protection.
type Kick struct {
	cfg    Config
	creds  platform.Credentials // expected: *auth.BrowserSession
	msgs   chan platform.Message
	status platform.Status
}

// New creates a new Kick platform instance.
func New(cfg Config, creds platform.Credentials) (*Kick, error) {
	panic("not implemented")
}

func (k *Kick) Name() string                      { return "kick" }
func (k *Kick) Connect(ctx context.Context) error { panic("not implemented") }
func (k *Kick) Disconnect() error                 { panic("not implemented") }
func (k *Kick) Messages() <-chan platform.Message  { panic("not implemented") }
func (k *Kick) Status() platform.Status           { panic("not implemented") }

var _ platform.Platform = (*Kick)(nil)
