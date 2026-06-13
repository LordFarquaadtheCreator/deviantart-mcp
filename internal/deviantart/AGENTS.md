# DeviantArt Package

Pure DeviantArt REST API client. Zero MCP knowledge.

## Responsibilities

- HTTP client with request building
- Error handling and rate limiting
- DeviantArt API endpoints:
  - browse.go: popular/newest/tags/search
  - deviation.go: metadata, stats, download links
  - gallery.go: user galleries, folders
  - collections.go: collections/favorites
  - messages.go: notifications, comments, messages
- Shared response structs in types.go

## Constraints

- No MCP types imported
- No OAuth logic (that's in internal/auth)
- Testable without MCP context
- Could be extracted as standalone Go DeviantArt client library
- Pure API client—no LLM-specific shaping
