package deviantart

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// GetDeviation retrieves a deviation by ID
func (c *Client) GetDeviation(ctx context.Context, deviationID string) (*Deviation, error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}

	path := fmt.Sprintf("deviation/%s", deviationID)
	resp, err := c.get(ctx, path, nil)
	if err != nil {
		return nil, err
	}

	var result Deviation
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDeviationContent retrieves deviation content
func (c *Client) GetDeviationContent(ctx context.Context, deviationID string, matureContent bool) (*DeviationContent, error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}

	params := url.Values{}
	params.Set("deviationid", deviationID)
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "deviation/content", params)
	if err != nil {
		return nil, err
	}

	var result DeviationContent
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDeviationDownload retrieves deviation download information
func (c *Client) GetDeviationDownload(ctx context.Context, deviationID string, matureContent bool) (*DeviationDownload, error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}

	params := url.Values{}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("deviation/download/%s", deviationID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result DeviationDownload
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDeviationEmbeddedContent retrieves content embedded in a deviation
func (c *Client) GetDeviationEmbeddedContent(ctx context.Context, deviationID string, offsetDeviationID string, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("deviationid", deviationID)
	if offsetDeviationID != "" {
		params.Set("offset_deviationid", offsetDeviationID)
	}

	resp, err := c.get(ctx, "deviation/embeddedcontent", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetDeviationMetadata retrieves deviation metadata for a set of deviations
func (c *Client) GetDeviationMetadata(ctx context.Context, deviationIDs []string, extSubmission, extCamera, extStats, extCollection bool) ([]DeviationMetadata, error) {
	if len(deviationIDs) == 0 {
		return nil, fmt.Errorf("at least one deviation ID is required")
	}
	if len(deviationIDs) > 50 {
		return nil, fmt.Errorf("maximum 50 deviation IDs allowed")
	}

	params := url.Values{}
	for _, id := range deviationIDs {
		params.Add("deviationids", id)
	}
	if extSubmission {
		params.Set("ext_submission", "true")
	}
	if extCamera {
		params.Set("ext_camera", "true")
	}
	if extStats {
		params.Set("ext_stats", "true")
	}
	if extCollection {
		params.Set("ext_collection", "true")
	}

	resp, err := c.get(ctx, "deviation/metadata", params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Metadata []DeviationMetadata `json:"metadata"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Metadata, nil
}

// DeviationWhoFaved retrieves users who favorited a deviation
func (c *Client) DeviationWhoFaved(ctx context.Context, deviationID string, offset, limit int) (*PaginatedResponse[User], error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("deviationid", deviationID)

	resp, err := c.get(ctx, "deviation/whofaved", params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pagination
		Results []struct {
			User User     `json:"user"`
			Time time.Time `json:"time"`
		} `json:"results"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	// Convert to simpler response
	users := make([]User, len(result.Results))
	for i, r := range result.Results {
		users[i] = r.User
	}

	return &PaginatedResponse[User]{
		Pagination: result.Pagination,
		Results:    users,
	}, nil
}
