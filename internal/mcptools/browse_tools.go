package mcptools

import (
	"context"
	"encoding/json"
	"time"

	"github.com/farquaad/deviantart-mcp/internal/deviantart"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func mustJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		return `{"error": "json marshal failed"}`
	}
	return string(b)
}

type browseDailyDeviationsArgs struct {
	Date          string `json:"date,omitempty" jsonschema:"description=Date in YYYY-MM-DD format (defaults to today if omitted)"`
	Geo           string `json:"geo,omitempty" jsonschema:"description=Geographic filter e.g. us"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty" jsonschema:"description=Pagination offset"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseJournalsArgs struct {
	Featured      bool `json:"featured,omitempty"`
	MatureContent bool `json:"mature_content,omitempty"`
	Offset        int  `json:"offset,omitempty"`
	Limit         int  `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseMoreLikeThisArgs struct {
	Seed     string `json:"seed" jsonschema:"description=Deviation UUID to find similar deviations for (required)"`
	Category string `json:"category,omitempty" jsonschema:"description=Category path filter"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseNewestArgs struct {
	CategoryPath  string `json:"category_path,omitempty" jsonschema:"description=Category path e.g. digitalart/drawings"`
	Query         string `json:"query,omitempty" jsonschema:"description=Search query"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browsePopularArgs struct {
	CategoryPath  string `json:"category_path,omitempty" jsonschema:"description=Category path e.g. digitalart/drawings"`
	Query         string `json:"query,omitempty" jsonschema:"description=Search query"`
	Timerange     string `json:"timerange,omitempty" jsonschema:"description=Time range: 8hr, 24hr, 3days, 1week, 1month, alltime"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseTagsArgs struct {
	Tag           string `json:"tag" jsonschema:"description=Tag to search for (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseTopicArgs struct {
	TopicID       string `json:"topic_id" jsonschema:"description=Topic ID from browse_hot_topics (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseCategoryTreeArgs struct {
	CategoryPath string `json:"category_path,omitempty" jsonschema:"description=Starting category path, omit for root"`
	Depth        int    `json:"depth,omitempty" jsonschema:"description=How many levels deep to return"`
}

type browseUserFriendsArgs struct {
	Username string `json:"username" jsonschema:"description=DeviantArt username (required)"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseUserJournalsArgs struct {
	Username      string `json:"username" jsonschema:"description=DeviantArt username (required)"`
	Featured      bool   `json:"featured,omitempty"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type browseUserLiteratureArgs struct {
	Username      string `json:"username" jsonschema:"description=DeviantArt username (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

func registerBrowseTools(s *mcp.Server, c *deviantart.Client) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_dailydeviations",
		Description: "Returns daily deviations for a specific date. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseDailyDeviationsArgs) (*mcp.CallToolResult, any, error) {
		var date time.Time
		if args.Date != "" {
			var err error
			date, err = time.Parse("2006-01-02", args.Date)
			if err != nil {
				date = time.Now()
			}
		}
		results, err := c.BrowseDailyDeviations(ctx, date, args.Geo, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_hot_topics",
		Description: "Returns currently trending/hot topics on DeviantArt. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseHotTopics(ctx)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_journals",
		Description: "Returns DeviantArt journals, optionally filtered to featured only. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseJournalsArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseJournals(ctx, args.Featured, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_more_like_this",
		Description: "Returns deviations similar to a seed deviation. Requires a deviation UUID. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseMoreLikeThisArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseMoreLikeThis(ctx, args.Seed, args.Category, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_newest",
		Description: "Returns the newest deviations on DeviantArt, optionally filtered by category path and search query. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseNewestArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseNewest(ctx, args.CategoryPath, args.Query, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_popular",
		Description: "Returns the most popular deviations on DeviantArt, optionally filtered by category path, search query, and timerange. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browsePopularArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowsePopular(ctx, args.CategoryPath, args.Query, args.Timerange, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_tags",
		Description: "Returns deviations tagged with a specific tag. The tag parameter is required. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseTagsArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseTags(ctx, args.Tag, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_topic",
		Description: "Returns deviations for a specific topic. Requires a topic_id from browse_hot_topics. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseTopicArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseTopic(ctx, args.TopicID, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_category_tree",
		Description: "Returns the DeviantArt category tree, optionally starting from a specific category path. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseCategoryTreeArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseCategoryTree(ctx, args.CategoryPath, args.Depth)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_user_friends",
		Description: "Returns a DeviantArt user's friends (users they watch). Requires a username. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseUserFriendsArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseUserFriends(ctx, args.Username, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_user_journals",
		Description: "Returns a DeviantArt user's journals. Requires a username. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseUserJournalsArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseUserJournals(ctx, args.Username, args.Featured, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "browse_user_literature",
		Description: "Returns a DeviantArt user's literature deviations. Requires a username. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args browseUserLiteratureArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.BrowseUserLiterature(ctx, args.Username, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})
}
