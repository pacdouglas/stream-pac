package overlay

import (
	"context"
	"time"

	"github.com/pacdouglas/stream-pac/internal/event"
)

// Meta holds the metadata for an overlay stored on disk.
// Serialized as overlay.json alongside the overlay's index.html.
type Meta struct {
	Name        string
	Description string
	Events      []event.Type // which bus events this overlay subscribes to
	Version     int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Manager manages overlays on the filesystem.
type Manager interface {
	// List returns metadata for all overlays.
	List(ctx context.Context) ([]Meta, error)
	// Get returns the metadata and HTML content of the named overlay.
	Get(ctx context.Context, name string) (*Meta, []byte, error)
	// Save creates or updates an overlay (HTML content + metadata).
	Save(ctx context.Context, name string, html []byte, meta Meta) error
	// Delete removes the overlay directory.
	Delete(ctx context.Context, name string) error
}
