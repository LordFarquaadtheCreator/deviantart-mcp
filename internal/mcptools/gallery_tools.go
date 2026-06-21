package mcptools

import (
	"context"

	"github.com/farquaad/deviantart-mcp/internal/deviantart"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type galleryAllArgs struct {
	Username string `json:"username" jsonschema:"description=DeviantArt username (optional)"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type galleryFoldersArgs struct {
	Username      string `json:"username" jsonschema:"description=DeviantArt username (optional)"`
	CalculateSize bool   `json:"calculate_size,omitempty"`
	ExtPreload    bool   `json:"ext_preload,omitempty"`
	Offset        int    `json:"offset,omitempty"`
	Limit         int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type galleryFolderArgs struct {
	Username string `json:"username" jsonschema:"description=DeviantArt username (optional)"`
	FolderID string `json:"folder_id,omitempty" jsonschema:"description=Gallery folder UUID (optional, if not provided returns root gallery)"`
	Mode     string `json:"mode,omitempty" jsonschema:"description=Gallery mode (optional)"`
	Offset   int    `json:"offset,omitempty"`
	Limit    int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type galleryFolderCreateArgs struct {
	FolderName string `json:"folder_name" jsonschema:"description=New folder name (required)"`
}

type galleryFolderDeleteArgs struct {
	FolderID string `json:"folder_id" jsonschema:"description=Folder UUID to delete (required)"`
}

func registerGalleryTools(s *mcp.Server, c *deviantart.Client) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "gallery_all",
		Description: "Returns the 'all' view of a user's gallery. Does not require user authentication (for public galleries).",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args galleryAllArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.GetGalleryAll(ctx, args.Username, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "gallery_folders",
		Description: "Returns a list of gallery folders for a DeviantArt user. Does not require user authentication (for public galleries).",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args galleryFoldersArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.GetGalleryFolders(ctx, args.Username, args.CalculateSize, args.ExtPreload, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "gallery_folder",
		Description: "Returns deviations in a specific gallery folder or root gallery. Does not require user authentication (for public galleries).",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args galleryFolderArgs) (*mcp.CallToolResult, any, error) {
		results, err := c.GetGallery(ctx, args.Username, args.FolderID, args.Mode, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(results)}}}, results, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "gallery_folder_create",
		Description: "Creates a new gallery folder. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args galleryFolderCreateArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.CreateGalleryFolder(ctx, args.FolderName)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "gallery_folder_delete",
		Description: "Deletes a gallery folder. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args galleryFolderDeleteArgs) (*mcp.CallToolResult, any, error) {
		if err := c.RemoveGalleryFolder(ctx, args.FolderID); err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: `{"success": true}`}}}, map[string]bool{"success": true}, nil
	})
}
