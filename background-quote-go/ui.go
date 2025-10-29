package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// MainWindow represents the main application window
type MainWindow struct {
	app     fyne.App
	window  fyne.Window
	config  *Config
	updater *Updater

	// UI elements
	keywordsEntry *widget.Entry
	delayEntry    *widget.Entry
	localDirEntry *widget.Entry
	localCheck    *widget.Check
	startupCheck  *widget.Check
	statusLabel   *widget.Label
}

// NewMainWindow creates the main application window
func NewMainWindow(app fyne.App, config *Config) *MainWindow {
	w := &MainWindow{
		app:    app,
		config: config,
		window: app.NewWindow("Background Quote"),
	}

	w.setupUI()
	w.setupSystemTray()

	return w
}

func (w *MainWindow) setupUI() {
	// Create form entries
	w.keywordsEntry = widget.NewEntry()
	w.keywordsEntry.SetText(w.config.BackgroundKeywords)
	w.keywordsEntry.SetPlaceHolder("e.g., nature,mountains")

	w.delayEntry = widget.NewEntry()
	w.delayEntry.SetText(strconv.Itoa(w.config.UpdateDelay))
	w.delayEntry.SetPlaceHolder("Seconds between updates")

	w.localDirEntry = widget.NewEntry()
	w.localDirEntry.SetText(w.config.LocalImageDirectory)
	w.localDirEntry.SetPlaceHolder("Path to local images")

	w.localCheck = widget.NewCheck("Use local images", func(checked bool) {
		w.config.LocalImagesEnabled = checked
	})
	w.localCheck.Checked = w.config.LocalImagesEnabled

	w.startupCheck = widget.NewCheck("Run on startup", func(checked bool) {
		w.config.RunOnStartup = checked
		w.handleStartup(checked)
	})
	w.startupCheck.Checked = w.config.RunOnStartup

	// Browse button for local directory
	browseBtn := widget.NewButton("Browse...", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err == nil && uri != nil {
				w.localDirEntry.SetText(uri.Path())
			}
		}, w.window)
	})

	localBox := container.NewBorder(nil, nil, nil, browseBtn, w.localDirEntry)

	// Status label
	w.statusLabel = widget.NewLabel("Ready")

	// Buttons
	updateNowBtn := widget.NewButton("Update Now", func() {
		go func() {
			w.statusLabel.SetText("Updating...")
			if err := w.updater.Update(); err != nil {
				dialog.ShowError(err, w.window)
				w.statusLabel.SetText(fmt.Sprintf("Error: %v", err))
			} else {
				w.statusLabel.SetText("Update complete")
			}
		}()
	})

	saveBtn := widget.NewButton("Save Settings", func() {
		w.saveSettings()
		dialog.ShowInformation("Saved", "Settings saved successfully", w.window)
	})

	buttons := container.NewHBox(updateNowBtn, saveBtn)

	// Create form
	form := container.NewVBox(
		widget.NewLabel("Background Keywords:"),
		w.keywordsEntry,
		widget.NewSeparator(),
		widget.NewLabel("Update Interval (seconds):"),
		w.delayEntry,
		widget.NewSeparator(),
		widget.NewLabel("Local Image Directory:"),
		localBox,
		w.localCheck,
		widget.NewSeparator(),
		w.startupCheck,
		widget.NewSeparator(),
		buttons,
		widget.NewSeparator(),
		widget.NewLabel("Status:"),
		w.statusLabel,
	)

	// Set window content
	w.window.SetContent(container.NewPadded(form))
	w.window.Resize(fyne.NewSize(500, 400))

	// Hide on close instead of exit
	w.window.SetCloseIntercept(func() {
		w.window.Hide()
	})
}

func (w *MainWindow) setupSystemTray() {
	// Check if desktop driver supports system tray
	if desk, ok := w.app.(desktop.App); ok {
		// Create menu
		menu := fyne.NewMenu("Background Quote",
			fyne.NewMenuItem("Show", func() {
				w.window.Show()
			}),
			fyne.NewMenuItem("Update Now", func() {
				go w.updater.Update()
			}),
			fyne.NewMenuItemSeparator(),
			fyne.NewMenuItem("Quit", func() {
				w.app.Quit()
			}),
		)

		desk.SetSystemTrayMenu(menu)
	}
}

func (w *MainWindow) saveSettings() {
	// Parse and validate delay
	delay, err := strconv.Atoi(w.delayEntry.Text)
	if err != nil || delay < 60 {
		dialog.ShowError(fmt.Errorf("update interval must be at least 60 seconds"), w.window)
		return
	}

	// Update config
	w.config.BackgroundKeywords = w.keywordsEntry.Text
	w.config.UpdateDelay = delay
	w.config.LocalImageDirectory = w.localDirEntry.Text
	w.config.LocalImagesEnabled = w.localCheck.Checked
	w.config.RunOnStartup = w.startupCheck.Checked

	// Save to preferences
	w.config.Save()

	// Restart updater with new settings
	w.updater.Restart()
}

func (w *MainWindow) handleStartup(enabled bool) {
	// TODO: Implement autostart for Linux
	// This would involve creating/removing a .desktop file in ~/.config/autostart/
	if enabled {
		w.createAutostartEntry()
	} else {
		w.removeAutostartEntry()
	}
}

func (w *MainWindow) createAutostartEntry() {
	// Implementation would create ~/.config/autostart/background-quote.desktop
	// For now, just a placeholder
}

func (w *MainWindow) removeAutostartEntry() {
	// Implementation would remove ~/.config/autostart/background-quote.desktop
	// For now, just a placeholder
}

// ShowAndRun shows the window and runs the application
func (w *MainWindow) ShowAndRun() {
	// Don't show window initially, just run in tray
	w.app.Run()
}

// SetUpdater sets the updater instance
func (w *MainWindow) SetUpdater(updater *Updater) {
	w.updater = updater
	updater.SetStatusFunc(func(msg string) {
		w.statusLabel.SetText(msg)
	})
}
