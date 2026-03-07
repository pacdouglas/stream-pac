package auth

import "time"

// Credentials provides authentication to a platform.
// All three concrete types satisfy platform.Credentials via duck typing.
type Credentials interface {
	Valid() bool
	NeedsRefresh() bool
	AccessToken() string
}

// StaticToken is a manually configured token (no OAuth, sourced from config file).
type StaticToken struct {
	Token string
}

func (t *StaticToken) Valid() bool         { return t.Token != "" }
func (t *StaticToken) NeedsRefresh() bool  { return false }
func (t *StaticToken) AccessToken() string { return t.Token }

// OAuthSession holds a full OAuth2 session with refresh capability.
// Used for Twitch, YouTube, Facebook, and any standard OAuth2 platform.
type OAuthSession struct {
	Access    string
	Refresh   string
	ExpiresAt time.Time
	Scope     []string
	Platform  string
}

func (s *OAuthSession) Valid() bool { return s.Access != "" && time.Now().Before(s.ExpiresAt) }
func (s *OAuthSession) NeedsRefresh() bool {
	return time.Now().After(s.ExpiresAt.Add(-5 * time.Minute))
}
func (s *OAuthSession) AccessToken() string { return s.Access }

// BrowserSession holds cookies/session data for platforms behind Cloudflare (Kick).
// Populated by capturing a browser session rather than a standard OAuth flow.
type BrowserSession struct {
	Cookies   map[string]string
	UserAgent string
	ExpiresAt time.Time
}

func (s *BrowserSession) Valid() bool { return len(s.Cookies) > 0 && time.Now().Before(s.ExpiresAt) }
func (s *BrowserSession) NeedsRefresh() bool {
	return time.Now().After(s.ExpiresAt.Add(-10 * time.Minute))
}
func (s *BrowserSession) AccessToken() string { return "" } // not applicable for Kick
