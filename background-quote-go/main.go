package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	// Create application
	a := app.NewWithID("com.github.aeolun.backgroundquote")

	// Load configuration
	cfg := LoadConfig(a.Preferences())

	// Create main window
	win := NewMainWindow(a, cfg)

	// Create updater
	updater := NewUpdater(cfg, a)
	win.SetUpdater(updater)

	// Start background updater
	go updater.Start()

	// Run application (window hidden by default, runs in system tray)
	win.ShowAndRun()
}
