package platform

import (
	"context"
	"fmt"
	"sync"
)

// Factory creates a Platform from a generic config map and credentials.
type Factory func(cfg map[string]any, creds Credentials) (Platform, error)

var (
	mu       sync.RWMutex
	registry = map[string]Factory{}
)

// Register adds a platform factory to the global registry.
// Each platform package calls this from its init() function.
func Register(name string, fn Factory) {
	mu.Lock()
	defer mu.Unlock()
	registry[name] = fn
}

// New creates a Platform by name using the registered factory.
func New(name string, cfg map[string]any, creds Credentials) (Platform, error) {
	mu.RLock()
	fn, ok := registry[name]
	mu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("platform %q not registered", name)
	}
	return fn(cfg, creds)
}

// Registered returns the names of all registered platforms.
func Registered() []string {
	mu.RLock()
	defer mu.RUnlock()
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}

// ConnectAll connects a slice of platforms concurrently.
// Returns the first error encountered but attempts all connections.
func ConnectAll(ctx context.Context, platforms []Platform) error {
	panic("not implemented")
}
