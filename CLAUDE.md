# Claude Code Configuration

This file contains configuration and context for Claude Code to help it understand and work with this Go project.

## Project Overview
This is a Go backend service for "seiji" application, featuring:
- Authentication with Firebase
- Database operations with GORM
- REST API endpoints
- JWT token handling
- User management and social features

## Project Structure
- `main.go` - Entry point
- `controller/` - HTTP request handlers
- `db/` - Database connection, migration, and seeding
- `lib/` - External service integrations (Firebase, JWT)
- `middleware/` - HTTP middleware
- `models/` - Data models
- `routes/` - Route definitions
- `service/` - Business logic
- `utils/` - Utility functions

## Development Commands
```bash
# Run the application
go run main.go

# Install dependencies
go mod tidy

# Build
go build

# Test (if tests exist)
go test ./...
```

## Environment
- Language: Go
- Framework: Gin (likely, based on structure)
- Database: Uses GORM
- Authentication: Firebase Auth
- Deployment: Render (based on render.yaml)

## Notes
- Development and production environments are handled via environment variables
- Firebase configuration is in `config/` directory
- Current branch: develop
- Main branch: main