package auth

import (
	"golang.org/x/oauth2"
)

// TokenStore handles file-based token persistence
type TokenStore struct {
	// TODO: Add storage path field
}

// NewTokenStore creates a new token store
func NewTokenStore(path string) *TokenStore {
	// TODO: Implement token store initialization
	return &TokenStore{}
}

// Load loads a token from storage
func (ts *TokenStore) Load() (*oauth2.Token, error) {
	// TODO: Implement token loading
	return nil, nil
}

// Save saves a token to storage
func (ts *TokenStore) Save(token *oauth2.Token) error {
	// TODO: Implement token saving
	return nil
}
