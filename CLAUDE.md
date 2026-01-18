# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go WebDAV is a minimal, single-user Network Attached Storage (NAS) server written in Go. It provides WebDAV protocol access to a local directory with HTTP Basic Authentication.

## Build Commands

```bash
# Local build (native architecture)
make local
# Output: bin/Local/Go-WebDAV

# Linux AMD64 build
make linux-amd64
# Output: bin/Linux-AMD64/Go-WebDAV

# Linux ARM64 build
make linux-arm64
# Output: bin/Linux-ARM64/Go-WebDAV

# Build all platforms
make all

# Clean build artifacts
make clean

# Show all available targets
make help

# Docker build and run
make docker-up

# Manual build (without Makefile)
go build -o Go-WebDAV .
```

## Running the Server

**Environment Variables Required:**
- `USERNAME` - Authentication username
- `PASSWORD` - Authentication password

```bash
# Run locally (builds and runs)
USERNAME=myuser PASSWORD=mypass make run

# Or build first, then run
make local
export USERNAME=myuser
export PASSWORD=mypass
./bin/Local/Go-WebDAV

# Or with Docker
# Edit docker-compose.yml to set USERNAME and PASSWORD
make docker-up
```

The server listens on port 8080 and serves files from the `./data` directory.

## Architecture

This is a flat, simple architecture with no subdirectories:

- [main.go](main.go) - Entry point; sets up WebDAV handler with Basic Auth middleware
- [auth.go](auth.go) - Authentication logic comparing credentials against environment variables
- [checks.go](checks.go) - Startup validation (checks for USERNAME/PASSWORD env vars and creates `./data` directory if missing)
- [return_ip.go](return_ip.go) - Helper to determine local IP for logging connection URL

**Key Architecture Details:**

1. **WebDAV Handler**: Uses `golang.org/x/net/webdav` package with an in-memory lock system (`NewMemLS()`) and silent logger for performance
2. **File System**: Serves `./data` directory at root path `/`
3. **Authentication**: HTTP Basic Auth middleware wraps the WebDAV handler; credentials validated on every request (only failed attempts are logged to reduce CPU overhead)
4. **HTTP Server**: Custom `http.Server` instance with optimizations for large file transfers:
   - No read/write timeouts (supports unlimited file sizes)
   - 120-second idle timeout (enables connection reuse)
   - 1MB max header size
5. **Startup Flow**:
   - Check USERNAME/PASSWORD env vars exist and warn if defaults
   - Create `./data` directory if missing
   - Initialize WebDAV handler with silent logger
   - Wrap with auth middleware
   - Start optimized HTTP server on `:8080`

## Performance Optimizations

The server is optimized for high-speed large file transfers with minimal CPU usage:

1. **Minimal Logging**: Only failed login attempts are logged; successful authentications and WebDAV operations use silent logging to eliminate I/O overhead during transfers
2. **Optimized HTTP Server Configuration** (`main.go:42-49`):
   - `ReadTimeout: 0` - No timeout for large uploads
   - `WriteTimeout: 0` - No timeout for large downloads
   - `IdleTimeout: 120s` - Keeps TCP connections alive for reuse across chunked requests
   - `MaxHeaderBytes: 1MB` - Reasonable security limit
3. **Silent WebDAV Logger** (`main.go:19-21`): Suppresses internal WebDAV debug output

These optimizations eliminate CPU spikes and maintain consistent transfer speeds regardless of file size.

## Dependencies

- Go 1.25.5
- `golang.org/x/net/webdav` - WebDAV protocol implementation

Run `go mod download` to fetch dependencies.

## Docker Deployment

The Dockerfile uses a multi-stage build:
1. Build stage: golang:1.25 image compiles static binary
2. Runtime stage: alpine:latest for minimal footprint

Data persistence is via volume mount of `./data` to `/data` in the container.

## Security Notes

- HTTP Basic Auth credentials are validated from environment variables
- Default credentials (`username`/`password`) trigger a warning
- Recommended to deploy behind reverse proxy (Caddy/NGINX) for HTTPS
- Server binds to all interfaces (`:8080`); use firewall rules or Docker port binding (`127.0.0.1:8080:8080`) to restrict access
