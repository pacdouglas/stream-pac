package overlay

import "context"

// Watcher watches the overlays directory for file changes and notifies
// subscribers so OBS Browser Sources can receive a hot-reload SSE event.
type Watcher struct {
	overlaysDir string
}

// NewWatcher creates a Watcher for the given directory.
func NewWatcher(overlaysDir string) *Watcher {
	return &Watcher{overlaysDir: overlaysDir}
}

// Watch starts watching the overlays directory.
// Sends the changed overlay name on the returned channel for each file change.
// Blocks until ctx is cancelled.
// Implementation note: use polling or github.com/fsnotify/fsnotify.
func (w *Watcher) Watch(ctx context.Context) (<-chan string, error) {
	panic("not implemented")
}
