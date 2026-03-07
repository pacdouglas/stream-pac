package store

import (
	"context"
	"database/sql"
	_ "embed"

	"github.com/pacdouglas/stream-pac/internal/crypto"
)

//go:embed migrations/001_init.sql
var initSQL string

// Store is the SQLite-backed persistence layer.
// Sensitive fields (tokens, cookies) are AES-GCM encrypted before insertion.
// Uses modernc.org/sqlite (pure Go, no CGO). Register the driver in main.go:
//
//	import _ "modernc.org/sqlite"
type Store struct {
	db  *sql.DB
	enc *crypto.Encrypter
}

// New opens or creates the SQLite database at dbPath and runs migrations.
func New(ctx context.Context, dbPath string, enc *crypto.Encrypter) (*Store, error) {
	panic("not implemented")
}

// Close closes the underlying database connection.
func (s *Store) Close() error {
	panic("not implemented")
}

// migrate executes the embedded SQL migrations.
func (s *Store) migrate() error {
	panic("not implemented")
}
