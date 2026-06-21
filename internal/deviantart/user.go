package deviantart

import (
	"context"
	"fmt"
	"net/url"
)

// GetProfile retrieves user profile information
func (c *Client) GetProfile(ctx context.Context, username string, extCollections, extGalleries bool) (*UserProfileResponse, error) {
	params := url.Values{}
	if extCollections {
		params.Set("ext_collections", "true")
	}
	if extGalleries {
		params.Set("ext_galleries", "true")
	}

	path := fmt.Sprintf("user/profile/%s", username)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result UserProfileResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateProfile updates the user's profile information
func (c *Client) UpdateProfile(ctx context.Context, userIsArtist *bool, artistLevel, artistSpecialty, realName, tagline string, countryID int, website, bio string) error {
	params := url.Values{}
	if userIsArtist != nil {
		if *userIsArtist {
			params.Set("user_is_artist", "true")
		} else {
			params.Set("user_is_artist", "false")
		}
	}
	if artistLevel != "" {
		params.Set("artist_level", artistLevel)
	}
	if artistSpecialty != "" {
		params.Set("artist_specialty", artistSpecialty)
	}
	if realName != "" {
		params.Set("real_name", realName)
	}
	if tagline != "" {
		params.Set("tagline", tagline)
	}
	if countryID > 0 {
		params.Set("countryid", fmt.Sprintf("%d", countryID))
	}
	if website != "" {
		params.Set("website", website)
	}
	if bio != "" {
		params.Set("bio", bio)
	}

	resp, err := c.post(ctx, "user/profile/update", params)
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
		return fmt.Errorf("failed to update profile")
	}

	return nil
}

// GetFriends retrieves the user's list of friends
func (c *Client) GetFriends(ctx context.Context, username string, offset, limit int) (*PaginatedResponse[UserFriend], error) {
	params := buildPaginationParams(offset, limit)

	path := fmt.Sprintf("user/friends/%s", username)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[UserFriend]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// WhoIs fetches user info for given usernames
func (c *Client) WhoIs(ctx context.Context, usernames []string) ([]User, error) {
	if len(usernames) == 0 {
		return nil, fmt.Errorf("at least one username is required")
	}

	params := url.Values{}
	for _, username := range usernames {
		params.Add("usernames", username)
	}

	resp, err := c.post(ctx, "user/whois", params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Results []User `json:"results"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Results, nil
}

// WhoAmI fetches user info of authenticated user
func (c *Client) WhoAmI(ctx context.Context) (*User, error) {
	resp, err := c.get(ctx, "user/whoami", nil)
	if err != nil {
		return nil, err
	}

	var result User
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// SearchFriends searches friends by username
func (c *Client) SearchFriends(ctx context.Context, query, username string, offset, limit int) (*PaginatedResponse[UserFriend], error) {
	if query == "" {
		return nil, fmt.Errorf("query is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("query", query)
	if username != "" {
		params.Set("username", username)
	}

	resp, err := c.get(ctx, "user/friends/search", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[UserFriend]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetStatuses retrieves user statuses
func (c *Client) GetStatuses(ctx context.Context, username string, matureContent bool, offset, limit int) (*PaginatedResponse[Status], error) {
	params := buildPaginationParams(offset, limit)
	params.Set("username", username)
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "user/statuses/", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Status]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetStatus fetches the status
func (c *Client) GetStatus(ctx context.Context, statusID string, matureContent bool) (*Status, error) {
	if statusID == "" {
		return nil, fmt.Errorf("status ID is required")
	}

	params := url.Values{}
	if matureContent {
		params.Set("mature_content", "true")
	}

	path := fmt.Sprintf("user/statuses/%s", statusID)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result Status
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetWatchers retrieves the user's list of watchers
func (c *Client) GetWatchers(ctx context.Context, username string, offset, limit int) (*PaginatedResponse[UserFriend], error) {
	params := buildPaginationParams(offset, limit)

	path := fmt.Sprintf("user/watchers/%s", username)
	resp, err := c.get(ctx, path, params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[UserFriend]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// WatchStatus checks if user is being watched by the given user
func (c *Client) WatchStatus(ctx context.Context, username string) (bool, error) {
	path := fmt.Sprintf("user/friends/watching/%s", username)
	resp, err := c.get(ctx, path, nil)
	if err != nil {
		return false, err
	}

	var result struct {
		IsWatching bool `json:"is_watching"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return false, err
	}

	return result.IsWatching, nil
}

// Watch watches a user
func (c *Client) Watch(ctx context.Context, username string, watch WatchOptions) error {
	if username == "" {
		return fmt.Errorf("username is required")
	}

	params := url.Values{}
	watchParams := url.Values{}
	if watch.Friend {
		watchParams.Set("friend", "true")
	}
	if watch.Deviations {
		watchParams.Set("deviations", "true")
	}
	if watch.Journals {
		watchParams.Set("journals", "true")
	}
	if watch.ForumThreads {
		watchParams.Set("forum_threads", "true")
	}
	if watch.Critiques {
		watchParams.Set("critiques", "true")
	}
	if watch.Scraps {
		watchParams.Set("scraps", "true")
	}
	if watch.Activity {
		watchParams.Set("activity", "true")
	}
	if watch.Collections {
		watchParams.Set("collections", "true")
	}
	params.Set("watch", watchParams.Encode())

	path := fmt.Sprintf("user/friends/watch/%s", username)
	resp, err := c.post(ctx, path, params)
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
		return fmt.Errorf("failed to watch user")
	}

	return nil
}

// Unwatch unwatches a user
func (c *Client) Unwatch(ctx context.Context, username string) error {
	if username == "" {
		return fmt.Errorf("username is required")
	}

	path := fmt.Sprintf("user/friends/unwatch/%s", username)
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
		return fmt.Errorf("failed to unwatch user")
	}

	return nil
}

// DamnToken retrieves the dAmn auth token required to connect to the dAmn servers
func (c *Client) DamnToken(ctx context.Context) (string, error) {
	resp, err := c.get(ctx, "user/damntoken", nil)
	if err != nil {
		return "", err
	}

	var result struct {
		Token string `json:"token"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return "", err
	}

	return result.Token, nil
}

// WatchOptions specifies watch options
type WatchOptions struct {
	Friend       bool
	Deviations   bool
	Journals     bool
	ForumThreads bool
	Critiques    bool
	Scraps       bool
	Activity     bool
	Collections  bool
}
