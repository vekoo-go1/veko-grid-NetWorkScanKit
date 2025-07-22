@echo off
echo.
echo ╔═══════════════════════════════════════════════════════════════╗
echo ║  🛰️  VEKO GRID - BUILD SCRIPT FOR WINDOWS                    ║
echo ╚═══════════════════════════════════════════════════════════════╝
echo.

echo [INFO] Downloading dependencies...
go mod tidy

if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Failed to download dependencies
    pause
    exit /b 1
)

echo [INFO] Building Veko Grid for Windows...
go build -ldflags="-s -w" -o veko-grid.exe .

if %ERRORLEVEL% NEQ 0 (
    echo [ERROR] Build failed
    pause
    exit /b 1
)

echo.
echo ✅ Build successful!
echo 📦 Binary created: veko-grid.exe
echo 💾 Size: 
for %%I in (veko-grid.exe) do echo    %%~zI bytes

echo.
echo 🚀 Testing binary...
veko-grid.exe --help

echo.
echo ✅ Veko Grid berhasil dibuild!
echo 📖 Baca README.md untuk petunjuk penggunaan
echo 🎯 Contoh: veko-grid.exe scan --input targets.txt --output results.json
echo.
pause