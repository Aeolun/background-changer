#!/bin/bash
# Build script for Background Quote

set -e

echo "=== Background Quote Build Script ==="
echo ""

# Check Go installation
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go 1.21 or later."
    echo "Download from: https://go.dev/dl/"
    exit 1
fi

echo "Go version: $(go version)"
echo ""

# Download dependencies
echo "Downloading dependencies..."
go mod tidy
go mod download
echo ""

# Build
echo "Building application..."
go build -ldflags="-s -w" -o background-quote .

if [ $? -eq 0 ]; then
    echo ""
    echo "=== Build successful! ==="
    echo ""
    echo "Binary created: ./background-quote"
    echo "Size: $(du -h background-quote | cut -f1)"
    echo ""
    echo "To run: ./background-quote"
    echo ""
else
    echo ""
    echo "Build failed!"
    exit 1
fi
