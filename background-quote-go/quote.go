package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Quote represents a quote with text and author
type Quote struct {
	Text   string `json:"quoteText"`
	Author string `json:"quoteAuthor"`
}

// FetchQuote retrieves a quote from the API
func FetchQuote(url string) (*Quote, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch quote: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("quote API returned status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	var quote Quote
	if err := json.Unmarshal(body, &quote); err != nil {
		return nil, fmt.Errorf("failed to parse quote JSON: %w", err)
	}

	// Trim whitespace
	quote.Text = strings.TrimSpace(quote.Text)
	quote.Author = strings.TrimSpace(quote.Author)

	// Handle empty author
	if quote.Author == "" {
		quote.Author = "Unknown"
	}

	return &quote, nil
}
