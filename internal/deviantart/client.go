package deviantart

import (
	"golang.org/x/oauth2"
)

// Client is the DeviantArt API client
type Client struct {
	// TODO: Add HTTP client, base URL, token source fields
}

// NewClient creates a new DeviantArt API client
func NewClient(tokenSource oauth2.TokenSource) *Client {
	// TODO: Implement client initialization
	return &Client{}
}
