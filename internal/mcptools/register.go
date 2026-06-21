package mcptools

import (
	"github.com/farquaad/deviantart-mcp/internal/deviantart"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Register wires all tool groups into the MCP server.
func Register(s *mcp.Server, c *deviantart.Client) {
	registerBrowseTools(s, c)
	registerDeviationTools(s, c)
	registerGalleryTools(s, c)
	registerCollectionsTools(s, c)
	registerMessagesTools(s, c)
}
