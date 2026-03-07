package hub

import (
	"context"
	"sync"

	"github.com/pacdouglas/stream-pac/internal/event"
	"github.com/pacdouglas/stream-pac/internal/platform"
)

// Hub orchestrates multiple platforms.
// It fans-in all platform message channels into a single output channel
// and publishes lifecycle events to the event bus.
type Hub struct {
	mu        sync.RWMutex
	platforms map[string]platform.Platform
	bus       event.Bus
	messages  chan platform.Message // merged output of all connected platforms
}

// New creates a Hub backed by the given event bus.
func New(bus event.Bus) *Hub {
	return &Hub{
		platforms: make(map[string]platform.Platform),
		bus:       bus,
		messages:  make(chan platform.Message, 256),
	}
}

// Add registers a platform with the hub. Must be called before Connect.
func (h *Hub) Add(p platform.Platform) {
	panic("not implemented")
}

// Connect connects the named platform and begins forwarding its messages.
func (h *Hub) Connect(ctx context.Context, name string) error {
	panic("not implemented")
}

// Disconnect disconnects the named platform gracefully.
func (h *Hub) Disconnect(name string) error {
	panic("not implemented")
}

// Platform returns the platform registered under name.
func (h *Hub) Platform(name string) (platform.Platform, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	p, ok := h.platforms[name]
	return p, ok
}

// Platforms returns all registered platforms.
func (h *Hub) Platforms() []platform.Platform {
	panic("not implemented")
}

// Messages returns the fan-in channel of all platform messages.
func (h *Hub) Messages() <-chan platform.Message {
	return h.messages
}

// Statuses returns a snapshot map of name → status for all platforms.
func (h *Hub) Statuses() map[string]platform.Status {
	panic("not implemented")
}
