package store

import (
	"context"
	"time"
)

// Session is a stored OAuth or browser session for a platform.
// AccessToken, RefreshToken, and RawSession are stored encrypted in the database.
type Session struct {
	Platform     string
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
	Scope        string
	RawSession   []byte // Kick: serialized cookie map
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// SaveSession inserts or replaces the session for session.Platform.
func (s *Store) SaveSession(ctx context.Context, session Session) error {
	panic("not implemented")
}

// GetSession retrieves and decrypts the session for the given platform.
// Returns (nil, nil) if no session exists.
func (s *Store) GetSession(ctx context.Context, platform string) (*Session, error) {
	panic("not implemented")
}

// DeleteSession removes the session for the given platform.
func (s *Store) DeleteSession(ctx context.Context, platform string) error {
	panic("not implemented")
}

// ListSessions returns metadata for all sessions (tokens are not decrypted).
func (s *Store) ListSessions(ctx context.Context) ([]Session, error) {
	panic("not implemented")
}
