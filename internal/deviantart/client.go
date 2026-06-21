package deviantart

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
)

const (
	baseURL         = "https://www.deviantart.com/api/v1/oauth2/"
	minorVersion    = "20240701"
	defaultTimeout  = 30 * time.Second
	defaultLimit    = 10
	maxLimit        = 50
)

// Client is the DeviantArt API client
type Client struct {
	httpClient   *http.Client
	tokenSource oauth2.TokenSource
	baseURL     string
}

// NewClient creates a new DeviantArt API client
func NewClient(tokenSource oauth2.TokenSource) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
		tokenSource: tokenSource,
		baseURL:     baseURL,
	}
}

// doRequest performs an HTTP request with OAuth authentication
func (c *Client) doRequest(ctx context.Context, method, path string, params url.Values, body io.Reader) (*http.Response, error) {
	token, err := c.tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	fullURL := c.baseURL + path
	if params != nil && len(params) > 0 {
		fullURL += "?" + params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("dA-minor-version", minorVersion)

	if body != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return c.httpClient.Do(req)
}

// get performs a GET request
func (c *Client) get(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	return c.doRequest(ctx, http.MethodGet, path, params, nil)
}

// post performs a POST request with form data
func (c *Client) post(ctx context.Context, path string, params url.Values) (*http.Response, error) {
	body := params.Encode()
	return c.doRequest(ctx, http.MethodPost, path, nil, bytes.NewBufferString(body))
}

// decodeResponse decodes a JSON response into the target struct
func decodeResponse(resp *http.Response, target any) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("HTTP %d: failed to decode error response: %w", resp.StatusCode, err)
		}
		return fmt.Errorf("HTTP %d: %s - %s", resp.StatusCode, errResp.Error, errResp.ErrorDescription)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

// buildPaginationParams builds pagination parameters
func buildPaginationParams(offset, limit int) url.Values {
	params := url.Values{}
	if offset > 0 {
		params.Set("offset", fmt.Sprintf("%d", offset))
	}
	if limit > 0 {
		if limit > maxLimit {
			limit = maxLimit
		}
		params.Set("limit", fmt.Sprintf("%d", limit))
	} else {
		params.Set("limit", fmt.Sprintf("%d", defaultLimit))
	}
	return params
}
