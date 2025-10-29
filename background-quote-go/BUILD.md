# Build Instructions

## Issue in Claude Code Environment

The code is complete and ready to build, but the Claude Code container environment has a DNS configuration issue where Go cannot resolve hostnames (it tries to use `[::1]:53` which isn't available).

**You'll need to build this on your local machine.**

## Local Build Instructions

### Prerequisites

1. **Install Go 1.21 or later:**
   ```bash
   # Check if Go is installed
   go version

   # If not, download from https://go.dev/dl/
   ```

2. **Install build dependencies (Linux only):**
   ```bash
   # Ubuntu/Debian
   sudo apt-get update
   sudo apt-get install -y gcc libgl1-mesa-dev xorg-dev

   # Fedora
   sudo dnf install -y gcc mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel

   # Arch Linux
   sudo pacman -S go gcc libgl libxcursor libxrandr libxinerama libxi
   ```

### Building

1. **Navigate to the project directory:**
   ```bash
   cd background-quote-go
   ```

2. **Download dependencies:**
   ```bash
   go mod tidy
   ```

3. **Build the application:**
   ```bash
   # Standard build
   go build -o background-quote .

   # Optimized build (smaller binary)
   go build -ldflags="-s -w" -o background-quote .
   ```

4. **Run the application:**
   ```bash
   ./background-quote
   ```

## Expected Dependencies

The `go mod tidy` command will download these dependencies:

```
fyne.io/fyne/v2 v2.7.0                           # GUI framework
github.com/golang/freetype v0.0.0-20170609...   # Font rendering
github.com/nfnt/resize v0.0.0-20180221...       # Image resizing
golang.org/x/image v0.32.0                       # Image utilities and fonts
```

## Troubleshooting

### "gcc: command not found"
Install GCC and development headers as shown in Prerequisites section.

### "cannot find package"
Run `go mod download` to fetch all dependencies.

### Wallpaper not changing
- Check that you have one of the supported desktop environments
- Install fallback tools: `sudo apt-get install feh` or `sudo apt-get install nitrogen`
- Run from terminal to see error messages

### Binary is large (~15-20MB)
This is normal for Fyne applications as they bundle GUI resources. You can use UPX to compress:
```bash
upx --best --lzma background-quote
```

## Cross-Platform Building

### For Linux (from any OS):
```bash
GOOS=linux GOARCH=amd64 go build -o background-quote-linux .
```

### For macOS (from any OS):
```bash
GOOS=darwin GOARCH=amd64 go build -o background-quote-macos .
```

### For Windows (from any OS):
```bash
GOOS=windows GOARCH=amd64 go build -o background-quote.exe .
```

## Development

To run without building:
```bash
go run .
```

To run tests (when added):
```bash
go test ./...
```

## First Run

1. The app starts minimized to system tray
2. Look for the icon in your system tray (may need to click "Show hidden icons")
3. Right-click → "Show" to open settings
4. Configure and click "Save Settings"
5. Click "Update Now" to test

Enjoy your inspirational wallpapers!
