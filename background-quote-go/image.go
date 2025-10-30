package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/nfnt/resize"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

var (
	TargetWidth  = 1920
	TargetHeight = 1080
)

// SetTargetDimensions updates the target dimensions for wallpapers
func SetTargetDimensions(width, height int) {
	TargetWidth = width
	TargetHeight = height
}

// DownloadImage fetches an image from URL
func DownloadImage(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("image API returned status: %d", resp.StatusCode)
	}

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

// LoadLocalImage picks a random image from a directory
func LoadLocalImage(directory string) (image.Image, error) {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	// Filter for image files
	var imageFiles []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		name := strings.ToLower(entry.Name())
		if strings.HasSuffix(name, ".jpg") || strings.HasSuffix(name, ".jpeg") ||
		   strings.HasSuffix(name, ".png") {
			imageFiles = append(imageFiles, filepath.Join(directory, entry.Name()))
		}
	}

	if len(imageFiles) == 0 {
		return nil, fmt.Errorf("no image files found in directory")
	}

	// Pick random image
	selectedFile := imageFiles[rand.Intn(len(imageFiles))]

	f, err := os.Open(selectedFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open image: %w", err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	return img, nil
}

// ResizeAndCrop resizes image to target dimensions with center crop
func ResizeAndCrop(img image.Image, width, height int) image.Image {
	bounds := img.Bounds()
	imgWidth := bounds.Dx()
	imgHeight := bounds.Dy()

	// Calculate scaling to cover target dimensions
	scaleX := float64(width) / float64(imgWidth)
	scaleY := float64(height) / float64(imgHeight)
	scale := math.Max(scaleX, scaleY)

	newWidth := uint(float64(imgWidth) * scale)
	newHeight := uint(float64(imgHeight) * scale)

	// Resize
	resized := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)

	// Crop to center
	resizedBounds := resized.Bounds()
	cropX := (resizedBounds.Dx() - width) / 2
	cropY := (resizedBounds.Dy() - height) / 2

	cropped := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(cropped, cropped.Bounds(), resized, image.Pt(cropX, cropY), draw.Src)

	return cropped
}

// OverlayQuote draws quote text on the image
func OverlayQuote(img image.Image, quote *Quote) (image.Image, error) {
	// Convert to RGBA for drawing
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	// Load font
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	// Setup text context
	fontSize := 26.0
	dpi := 72.0

	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)

	// Prepare quote text with author
	fullText := quote.Text + "\n - " + quote.Author

	// Calculate text box dimensions (similar to C# version)
	width := float64(bounds.Dx()) * 0.5
	charPerLine := 67
	lines := int(math.Ceil(float64(len(quote.Text))/float64(charPerLine))) + 1
	lineHeight := float64(bounds.Dy()) * 0.038
	height := float64(lines) * lineHeight

	posX := float64(bounds.Dx()) - width
	posY := float64(bounds.Dy()) * 0.2
	padding := 20.0

	// Draw semi-transparent background box
	bgColor := color.RGBA{255, 255, 255, 158}
	drawRect(rgba, int(posX-padding), int(posY-padding),
	         int(width+padding*2), int(height+padding*2), bgColor)

	// Draw text with shadow effect
	// Shadow (offset by 1,1)
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))
	pt := freetype.Pt(int(posX)+1, int(posY)+int(fontSize)+1)
	drawWrappedText(c, fullText, pt, int(width), int(lineHeight))

	// Main text
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))
	pt = freetype.Pt(int(posX), int(posY)+int(fontSize))
	drawWrappedText(c, fullText, pt, int(width), int(lineHeight))

	return rgba, nil
}

// drawRect draws a filled rectangle
func drawRect(img *image.RGBA, x, y, width, height int, col color.Color) {
	for dy := 0; dy < height; dy++ {
		for dx := 0; dx < width; dx++ {
			img.Set(x+dx, y+dy, col)
		}
	}
}

// drawWrappedText draws text with word wrapping
func drawWrappedText(c *freetype.Context, text string, pt fixed.Point26_6, maxWidth int, lineHeight int) error {
	words := strings.Fields(text)
	currentLine := ""
	y := pt.Y

	for _, word := range words {
		testLine := currentLine
		if testLine != "" {
			testLine += " "
		}
		testLine += word

		// Simple character-based wrapping (could be improved with actual text measurement)
		if len(testLine) > 67 {
			// Draw current line
			if currentLine != "" {
				_, err := c.DrawString(currentLine, fixed.Point26_6{X: pt.X, Y: y})
				if err != nil {
					return err
				}
				y += fixed.I(lineHeight)
			}
			currentLine = word
		} else {
			currentLine = testLine
		}
	}

	// Draw remaining line
	if currentLine != "" {
		_, err := c.DrawString(currentLine, fixed.Point26_6{X: pt.X, Y: y})
		if err != nil {
			return err
		}
	}

	return nil
}

// SaveImage saves the image to a file
func SaveImage(img image.Image, path string) error {
	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	// Save as JPEG with quality 85
	opts := &jpeg.Options{Quality: 85}
	if err := jpeg.Encode(f, img, opts); err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	return nil
}
