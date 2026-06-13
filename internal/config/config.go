package config

// Config holds environment-based configuration
type Config struct {
	// TODO: Add client ID, secret, token storage path fields
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// TODO: Implement env var parsing and validation
	return &Config{}, nil
}
