#!/bin/bash

echo "🛰️  Building Veko Grid for Windows (x64)..."
echo ""

# Clean previous builds
echo "🧹 Cleaning previous builds..."
rm -f veko-grid-windows.exe veko-grid-win64.exe

# Download dependencies
echo "📦 Downloading dependencies..."
go mod tidy

if [ $? -ne 0 ]; then
    echo "❌ Failed to download dependencies"
    exit 1
fi

# Build for Windows with optimizations
echo "🔨 Building optimized Windows executable..."
GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -X main.version=1.0.0" -o veko-grid-win64.exe .

if [ $? -ne 0 ]; then
    echo "❌ Windows build failed"
    exit 1
fi

# Check file size
SIZE=$(stat -f%z veko-grid-win64.exe 2>/dev/null || stat -c%s veko-grid-win64.exe 2>/dev/null)
SIZE_MB=$((SIZE / 1024 / 1024))

echo ""
echo "✅ Windows build successful!"
echo "📁 File: veko-grid-win64.exe"
echo "💾 Size: ${SIZE_MB}MB"
echo ""
echo "🚀 Ready for deployment on Windows!"
echo "📋 Usage: veko-grid-win64.exe scan --input targets.txt --output results.json"
echo ""