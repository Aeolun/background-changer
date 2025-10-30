package main

import (
	"fyne.io/fyne/v2"
)

const (
	DefaultQuoteURL  = "https://api.forismatic.com/api/1.0/?method=getQuote&format=json&lang=en"
	DefaultImageURL  = "https://picsum.photos/1920/1080"
	DefaultDelay     = 600 // seconds
)

// Config holds application configuration
type Config struct {
	prefs               fyne.Preferences
	QuoteURL            string
	ImageURL            string
	BackgroundKeywords  string
	UpdateDelay         int
	LocalImageDirectory string
	LocalImagesEnabled  bool
	RunOnStartup        bool
}

// LoadConfig loads configuration from preferences
func LoadConfig(prefs fyne.Preferences) *Config {
	cfg := &Config{
		prefs: prefs,
	}

	// Load with defaults
	cfg.QuoteURL = prefs.StringWithFallback("quote_url", DefaultQuoteURL)
	cfg.ImageURL = prefs.StringWithFallback("image_url", DefaultImageURL)
	cfg.BackgroundKeywords = prefs.StringWithFallback("keywords", "")
	cfg.UpdateDelay = prefs.IntWithFallback("update_delay", DefaultDelay)
	cfg.LocalImageDirectory = prefs.StringWithFallback("local_dir", "")
	cfg.LocalImagesEnabled = prefs.BoolWithFallback("local_enabled", false)
	cfg.RunOnStartup = prefs.BoolWithFallback("run_on_startup", false)

	return cfg
}

// Save persists the configuration
func (c *Config) Save() {
	c.prefs.SetString("quote_url", c.QuoteURL)
	c.prefs.SetString("image_url", c.ImageURL)
	c.prefs.SetString("keywords", c.BackgroundKeywords)
	c.prefs.SetInt("update_delay", c.UpdateDelay)
	c.prefs.SetString("local_dir", c.LocalImageDirectory)
	c.prefs.SetBool("local_enabled", c.LocalImagesEnabled)
	c.prefs.SetBool("run_on_startup", c.RunOnStartup)
}
