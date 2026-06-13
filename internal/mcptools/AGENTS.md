# MCP Tools Package

Adapter layer that shapes DeviantArt operations into MCP tools.

## Responsibilities

- Wire all tool groups into mcp-go server (register.go)
- Define MCP tool schemas (input/output) for each operation
- Call internal/deviantart client methods
- Shape responses into MCP content blocks for LLM consumption
- Tool groups:
  - browse_tools.go: search, popular, newest
  - deviation_tools.go: metadata, stats, downloads
  - gallery_tools.go: user galleries, folders
  - collections_tools.go: favorites, collections
  - messages_tools.go: notifications, comments

## Constraints

- Imports internal/deviantart and mcp-go server
- No OAuth logic (that's in internal/auth)
- No direct HTTP calls (use deviantart client)
- This is where "is this tool well-designed for an LLM to call" thinking lives
- Focus on clear input schemas and helpful output formatting
