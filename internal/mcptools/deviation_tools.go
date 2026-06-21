package mcptools

import (
	"context"

	"github.com/farquaad/deviantart-mcp/internal/deviantart"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type deviationGetArgs struct {
	DeviationID string `json:"deviation_id" jsonschema:"description=Deviation UUID (required, format: XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX)"`
}

type deviationContentArgs struct {
	DeviationID   string `json:"deviation_id" jsonschema:"description=Deviation UUID (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
}

type deviationDownloadArgs struct {
	DeviationID   string `json:"deviation_id" jsonschema:"description=Deviation UUID (required)"`
	MatureContent bool   `json:"mature_content,omitempty"`
}

type deviationWhoFavedArgs struct {
	DeviationID string `json:"deviation_id" jsonschema:"description=Deviation UUID (required)"`
	Offset      int    `json:"offset,omitempty"`
	Limit       int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type deviationEmbeddedContentArgs struct {
	DeviationID       string `json:"deviation_id" jsonschema:"description=Deviation UUID (required)"`
	OffsetDeviationID string `json:"offset_deviation_id,omitempty" jsonschema:"description=Offset deviation UUID for pagination"`
	Offset            int    `json:"offset,omitempty"`
	Limit             int    `json:"limit,omitempty" jsonschema:"description=Number of results 1-50,default=10"`
}

type deviationMetadataArgs struct {
	DeviationIDs   []string `json:"deviation_ids" jsonschema:"description=List of deviation UUIDs (required, max 50)"`
	ExtSubmission  bool     `json:"ext_submission,omitempty"`
	ExtCamera      bool     `json:"ext_camera,omitempty"`
	ExtStats       bool     `json:"ext_stats,omitempty"`
	ExtCollection  bool     `json:"ext_collection,omitempty"`
}

func registerDeviationTools(s *mcp.Server, c *deviantart.Client) {
	mcp.AddTool(s, &mcp.Tool{
		Name:        "deviation_get",
		Description: "Returns metadata for a single deviation by its UUID. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args deviationGetArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.GetDeviation(ctx, args.DeviationID)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "deviation_content",
		Description: "Returns the full content of a deviation (journal body, literature text). Requires user authentication for some deviations.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args deviationContentArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.GetDeviationContent(ctx, args.DeviationID, args.MatureContent)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "deviation_download",
		Description: "Returns download information (URL, dimensions, filesize) for a deviation. Requires user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args deviationDownloadArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.GetDeviationDownload(ctx, args.DeviationID, args.MatureContent)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "deviation_whofaved",
		Description: "Returns users who favorited a deviation. Requires a deviation UUID. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args deviationWhoFavedArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.DeviationWhoFaved(ctx, args.DeviationID, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "deviation_embedded_content",
		Description: "Returns content embedded in a deviation (journal and literature deviations support embedding). Requires deviation UUID.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args deviationEmbeddedContentArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.GetDeviationEmbeddedContent(ctx, args.DeviationID, args.OffsetDeviationID, args.Offset, args.Limit)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})

	mcp.AddTool(s, &mcp.Tool{
		Name:        "deviation_metadata",
		Description: "Returns deviation metadata for a set of deviations. Limited to 50 deviations per query. Does not require user authentication.",
	}, func(ctx context.Context, req *mcp.CallToolRequest, args deviationMetadataArgs) (*mcp.CallToolResult, any, error) {
		result, err := c.GetDeviationMetadata(ctx, args.DeviationIDs, args.ExtSubmission, args.ExtCamera, args.ExtStats, args.ExtCollection)
		if err != nil {
			return nil, nil, err
		}
		return &mcp.CallToolResult{Content: []mcp.Content{&mcp.TextContent{Text: mustJSON(result)}}}, result, nil
	})
}
