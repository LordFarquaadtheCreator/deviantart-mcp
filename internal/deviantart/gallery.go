package deviantart

import (
	"context"
	"fmt"
	"net/url"
)

// GetGalleryAll retrieves the "all" view of a user's gallery
func (c *Client) GetGalleryAll(ctx context.Context, username string, offset, limit int) (*PaginatedResponse[Deviation], error) {
	params := buildPaginationParams(offset, limit)
	if username != "" {
		params.Set("username", username)
	}

	resp, err := c.get(ctx, "gallery/all", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGalleryFolders retrieves gallery folders
func (c *Client) GetGalleryFolders(ctx context.Context, username string, calculateSize, extPreload bool, offset, limit int) (*PaginatedResponse[GalleryFolder], error) {
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

	resp, err := c.get(ctx, "gallery/folders", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[GalleryFolder]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetGallery retrieves gallery folder contents
func (c *Client) GetGallery(ctx context.Context, username, folderID, mode string, offset, limit int) (*PaginatedResponse[Deviation], error) {
	params := buildPaginationParams(offset, limit)
	if username != "" {
		params.Set("username", username)
	}
	if mode != "" {
		params.Set("mode", mode)
	}

	var path string
	if folderID != "" {
		path = fmt.Sprintf("gallery/%s", folderID)
	} else {
		path = "gallery/"
	}

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

// CreateGalleryFolder creates a new gallery folder
func (c *Client) CreateGalleryFolder(ctx context.Context, folderName string) (*GalleryFolder, error) {
	if folderName == "" {
		return nil, fmt.Errorf("folder name is required")
	}

	params := url.Values{}
	params.Set("folder", folderName)

	resp, err := c.post(ctx, "gallery/folders/create", params)
	if err != nil {
		return nil, err
	}

	var result GalleryFolder
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RemoveGalleryFolder deletes a gallery folder
func (c *Client) RemoveGalleryFolder(ctx context.Context, folderID string) error {
	if folderID == "" {
		return fmt.Errorf("folder ID is required")
	}

	path := fmt.Sprintf("gallery/folders/remove/%s", folderID)
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
		return fmt.Errorf("failed to remove gallery folder")
	}

	return nil
}
