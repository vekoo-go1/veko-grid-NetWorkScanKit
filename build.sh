#!/bin/bash

echo ""
echo "╔═══════════════════════════════════════════════════════════════╗"
echo "║  🛰️  VEKO GRID - BUILD SCRIPT FOR LINUX/MAC                  ║"
echo "╚═══════════════════════════════════════════════════════════════╝"
echo ""

echo "[INFO] Downloading dependencies..."
go mod tidy

if [ $? -ne 0 ]; then
    echo "[ERROR] Failed to download dependencies"
    exit 1
fi

echo "[INFO] Building Veko Grid..."

# Detect OS
OS=$(uname -s)
case "$OS" in
    Linux*)
        echo "[INFO] Building for Linux..."
        go build -ldflags="-s -w" -o veko-grid-linux .
        BINARY="veko-grid-linux"
        ;;
    Darwin*)
        echo "[INFO] Building for macOS..."
        go build -ldflags="-s -w" -o veko-grid-mac .
        BINARY="veko-grid-mac"
        ;;
    *)
        echo "[INFO] Building for generic Unix..."
        go build -ldflags="-s -w" -o veko-grid .
        BINARY="veko-grid"
        ;;
esac

if [ $? -ne 0 ]; then
    echo "[ERROR] Build failed"
    exit 1
fi

echo ""
echo "✅ Build successful!"
echo "📦 Binary created: $BINARY"
echo "💾 Size: $(du -h $BINARY | cut -f1)"

echo ""
echo "🚀 Testing binary..."
./$BINARY --help

echo ""
echo "✅ Veko Grid berhasil dibuild!"
echo "📖 Baca README.md untuk petunjuk penggunaan"
echo "🎯 Contoh: ./$BINARY scan --input targets.txt --output results.json"
echo ""