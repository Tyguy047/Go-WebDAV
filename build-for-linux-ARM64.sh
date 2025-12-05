#!/bin/bash
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o bin/Linux-ARM64/Go-WebDAV