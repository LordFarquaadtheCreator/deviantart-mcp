package deviantart

import (
	"context"
	"fmt"
)

// MessagesFeedback retrieves feedback messages for a user
func (c *Client) MessagesFeedback(ctx context.Context, userID string, matureContent bool, offset, limit int) (*PaginatedResponse[Feedback], error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("messages/feedback/%s", userID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Feedback]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// MessagesFeedbackStack retrieves feedback stack for a user
func (c *Client) MessagesFeedbackStack(ctx context.Context, userID, stackID string, matureContent bool, offset, limit int) (*PaginatedResponse[Feedback], error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if stackID != "" {
		params.Set("stackid", stackID)
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("messages/feedback/%s/stack", userID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Feedback]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// MessagesMentions retrieves mention messages for a user
func (c *Client) MessagesMentions(ctx context.Context, userID string, matureContent bool, offset, limit int) (*PaginatedResponse[Mention], error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("messages/mentions/%s", userID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Mention]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// MessagesMentionsStack retrieves mention stack for a user
func (c *Client) MessagesMentionsStack(ctx context.Context, userID, stackID string, matureContent bool, offset, limit int) (*PaginatedResponse[Mention], error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if stackID != "" {
		params.Set("stackid", stackID)
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("messages/mentions/%s/stack", userID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Mention]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
