# deviantart-mcp

MCP (Model Context Protocol) server for DeviantArt API. Provides tools for browsing, searching, and interacting with DeviantArt content through LLMs.

## Architecture

**Go Core** (`cmd/`, `internal/`): Self-contained Go MCP server with DeviantArt API client.

**npm Wrapper** (`npm-wrapper/`): Thin npm package that downloads and executes the compiled Go binary.

### Directory Structure

```
cmd/deviantart-mcp/main.go   # entrypoint: load config, build client, register tools, start stdio transport
internal/
  config/config.go            # env var parsing (client ID/secret, token storage path, etc.)
  auth/
    oauth.go                  # oauth2.Config setup, token refresh logic
    tokenstore.go             # persist/load tokens (file-based, since this runs locally)
  deviantart/
    client.go                 # base HTTP client: request building, error handling, rate limiting
    browse.go                 # popular/newest/tags/search endpoints
    deviation.go              # deviation metadata, stats, download links
    gallery.go                # user galleries, folders
    collections.go            # collections/favorites
    messages.go               # notifications, comments, messages
    types.go                  # shared structs for API responses
  mcptools/
    register.go               # wires all tool groups into the mcp-go server
    browse_tools.go           # MCP tool defs that call deviantart/browse.go
    deviation_tools.go
    gallery_tools.go
    collections_tools.go
    messages_tools.go
npm-wrapper/
  package.json                # name, version (synced to git tags), bin field
  scripts/install.js          # postinstall: detect platform/arch, download matching binary from GitHub Releases
  bin/deviantart-mcp.js       # thin shim: exec the downloaded binary, pipe stdio/args through
```

## Key Relationships

- `cmd/deviantart-mcp/main.go` is the only cross-package import—wires config, auth, client, and MCP tools
- `internal/deviantart` is pure DeviantArt API client—zero MCP knowledge, testable in isolation
- `internal/mcptools` is the adapter layer—shapes DeviantArt operations into MCP tools
- `internal/auth` handles OAuth flows (client-credentials for public browsing, authorization-code for user actions)
- `npm-wrapper` is decoupled from Go—downloads binary from GitHub Releases matching npm version

## Distribution

Version sync via git tags: `v1.2.0` triggers GoReleaser to build binaries, then npm publishes `package.json@1.2.0`.

## Development

```bash
# Run Go server locally
go run cmd/deviantart-mcp/main.go

# Build for current platform
go build -o deviantart-mcp cmd/deviantart-mcp/main.go

# Test npm wrapper locally
cd npm-wrapper && npm link && npx deviantart-mcp
```
