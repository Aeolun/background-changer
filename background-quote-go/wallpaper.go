package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// SetWallpaper sets the desktop wallpaper
func SetWallpaper(imagePath string) error {
	switch runtime.GOOS {
	case "linux":
		return setWallpaperLinux(imagePath)
	case "windows":
		return fmt.Errorf("Windows support not implemented yet")
	case "darwin":
		return setWallpaperMacOS(imagePath)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

// setWallpaperLinux attempts to set wallpaper on various Linux desktop environments
func setWallpaperLinux(imagePath string) error {
	// Get absolute path
	absPath, err := filepath.Abs(imagePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	// Convert to file:// URI
	fileURI := "file://" + absPath

	// Try different desktop environments
	// Try GNOME/Ubuntu
	if err := tryCommand("gsettings", "set", "org.gnome.desktop.background", "picture-uri", fileURI); err == nil {
		// Also set dark mode variant
		tryCommand("gsettings", "set", "org.gnome.desktop.background", "picture-uri-dark", fileURI)
		return nil
	}

	// Try Cinnamon
	if err := tryCommand("gsettings", "set", "org.cinnamon.desktop.background", "picture-uri", fileURI); err == nil {
		return nil
	}

	// Try MATE
	if err := tryCommand("gsettings", "set", "org.mate.background", "picture-filename", absPath); err == nil {
		return nil
	}

	// Try XFCE
	// XFCE requires setting for each monitor, try common ones
	if err := tryCommand("xfconf-query", "-c", "xfce4-desktop",
		"-p", "/backdrop/screen0/monitor0/workspace0/last-image",
		"-s", absPath); err == nil {
		return nil
	}

	// Try KDE Plasma 5
	if err := setWallpaperKDE(absPath); err == nil {
		return nil
	}

	// Try feh (works with many window managers)
	if err := tryCommand("feh", "--bg-fill", absPath); err == nil {
		return nil
	}

	// Try nitrogen
	if err := tryCommand("nitrogen", "--set-zoom-fill", absPath); err == nil {
		return nil
	}

	return fmt.Errorf("unable to set wallpaper - no supported desktop environment detected")
}

// setWallpaperKDE sets wallpaper on KDE Plasma using a JavaScript script
func setWallpaperKDE(imagePath string) error {
	script := fmt.Sprintf(`
var allDesktops = desktops();
for (i=0; i<allDesktops.length; i++) {
	d = allDesktops[i];
	d.wallpaperPlugin = "org.kde.image";
	d.currentConfigGroup = Array("Wallpaper", "org.kde.image", "General");
	d.writeConfig("Image", "file://%s");
}
`, imagePath)

	cmd := exec.Command("qdbus", "org.kde.plasmashell", "/PlasmaShell",
		"org.kde.PlasmaShell.evaluateScript", script)
	return cmd.Run()
}

// setWallpaperMacOS sets wallpaper on macOS
func setWallpaperMacOS(imagePath string) error {
	script := fmt.Sprintf(`tell application "System Events" to tell every desktop to set picture to "%s"`, imagePath)
	cmd := exec.Command("osascript", "-e", script)
	return cmd.Run()
}

// tryCommand tries to execute a command and returns error if it fails
func tryCommand(name string, args ...string) error {
	// Check if command exists
	if _, err := exec.LookPath(name); err != nil {
		return err
	}

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
