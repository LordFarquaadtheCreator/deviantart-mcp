package deviantart

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// BrowseDailyDeviations retrieves daily deviations for a specific date
func (c *Client) BrowseDailyDeviations(ctx context.Context, date time.Time, geo string, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	params := buildPaginationParams(offset, limit)
	if !date.IsZero() {
		params.Set("date", date.Format("2006-01-02"))
	}
	if geo != "" {
		params.Set("geo", geo)
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/dailydeviations", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseHotTopics retrieves hot topics
func (c *Client) BrowseHotTopics(ctx context.Context) ([]Topic, error) {
	resp, err := c.get(ctx, "browse/hottopics", nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		TopicList []Topic `json:"topic_list"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.TopicList, nil
}

// BrowseJournals retrieves journals
func (c *Client) BrowseJournals(ctx context.Context, featured, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	params := buildPaginationParams(offset, limit)
	if featured {
		params.Set("featured", "true")
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/journals", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseMoreLikeThis retrieves deviations similar to a seed deviation
func (c *Client) BrowseMoreLikeThis(ctx context.Context, seed string, category string, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if seed == "" {
		return nil, fmt.Errorf("seed deviation ID is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("seed", seed)
	if category != "" {
		params.Set("category", category)
	}

	resp, err := c.get(ctx, "browse/morelikethis", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseNewest retrieves newest deviations
func (c *Client) BrowseNewest(ctx context.Context, categoryPath, query string, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	params := buildPaginationParams(offset, limit)
	if categoryPath != "" {
		params.Set("category_path", categoryPath)
	}
	if query != "" {
		params.Set("q", query)
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/newest", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowsePopular retrieves popular deviations
func (c *Client) BrowsePopular(ctx context.Context, categoryPath, query, timeRange string, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	params := buildPaginationParams(offset, limit)
	if categoryPath != "" {
		params.Set("category_path", categoryPath)
	}
	if query != "" {
		params.Set("q", query)
	}
	if timeRange != "" {
		params.Set("timerange", timeRange)
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/popular", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseTags retrieves deviations by tag
func (c *Client) BrowseTags(ctx context.Context, tag string, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if tag == "" {
		return nil, fmt.Errorf("tag is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("tag", tag)
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/tags", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseTopic retrieves deviations for a specific topic
func (c *Client) BrowseTopic(ctx context.Context, topicID string, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if topicID == "" {
		return nil, fmt.Errorf("topic ID is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("topic_id", topicID)
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/topic", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseCategoryTree retrieves category tree
func (c *Client) BrowseCategoryTree(ctx context.Context, categoryPath string, depth int) ([]Category, error) {
	params := url.Values{}
	if categoryPath != "" {
		params.Set("category_path", categoryPath)
	}
	if depth > 0 {
		params.Set("depth", fmt.Sprintf("%d", depth))
	}

	resp, err := c.get(ctx, "browse/categorytree", params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Categories []Category `json:"categories"`
	}
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Categories, nil
}

// BrowseUserFriends retrieves a user's friends
func (c *Client) BrowseUserFriends(ctx context.Context, username string, offset, limit int) (*PaginatedResponse[UserFriend], error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("username", username)

	resp, err := c.get(ctx, "browse/user/friends", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[UserFriend]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseUserJournals retrieves a user's journals
func (c *Client) BrowseUserJournals(ctx context.Context, username string, featured, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("username", username)
	if featured {
		params.Set("featured", "true")
	}
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/user/journals", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BrowseUserLiterature retrieves a user's literature
func (c *Client) BrowseUserLiterature(ctx context.Context, username string, matureContent bool, offset, limit int) (*PaginatedResponse[Deviation], error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	params := buildPaginationParams(offset, limit)
	params.Set("username", username)
	if matureContent {
		params.Set("mature_content", "true")
	}

	resp, err := c.get(ctx, "browse/user/literature", params)
	if err != nil {
		return nil, err
	}

	var result PaginatedResponse[Deviation]
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
