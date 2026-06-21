package deviantart

import (
	"context"
)

// GetCountries retrieves a list of countries
func (c *Client) GetCountries(ctx context.Context) (*DataCountries, error) {
	resp, err := c.get(ctx, "data/countries", nil)
	if err != nil {
		return nil, err
	}

	var result DataCountries
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetPrivacy retrieves the DeviantArt Privacy Policy
func (c *Client) GetPrivacy(ctx context.Context) (*DataText, error) {
	resp, err := c.get(ctx, "data/privacy", nil)
	if err != nil {
		return nil, err
	}

	var result DataText
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetSubmission retrieves the DeviantArt Submission Policy
func (c *Client) GetSubmission(ctx context.Context) (*DataText, error) {
	resp, err := c.get(ctx, "data/submission", nil)
	if err != nil {
		return nil, err
	}

	var result DataText
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetTOS retrieves the DeviantArt Terms of Service
func (c *Client) GetTOS(ctx context.Context) (*DataText, error) {
	resp, err := c.get(ctx, "data/tos", nil)
	if err != nil {
		return nil, err
	}

	var result DataText
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
