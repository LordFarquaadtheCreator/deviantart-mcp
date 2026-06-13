# Config Package

Environment variable parsing only. No business logic.

## Responsibilities

- Load DeviantArt client ID/secret from env
- Load token storage path from env
- Validate required environment variables
- Provide configuration struct to main.go

## Constraints

- No network calls
- No file I/O beyond reading env
- No OAuth logic (that's in internal/auth)
- Pure data loading and validation
