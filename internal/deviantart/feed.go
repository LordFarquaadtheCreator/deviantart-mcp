package deviantart

import (
	"context"
	"net/url"
)

// FeedHome fetches watch feed
func (c *Client) FeedHome(ctx context.Context, matureContent bool, cursor string) (*FeedHome, error) {
	params := url.Values{}
	if cursor != "" {
		params.Set("cursor", cursor)
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "feed/home", params)
	if err != nil {
		return nil, err
	}

	var result FeedHome
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// FeedProfile fetches profile feed
func (c *Client) FeedProfile(ctx context.Context, cursor string) (*FeedProfile, error) {
	params := url.Values{}
	if cursor != "" {
		params.Set("cursor", cursor)
	}

	resp, err := c.get(ctx, "feed/profile", params)
	if err != nil {
		return nil, err
	}

	var result FeedProfile
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
