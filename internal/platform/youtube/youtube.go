package youtube

import (
	"context"
	"time"

	"github.com/pacdouglas/stream-pac/internal/platform"
)

func init() {
	platform.Register("youtube", func(cfg map[string]any, creds platform.Credentials) (platform.Platform, error) {
		c := Config{PollInterval: 5 * time.Second}
		if id, ok := cfg["video_id"].(string); ok {
			c.VideoID = id
		}
		if pi, ok := cfg["poll_interval"].(string); ok {
			if d, err := time.ParseDuration(pi); err == nil {
				c.PollInterval = d
			}
		}
		return New(c, creds)
	})
}

// OAuthEndpoints holds the YouTube / Google OAuth2 URLs.
// Required scopes: https://www.googleapis.com/auth/youtube.readonly
var OAuthEndpoints = struct {
	AuthURL  string
	TokenURL string
}{
	AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
	TokenURL: "https://oauth2.googleapis.com/token",
}

// Config holds YouTube-specific configuration.
type Config struct {
	VideoID      string
	PollInterval time.Duration // default: 5s
}

// YouTube polls YouTube live chat via the internal youtubei/v1/live_chat API.
// Endpoint: https://www.youtube.com/youtubei/v1/live_chat/get_live_chat
type YouTube struct {
	cfg    Config
	creds  platform.Credentials
	msgs   chan platform.Message
	status platform.Status
}

// New creates a new YouTube platform instance.
func New(cfg Config, creds platform.Credentials) (*YouTube, error) {
	panic("not implemented")
}

func (y *YouTube) Name() string                      { return "youtube" }
func (y *YouTube) Connect(ctx context.Context) error { panic("not implemented") }
func (y *YouTube) Disconnect() error                 { panic("not implemented") }
func (y *YouTube) Messages() <-chan platform.Message  { panic("not implemented") }
func (y *YouTube) Status() platform.Status           { panic("not implemented") }

var _ platform.Platform = (*YouTube)(nil)
