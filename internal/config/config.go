package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Config holds environment-based configuration for the DeviantArt MCP server.
type Config struct {
	ClientID      string
	ClientSecret  string
	RedirectURI   string
	TokenPath     string
	MatureContent bool
	LogLevel      string
}

// Load reads configuration from environment variables.
// DA_CLIENT_ID and DA_CLIENT_SECRET are required.
func Load() (*Config, error) {
	cfg := &Config{
		ClientID:     os.Getenv("DA_CLIENT_ID"),
		ClientSecret: os.Getenv("DA_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("DA_REDIRECT_URI"),
		TokenPath:    os.Getenv("DA_TOKEN_PATH"),
		LogLevel:     os.Getenv("DA_LOG_LEVEL"),
	}

	if cfg.ClientID == "" {
		return nil, fmt.Errorf("DA_CLIENT_ID environment variable is required")
	}
	if cfg.ClientSecret == "" {
		return nil, fmt.Errorf("DA_CLIENT_SECRET environment variable is required")
	}

	if cfg.RedirectURI == "" {
		cfg.RedirectURI = "http://localhost:9004/callback"
	}

	if cfg.TokenPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, fmt.Errorf("cannot determine home directory for token path: %w", err)
		}
		cfg.TokenPath = filepath.Join(home, ".config", "deviantart-mcp", "token.json")
	}

	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	if os.Getenv("DA_MATURE_CONTENT") == "true" {
		cfg.MatureContent = true
	}

	return cfg, nil
}
