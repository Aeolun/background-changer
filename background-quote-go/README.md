# Background Quote - Go Edition

A cross-platform desktop application that automatically changes your wallpaper with inspirational quotes overlaid on beautiful background images.

This is a Go + Fyne port of the original C# Windows application, now supporting Linux, macOS, and Windows.

## Features

- 🖼️ **Automatic wallpaper rotation** with configurable intervals
- 💬 **Inspirational quotes** fetched from API and overlaid on images
- 🎨 **Beautiful backgrounds** from Unsplash or your local collection
- 🔧 **System tray integration** - runs quietly in the background
- ⚙️ **Configurable settings** - keywords, update frequency, local images
- 🚀 **Auto-start support** - launch on system startup
- 🐧 **Multi-desktop support** - works with GNOME, KDE, XFCE, Cinnamon, and more

## Supported Desktop Environments

The application automatically detects and supports:
- **GNOME** (Ubuntu, Fedora, etc.)
- **Cinnamon** (Linux Mint)
- **MATE**
- **XFCE**
- **KDE Plasma**
- **Window managers** using `feh` or `nitrogen`
- **macOS** (via AppleScript)

## Requirements

### Linux
- Go 1.21 or later
- GCC (for CGO compilation)
- Development libraries:
  ```bash
  # Ubuntu/Debian
  sudo apt-get install gcc libgl1-mesa-dev xorg-dev

  # Fedora
  sudo dnf install gcc mesa-libGL-devel libXcursor-devel libXrandr-devel libXinerama-devel libXi-devel

  # Arch
  sudo pacman -S go gcc libgl libxcursor libxrandr libxinerama libxi
  ```

### macOS
- Go 1.21 or later
- Xcode Command Line Tools

### Windows
- Go 1.21 or later
- GCC (via MinGW or TDM-GCC)

## Installation

### From Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/aeolun/background-quote.git
   cd background-quote/background-quote-go
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Build the application:**
   ```bash
   go build -o background-quote
   ```

4. **Run:**
   ```bash
   ./background-quote
   ```

### Quick Build Script

```bash
# Build for your current platform
go build -ldflags="-s -w" -o background-quote .

# Or for a smaller binary
go build -ldflags="-s -w" -tags release -o background-quote .
```

## Usage

### First Run

1. The application starts minimized in the system tray
2. Double-click the tray icon to open the settings window
3. Configure your preferences:
   - **Background Keywords**: Add keywords for image search (e.g., "nature,mountains,forest")
   - **Update Interval**: Set how often to change the wallpaper (in seconds, minimum 60)
   - **Local Images**: Optionally use your own image collection instead of downloading
   - **Run on Startup**: Enable to start automatically when you log in

4. Click **"Update Now"** to test immediately
5. Click **"Save Settings"** to persist your configuration

### System Tray Menu

Right-click the system tray icon to access:
- **Show**: Open the settings window
- **Update Now**: Immediately change the wallpaper
- **Quit**: Exit the application

## Configuration

Settings are automatically saved using Fyne's preferences system:
- **Linux**: `~/.config/fyne/com.github.aeolun.backgroundquote.preferences.json`
- **macOS**: `~/Library/Preferences/com.github.aeolun.backgroundquote.plist`
- **Windows**: Registry under `HKCU\Software\fyne.io\backgroundquote`

Generated wallpapers are stored in:
- **Linux**: `~/.local/share/background-quote/`
- **macOS**: `~/Library/Application Support/background-quote/`
- **Windows**: `%APPDATA%\background-quote\`

## API Sources

- **Quotes**: [Forismatic API](https://api.forismatic.com/)
- **Images**: [Unsplash Source](https://source.unsplash.com/)

You can customize these URLs in the code if you prefer different sources.

## Development

### Project Structure

```
background-quote-go/
├── main.go          # Application entry point
├── config.go        # Configuration management
├── quote.go         # Quote fetching from API
├── image.go         # Image downloading and processing
├── wallpaper.go     # Wallpaper setting (multi-platform)
├── updater.go       # Background update timer
├── ui.go            # Fyne GUI and system tray
├── go.mod           # Go module definition
└── README.md        # This file
```

### Key Dependencies

- **Fyne v2**: Cross-platform GUI framework
- **freetype**: Font rendering for quote text
- **resize**: Image resizing library
- **golang.org/x/image**: Additional image utilities

### Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o background-quote-linux

# macOS
GOOS=darwin GOARCH=amd64 go build -o background-quote-macos

# Windows
GOOS=windows GOARCH=amd64 go build -o background-quote.exe
```

## Troubleshooting

### Wallpaper not changing

1. Check the logs for error messages
2. Ensure you have the correct desktop environment tools installed:
   ```bash
   # For GNOME
   gsettings --version

   # For KDE
   qdbus --version

   # For XFCE
   xfconf-query --version

   # Fallback options
   which feh
   which nitrogen
   ```

### Application doesn't start

1. Ensure all dependencies are installed
2. Try running from terminal to see error messages
3. Check that X11/Wayland display is available

### Quotes not loading

1. Verify internet connection
2. Check if Forismatic API is accessible: `curl https://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=en`
3. Consider using an alternative quote API

## Comparison with C# Version

| Feature | C# (Windows) | Go (Cross-platform) |
|---------|-------------|---------------------|
| Platform | Windows only | Linux, macOS, Windows |
| UI Framework | WinForms | Fyne |
| System Tray | ✅ | ✅ |
| Local Images | ✅ | ✅ |
| Auto-start | ✅ | ✅ |
| Image Processing | ImageProcessor | Go stdlib + resize |
| Binary Size | ~2MB | ~15MB (includes GUI) |

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## License

This project follows the same license as the original C# version.

## Credits

- Original C# Windows version by the repository owner
- Go port uses the excellent [Fyne](https://fyne.io/) framework
- Quotes from [Forismatic](https://forismatic.com/)
- Images from [Unsplash](https://unsplash.com/)

## Roadmap

- [ ] Add more quote sources
- [ ] Support for multi-monitor setups
- [ ] Custom font selection
- [ ] Quote position customization
- [ ] Image filters and effects
- [ ] Notification on wallpaper change
- [ ] Export/import settings
- [ ] Packaging for various Linux distributions
