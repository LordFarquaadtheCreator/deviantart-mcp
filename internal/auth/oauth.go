package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"

	"golang.org/x/oauth2"
)

const (
	authURL  = "https://www.deviantart.com/oauth2/authorize"
	tokenURL = "https://www.deviantart.com/oauth2/token"
)

// AllScopes is the complete set of DeviantArt OAuth scopes requested during login.
var AllScopes = []string{
	"basic",
	"browse",
	"collection",
	"comment.post",
	"feed",
	"message",
	"note",
	"stash",
	"user",
	"user.manage",
}

// OAuthConfig wraps oauth2.Config for DeviantArt
type OAuthConfig struct {
	config *oauth2.Config
}

// NewOAuthConfig creates a new OAuth configuration for DeviantArt.
// Scopes are set to basic+browse for client-credentials usage.
func NewOAuthConfig(clientID, clientSecret, redirectURI string) *OAuthConfig {
	return &OAuthConfig{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURI,
			Scopes:       []string{"basic", "browse"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
		},
	}
}

// NewUserOAuthConfig creates an OAuth configuration with all user scopes,
// used for the authorization-code flow during auth login.
func NewUserOAuthConfig(clientID, clientSecret, redirectURI string) *OAuthConfig {
	return &OAuthConfig{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURI,
			Scopes:       AllScopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  authURL,
				TokenURL: tokenURL,
			},
		},
	}
}

// Config returns the underlying oauth2.Config
func (oc *OAuthConfig) Config() *oauth2.Config {
	return oc.config
}

// GeneratePKCEVerifier creates a cryptographically random code verifier
// (43 bytes → 56 base64url chars, no padding, well within the 43-128 range).
func GeneratePKCEVerifier() (string, error) {
	b := make([]byte, 43)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// GeneratePKCEChallenge returns the S256 code challenge for a given verifier.
func GeneratePKCEChallenge(verifier string) string {
	h := sha256.Sum256([]byte(verifier))
	return base64.RawURLEncoding.EncodeToString(h[:])
}

// GenerateState creates a random CSRF state token.
func GenerateState() (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
