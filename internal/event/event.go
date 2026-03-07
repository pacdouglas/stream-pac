package event

import (
	"sync"
	"time"
)

// Type identifies a specific kind of event.
type Type string

const (
	TypeMessageReceived      Type = "message.received"
	TypePlatformConnected    Type = "platform.connected"
	TypePlatformDisconnected Type = "platform.disconnected"
	TypeViewerCountChanged   Type = "viewer.count.changed"
	TypeOverlayReloaded      Type = "overlay.reloaded"
	TypeOverlayGenerated     Type = "overlay.generated"
)

// Event carries a typed payload through the bus.
type Event struct {
	Type    Type
	Payload any
	At      time.Time
}

// Handler is a callback invoked on each matching event.
type Handler func(Event)

// Bus is an in-process publish/subscribe event bus.
type Bus interface {
	Publish(e Event)
	Subscribe(t Type, h Handler) (unsubscribe func())
	SubscribeAll(h Handler) (unsubscribe func())
}

type handlerEntry struct {
	id      uint64
	handler Handler
}

// InMemoryBus is a synchronous, in-memory Bus implementation.
type InMemoryBus struct {
	mu          sync.RWMutex
	subscribers map[Type][]handlerEntry
	all         []handlerEntry
	seq         uint64
}

// NewBus returns a new InMemoryBus.
func NewBus() Bus {
	return &InMemoryBus{
		subscribers: make(map[Type][]handlerEntry),
	}
}

func (b *InMemoryBus) Publish(e Event)                     { panic("not implemented") }
func (b *InMemoryBus) Subscribe(t Type, h Handler) func()  { panic("not implemented") }
func (b *InMemoryBus) SubscribeAll(h Handler) func()       { panic("not implemented") }

var _ Bus = (*InMemoryBus)(nil)
