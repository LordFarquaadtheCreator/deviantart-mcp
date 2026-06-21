package mcptools

import (
	"context"

	"github.com/farquaad/deviantart-mcp/internal/deviantart"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type collectionsFoldersArgs struct {
	Username      string `json:"username" jsonschema:"description=DeviantArt username (optional)"`
	CalculateSize bool   `json:"calculate_size,omitempty"`
	ExtPreload    bool   `json:"ext_preload,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type collectionsFolderArgs struct {
	FolderID      string `json:"folder_id" jsonschema:"description=Collection folder UUID (required)"`
	Username      string `json:"username" jsonschema:"description=DeviantArt username (optional)"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type collectionsCreateArgs struct {
	Name string `json:"name" jsonschema:"description=New collection folder name (required)"`
}

type collectionsDeleteArgs struct {
	FolderID string `json:"folder_id" jsonschema:"description=Collection folder UUID to delete (required)"`
}

type collectionsFaveArgs struct {
	DeviationID string `json:"deviation_id" jsonschema:"description=Deviation UUID to favorite (required)"`
	FolderID    string `json:"folder_id,omitempty" jsonschema:"description=Collection folder UUID (optional, defaults to favorites)"`
}

type collectionsUnfaveArgs struct {
	DeviationID string `json:"deviation_id" jsonschema:"description=Deviation UUID to unfavorite (required)"`
	FolderID    string `json:"folder_id,omitempty" jsonschema:"description=Collection folder UUID (optional)"`
}

func registerCollectionsTools(s *mcp.Server, c *deviantart.Client) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "collections_folders",
		Description: "Returns a list of collection folders for a DeviantArt user. Does not require user authentication (for public collections).",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args collectionsFoldersArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.GetCollectionsFolders(ctx, args.Username, args.CalculateSize, args.ExtPreload, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "collections_folder",
		Description: "Returns deviations in a specific collection folder. Requires collection folder UUID.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args collectionsFolderArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.GetCollections(ctx, args.FolderID, args.Username, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "collections_create",
		Description: "Creates a new collection folder. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args collectionsCreateArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.CreateCollectionFolder(ctx, args.Name)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "collections_delete",
		Description: "Deletes a collection folder. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args collectionsDeleteArgs) (*mcp.CallToolResult, any, error) {
		if err := c.RemoveCollectionFolder(ctx, args.FolderID); err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: `{"success": true}`}}}, map[string]bool{"success": true}, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "collections_fave",
		Description: "Adds a deviation to favorites. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args collectionsFaveArgs) (*mcp.CallToolResult, any, error) {
		if err := c.Fave(ctx, args.DeviationID, args.FolderID); err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: `{"success": true}`}}}, map[string]bool{"success": true}, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "collections_unfave",
		Description: "Removes a deviation from favorites. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args collectionsUnfaveArgs) (*mcp.CallToolResult, any, error) {
		if err := c.Unfave(ctx, args.DeviationID, args.FolderID); err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: `{"success": true}`}}}, map[string]bool{"success": true}, nil
	})
}
