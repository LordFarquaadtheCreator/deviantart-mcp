package auth

import (
	"encoding/json"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
)

// TokenStore handles file-based token persistence
type TokenStore struct {
	storagePath string
}

// NewTokenStore creates a new token store
func NewTokenStore(path string) *TokenStore {
	return &TokenStore{
		storagePath: path,
	}
}

// Load loads a token from storage
func (ts *TokenStore) Load() (*oauth2.Token, error) {
	data, err := os.ReadFile(ts.storagePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // No token file exists
		}
		return nil, err
	}

	var token oauth2.Token
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, err
	}

	return &token, nil
}

// Save saves a token to storage
func (ts *TokenStore) Save(token *oauth2.Token) error {
	if err := os.MkdirAll(filepath.Dir(ts.storagePath), 0755); err != nil {
		return err
	}

	data, err := json.Marshal(token)
	if err != nil {
		return err
	}

	return os.WriteFile(ts.storagePath, data, 0600)
}
