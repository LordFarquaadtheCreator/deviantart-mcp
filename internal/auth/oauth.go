package auth

import (
	"golang.org/x/oauth2"
)

// OAuthConfig wraps oauth2.Config for DeviantArt
type OAuthConfig struct {
	// TODO: Add oauth2.Config field
}

// NewOAuthConfig creates a new OAuth configuration for DeviantArt
func NewOAuthConfig(clientID, clientSecret string) *OAuthConfig {
	// TODO: Implement oauth2.Config setup
	return &OAuthConfig{}
}
