package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Updater handles periodic wallpaper updates
type Updater struct {
	config       *Config
	ticker       *time.Ticker
	stopChan     chan bool
	statusFunc   func(string)
	dataDir      string
}

// NewUpdater creates a new updater
func NewUpdater(config *Config) *Updater {
	// Get data directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dataDir := filepath.Join(homeDir, ".local", "share", "background-quote")

	return &Updater{
		config:   config,
		stopChan: make(chan bool),
		dataDir:  dataDir,
	}
}

// SetStatusFunc sets the function to call for status updates
func (u *Updater) SetStatusFunc(f func(string)) {
	u.statusFunc = f
}

// Start begins the update cycle
func (u *Updater) Start() {
	// Do initial update
	u.updateStatus("Performing initial update...")
	if err := u.Update(); err != nil {
		log.Printf("Initial update failed: %v", err)
		u.updateStatus(fmt.Sprintf("Error: %v", err))
	} else {
		u.updateStatus("Update complete")
	}

	// Start ticker
	duration := time.Duration(u.config.UpdateDelay) * time.Second
	u.ticker = time.NewTicker(duration)

	for {
		select {
		case <-u.ticker.C:
			u.updateStatus("Updating wallpaper...")
			if err := u.Update(); err != nil {
				log.Printf("Update failed: %v", err)
				u.updateStatus(fmt.Sprintf("Error: %v", err))
			} else {
				nextUpdate := time.Now().Add(duration)
				u.updateStatus(fmt.Sprintf("Next update: %s", nextUpdate.Format("15:04:05")))
			}
		case <-u.stopChan:
			u.ticker.Stop()
			return
		}
	}
}

// Stop stops the updater
func (u *Updater) Stop() {
	u.stopChan <- true
}

// Update performs a single wallpaper update
func (u *Updater) Update() error {
	log.Println("Starting wallpaper update...")

	// Fetch quote
	quote, err := FetchQuote(u.config.QuoteURL)
	if err != nil {
		return fmt.Errorf("failed to fetch quote: %w", err)
	}
	log.Printf("Quote: %s - %s", quote.Text, quote.Author)

	// Get background image
	var img image.Image
	if u.config.LocalImagesEnabled && u.config.LocalImageDirectory != "" {
		log.Println("Loading local image...")
		img, err = LoadLocalImage(u.config.LocalImageDirectory)
		if err != nil {
			return fmt.Errorf("failed to load local image: %w", err)
		}
	} else {
		log.Println("Downloading image...")
		imageURL := u.config.ImageURL
		if u.config.BackgroundKeywords != "" {
			imageURL += "?" + u.config.BackgroundKeywords
		}
		img, err = DownloadImage(imageURL)
		if err != nil {
			return fmt.Errorf("failed to download image: %w", err)
		}
	}

	// Resize and crop
	log.Println("Processing image...")
	img = ResizeAndCrop(img, TargetWidth, TargetHeight)

	// Overlay quote
	img, err = OverlayQuote(img, quote)
	if err != nil {
		return fmt.Errorf("failed to overlay quote: %w", err)
	}

	// Save image
	outputPath := filepath.Join(u.dataDir, "current.jpg")
	log.Printf("Saving to: %s", outputPath)
	if err := SaveImage(img, outputPath); err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}

	// Set as wallpaper
	log.Println("Setting wallpaper...")
	if err := SetWallpaper(outputPath); err != nil {
		return fmt.Errorf("failed to set wallpaper: %w", err)
	}

	log.Println("Wallpaper update complete!")
	return nil
}

// updateStatus calls the status function if set
func (u *Updater) updateStatus(msg string) {
	if u.statusFunc != nil {
		u.statusFunc(msg)
	}
}

// Restart restarts the updater with new configuration
func (u *Updater) Restart() {
	if u.ticker != nil {
		u.ticker.Stop()
	}

	duration := time.Duration(u.config.UpdateDelay) * time.Second
	u.ticker = time.NewTicker(duration)
}
