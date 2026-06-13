# npm Wrapper

Thin npm package that downloads and executes Go binaries. Decoupled from Go code.

## Responsibilities

- package.json: name, version (synced to git tags), bin field
- scripts/install.js: postinstall hook that detects platform/arch and downloads matching binary from GitHub Releases
- bin/deviantart-mcp.js: thin shim that execs the downloaded binary and forwards stdin/stdout/stderr and argv

## Constraints

- No shared code with Go module
- Downloads binary from GitHub Releases tagged with same version as npm package
- Pure distribution layer—no business logic
- Makes `npx deviantart-mcp` work identically to running Go binary directly
