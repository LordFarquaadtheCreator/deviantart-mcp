# Repository Guidelines

## Project Structure

**Go Core**: `cmd/` and `internal/` contain the MCP server and DeviantArt API client.

**npm Wrapper**: `npm-wrapper/` is a thin distribution layer that downloads and executes Go binaries.

## Architecture Principles

### Layer Separation

- `cmd/deviantart-mcp/main.go` is the only cross-package import point
- `internal/deviantart` is a pure API client—no MCP knowledge
- `internal/mcptools` adapts DeviantArt operations to MCP tools
- `internal/auth` handles OAuth flows independently
- `npm-wrapper` is decoupled from Go code

### Package Boundaries

**internal/deviantart**: Pure DeviantArt REST API client. Should be testable without MCP context. No MCP types imported.

**internal/mcptools**: MCP adapter layer. Imports DeviantArt client and mcp-go server. Shapes responses for LLM consumption.

**internal/auth**: OAuth2 token management. File-based token storage for local execution. Two flows: client-credentials (public) and authorization-code (user-specific).

**internal/config**: Environment variable parsing only. No business logic.

## Development Commands

```bash
# Run Go server
go run cmd/deviantart-mcp/main.go

# Build
go build -o deviantart-mcp cmd/deviantart-mcp/main.go

# Test
go test ./...

# Release (requires git tag)
git tag v1.2.0
git push origin v1.2.0
```

## Distribution

Version sync via git tags. GoReleaser builds binaries for all platforms, npm wrapper downloads matching binary on install.

## Coding Style

- Go standard formatting (`gofmt`)
- Package comments for public APIs
- Error handling with explicit checks—no silent failures
- Interface-based design where appropriate for testability
