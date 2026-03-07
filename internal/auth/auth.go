package auth

import "context"

// OAuthConfig holds the parameters for an OAuth2 PKCE flow.
type OAuthConfig struct {
	ClientID     string
	ClientSecret string // empty for pure PKCE (public client)
	AuthURL      string
	TokenURL     string
	RedirectPort int      // callback server listens at localhost:PORT
	Scopes       []string
	Platform     string
}

// Authorizer performs OAuth2 PKCE flows and token refresh.
type Authorizer interface {
	// Authorize opens the system browser and waits for the OAuth callback.
	Authorize(ctx context.Context) (*OAuthSession, error)
	// Refresh exchanges a refresh token for a new access token.
	Refresh(ctx context.Context, session *OAuthSession) (*OAuthSession, error)
}

// NewOAuthAuthorizer creates an Authorizer for the given OAuth config.
func NewOAuthAuthorizer(cfg OAuthConfig) Authorizer {
	panic("not implemented")
}
