# Auth Package

OAuth2 token management for DeviantArt API.

## Responsibilities

- oauth2.Config setup for DeviantArt endpoints
- Token refresh logic
- File-based token storage (local execution)
- Two OAuth flows:
  - client-credentials for public browsing
  - authorization-code for user-specific actions

## Constraints

- No MCP types imported
- No DeviantArt API calls (that's in internal/deviantart)
- Pure OAuth flow and token persistence
- Token storage is file-based (runs locally)
