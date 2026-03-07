package overlay

import (
	"context"
	"path/filepath"
)

// FileManager is the filesystem-backed implementation of Manager.
// Each overlay lives under overlaysDir/{name}/:
//
//	index.html   — the overlay HTML/CSS/JS
//	overlay.json — serialized Meta
type FileManager struct {
	overlaysDir string
}

// NewFileManager creates a FileManager rooted at overlaysDir.
func NewFileManager(overlaysDir string) *FileManager {
	return &FileManager{overlaysDir: overlaysDir}
}

func (m *FileManager) dir(name string) string {
	return filepath.Join(m.overlaysDir, name)
}

func (m *FileManager) List(ctx context.Context) ([]Meta, error)                            { panic("not implemented") }
func (m *FileManager) Get(ctx context.Context, name string) (*Meta, []byte, error)         { panic("not implemented") }
func (m *FileManager) Save(ctx context.Context, name string, html []byte, meta Meta) error { panic("not implemented") }
func (m *FileManager) Delete(ctx context.Context, name string) error                       { panic("not implemented") }

var _ Manager = (*FileManager)(nil)
