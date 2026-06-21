package mcptools

import (
	"context"

	"github.com/farquaad/deviantart-mcp/internal/deviantart"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type messagesFeedbackArgs struct {
	UserID        string `json:"user_id" jsonschema:"description=Your DeviantArt user UUID (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type messagesFeedbackStackArgs struct {
	UserID        string `json:"user_id" jsonschema:"description=Your DeviantArt user UUID (required)"`
	StackID       string `json:"stack_id,omitempty" jsonschema:"description=Stack ID for filtering feedback"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type messagesMentionsArgs struct {
	UserID        string `json:"user_id" jsonschema:"description=Your DeviantArt user UUID (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type messagesMentionsStackArgs struct {
	UserID        string `json:"user_id" jsonschema:"description=Your DeviantArt user UUID (required)"`
	StackID       string `json:"stack_id,omitempty" jsonschema:"description=Stack ID for filtering mentions"`
	MatureContent bool   `json:"mature_content,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

func registerMessagesTools(s *mcp.Server, c *deviantart.Client) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "messages_feedback",
		Description: "Returns feedback messages (comments, replies, favourites) for a user. Requires the user's own UUID. User authentication required.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args messagesFeedbackArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.MessagesFeedback(ctx, args.UserID, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "messages_feedback_stack",
		Description: "Returns feedback messages from a specific stack for a user. Requires the user's own UUID. User authentication required.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args messagesFeedbackStackArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.MessagesFeedbackStack(ctx, args.UserID, args.StackID, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "messages_mentions",
		Description: "Returns mention notifications for a user. Requires the user's own UUID. User authentication required.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args messagesMentionsArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.MessagesMentions(ctx, args.UserID, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "messages_mentions_stack",
		Description: "Returns mention notifications from a specific stack for a user. Requires the user's own UUID. User authentication required.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args messagesMentionsStackArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.MessagesMentionsStack(ctx, args.UserID, args.StackID, args.MatureContent, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})
}
