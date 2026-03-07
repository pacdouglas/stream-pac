package platform

import (
	"context"
	"time"
)

// Message is the unified chat message across all platforms.
type Message struct {
	Platform    string
	Author      string
	AuthorColor string    // hex color; empty if the platform does not provide it
	Text        string
	Timestamp   time.Time
	Raw         any       // original platform payload, useful for debugging
}

// Status is the current connection state of a platform.
type Status struct {
	Connected bool
	Viewers   int
	Err       error
}

// Credentials provides authentication tokens to a platform.
// Implemented by auth.StaticToken, auth.OAuthSession, auth.BrowserSession.
// Duck typing: no import of the auth package required here.
type Credentials interface {
	Valid() bool
	NeedsRefresh() bool
	AccessToken() string
}

// Platform is the interface every streaming platform must implement.
type Platform interface {
	// Name returns the lowercase identifier: "twitch", "kick", "youtube".
	Name() string
	// Connect opens the connection to the platform's chat stream.
	Connect(ctx context.Context) error
	// Disconnect closes the connection gracefully.
	Disconnect() error
	// Messages returns a read-only channel of incoming chat messages.
	Messages() <-chan Message
	// Status returns the current connection status.
	Status() Status
}
