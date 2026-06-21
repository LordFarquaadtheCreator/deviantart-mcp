package deviantart

import (
	"context"
	"fmt"
	"net/url"
)

// CommentsPost posts a comment on a deviation
func (c *Client) CommentsPost(ctx context.Context, deviationID, body, commentID, requestID string) (*Comment, error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}
	if body == "" {
		return nil, fmt.Errorf("comment body is required")
	}

	params := url.Values{}
	params.Set("body", body)
	if commentID != "" {
		params.Set("commentid", commentID)
	}
	if requestID != "" {
		params.Set("request_id", requestID)
	}

	path := fmt.Sprintf("comments/post/%s", deviationID)
	resp, err := c.post(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result Comment
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CommentsGetDeviationComments retrieves comments for a deviation
func (c *Client) CommentsGetDeviationComments(ctx context.Context, deviationID, commentID string, maxDepth int, matureContent bool, offset, limit int) (*PaginatedResponse[Comment], error) {
	if deviationID == "" {
		return nil, fmt.Errorf("deviation ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if commentID != "" {
		params.Set("commentid", commentID)
	}
	if maxDepth > 0 {
		params.Set("maxdepth", fmt.Sprintf("%d", maxDepth))
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("comments/%s", deviationID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pagination
		Thread []Comment `json:"thread"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &PaginatedResponse[Comment]{
		Pagination: result.Pagination,
		Results:    result.Thread,
	}, nil
}

// CommentsGetProfileComments retrieves comments on a user's profile
func (c *Client) CommentsGetProfileComments(ctx context.Context, userID, commentID string, maxDepth int, matureContent bool, offset, limit int) (*PaginatedResponse[Comment], error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if commentID != "" {
		params.Set("commentid", commentID)
	}
	if maxDepth > 0 {
		params.Set("maxdepth", fmt.Sprintf("%d", maxDepth))
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("comments/profile/%s", userID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pagination
		Thread []Comment `json:"thread"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &PaginatedResponse[Comment]{
		Pagination: result.Pagination,
		Results:    result.Thread,
	}, nil
}

// CommentsPostProfileComment posts a comment on a user's profile
func (c *Client) CommentsPostProfileComment(ctx context.Context, userID, body, commentID, requestID string) (*Comment, error) {
	if userID == "" {
		return nil, fmt.Errorf("user ID is required")
	}
	if body == "" {
		return nil, fmt.Errorf("comment body is required")
	}

	params := url.Values{}
	params.Set("body", body)
	if commentID != "" {
		params.Set("commentid", commentID)
	}
	if requestID != "" {
		params.Set("request_id", requestID)
	}

	path := fmt.Sprintf("comments/post/profile/%s", userID)
	resp, err := c.post(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result Comment
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CommentsGetStatusComments retrieves comments on a status
func (c *Client) CommentsGetStatusComments(ctx context.Context, statusID, commentID string, maxDepth int, matureContent bool, offset, limit int) (*PaginatedResponse[Comment], error) {
	if statusID == "" {
		return nil, fmt.Errorf("status ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if commentID != "" {
		params.Set("commentid", commentID)
	}
	if maxDepth > 0 {
		params.Set("maxdepth", fmt.Sprintf("%d", maxDepth))
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("comments/status/%s", statusID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Pagination
		Thread []Comment `json:"thread"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &PaginatedResponse[Comment]{
		Pagination: result.Pagination,
		Results:    result.Thread,
	}, nil
}

// CommentsPostStatusComment posts a comment on a status
func (c *Client) CommentsPostStatusComment(ctx context.Context, statusID, body, commentID, requestID string) (*Comment, error) {
	if statusID == "" {
		return nil, fmt.Errorf("status ID is required")
	}
	if body == "" {
		return nil, fmt.Errorf("comment body is required")
	}

	params := url.Values{}
	params.Set("body", body)
	if commentID != "" {
		params.Set("commentid", commentID)
	}
	if requestID != "" {
		params.Set("request_id", requestID)
	}

	path := fmt.Sprintf("comments/post/status/%s", statusID)
	resp, err := c.post(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result Comment
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
