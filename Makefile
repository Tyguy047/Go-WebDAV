.PHONY: all local linux-amd64 linux-arm64 clean docker-build docker-up docker-down run help

# Build variables
BINARY_NAME=Go-WebDAV
BUILD_FLAGS=CGO_ENABLED=0 go build -trimpath -ldflags="-s -w"

# Output directories
BIN_LOCAL=bin/Local
BIN_LINUX_AMD64=bin/Linux-AMD64
BIN_LINUX_ARM64=bin/Linux-ARM64

# Default target
all: local linux-amd64 linux-arm64

# Build for local architecture
local:
	@mkdir -p $(BIN_LOCAL)
	$(BUILD_FLAGS) -o $(BIN_LOCAL)/$(BINARY_NAME)
	@echo "Built $(BIN_LOCAL)/$(BINARY_NAME)"

# Build for Linux AMD64
linux-amd64:
	@mkdir -p $(BIN_LINUX_AMD64)
	GOOS=linux GOARCH=amd64 $(BUILD_FLAGS) -o $(BIN_LINUX_AMD64)/$(BINARY_NAME)
	@echo "Built $(BIN_LINUX_AMD64)/$(BINARY_NAME)"

# Build for Linux ARM64
linux-arm64:
	@mkdir -p $(BIN_LINUX_ARM64)
	GOOS=linux GOARCH=arm64 $(BUILD_FLAGS) -o $(BIN_LINUX_ARM64)/$(BINARY_NAME)
	@echo "Built $(BIN_LINUX_ARM64)/$(BINARY_NAME)"

# Clean build artifacts
clean:
	rm -rf bin/
	@echo "Cleaned build artifacts"

# Docker build
docker-build:
	docker compose build
	@echo "Docker image built"

# Docker up (build and run)
docker-up:
	docker compose up -d
	@echo "Docker container started"

# Docker down
docker-down:
	docker compose down
	@echo "Docker container stopped"

# Run local build (requires USERNAME and PASSWORD environment variables)
run: local
	@if [ -z "$(USERNAME)" ] || [ -z "$(PASSWORD)" ]; then \
		echo "Error: USERNAME and PASSWORD environment variables must be set"; \
		echo "Example: USERNAME=myuser PASSWORD=mypass make run"; \
		exit 1; \
	fi
	$(BIN_LOCAL)/$(BINARY_NAME)

# Show help
help:
	@echo "Go WebDAV - Build Targets"
	@echo ""
	@echo "  make local       - Build for local architecture (bin/Local/Go-WebDAV)"
	@echo "  make linux-amd64 - Build for Linux AMD64 (bin/Linux-AMD64/Go-WebDAV)"
	@echo "  make linux-arm64 - Build for Linux ARM64 (bin/Linux-ARM64/Go-WebDAV)"
	@echo "  make all         - Build all platforms (default)"
	@echo "  make clean       - Remove build artifacts"
	@echo ""
	@echo "Docker Targets:"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-up    - Start Docker container"
	@echo "  make docker-down  - Stop Docker container"
	@echo ""
	@echo "Run Targets:"
	@echo "  make run         - Build and run local binary"
	@echo "                     (requires USERNAME and PASSWORD env vars)"
	@echo ""
	@echo "Example: USERNAME=myuser PASSWORD=mypass make run"
