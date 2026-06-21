package deviantart

import (
	"context"
	"fmt"
	"net/url"
)

// GetCollectionsFolders retrieves collection folders
func (c *Client) GetCollectionsFolders(ctx context.Context, username string, calculateSize, extPreload bool, offset, limit int) (*PaginatedResponse[Collection], error) {
	params := buildPaginationParams(offset, limit)
	if username != "" {
		params.Set("username", username)
	}
	if calculateSize {
		params.Set("calculate_size", "true")
	}
	if extPreload {
		params.Set("ext_preload", "true")
	}

	resp, err := c.get(ctx, "collections/folders", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Collection]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetCollections retrieves collection folder contents
func (c *Client) GetCollections(ctx context.Context, folderID, username string, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if folderID == "" {
		return nil, fmt.Errorf("folder ID is required")
	}

	params := buildPaginationParams(offset, limit)
	if username != "" {
		params.Set("username", username)
	}

	path := fmt.Sprintf("collections/%s", folderID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Fave adds deviation to favourites
func (c *Client) Fave(ctx context.Context, deviationID, folderID string) error {
	if deviationID == "" {
		return fmt.Errorf("deviation ID is required")
	}

	params := url.Values{}
	params.Set("deviationid", deviationID)
	if folderID != "" {
		params.Set("folderid", folderID)
	}

	resp, err := c.post(ctx, "collections/fave", params)
	if err != nil {
		return err
	}

	var result struct {
		Success bool `json:"success"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return err
	}

	if !result.Success {
		return fmt.Errorf("failed to favorite deviation")
	}

	return nil
}

// Unfave removes deviation from favourites
func (c *Client) Unfave(ctx context.Context, deviationID, folderID string) error {
	if deviationID == "" {
		return fmt.Errorf("deviation ID is required")
	}

	params := url.Values{}
	params.Set("deviationid", deviationID)
	if folderID != "" {
		params.Set("folderid", folderID)
	}

	resp, err := c.post(ctx, "collections/unfave", params)
	if err != nil {
		return err
	}

	var result struct {
		Success bool `json:"success"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return err
	}

	if !result.Success {
		return fmt.Errorf("failed to unfavorite deviation")
	}

	return nil
}

// CreateCollectionFolder creates a new collection folder
func (c *Client) CreateCollectionFolder(ctx context.Context, name string) (*Collection, error) {
	if name == "" {
		return nil, fmt.Errorf("folder name is required")
	}

	params := url.Values{}
	params.Set("folder", name)

	resp, err := c.post(ctx, "collections/folders/create", params)
	if err != nil {
		return nil, err
	}

	var result Collection
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RemoveCollectionFolder deletes a collection folder
func (c *Client) RemoveCollectionFolder(ctx context.Context, folderID string) error {
	if folderID == "" {
		return fmt.Errorf("folder ID is required")
	}

	path := fmt.Sprintf("collections/folders/remove/%s", folderID)
	resp, err := c.get(ctx, path, nil)
	if err != nil {
		return err
	}

	var result struct {
		Success bool `json:"success"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return err
	}

	if !result.Success {
		return fmt.Errorf("failed to remove collection folder")
	}

	return nil
}
