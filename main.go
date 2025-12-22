package main

import (
	"errors"
	"fmt"
	"os"

	"ascii-art/parser"
	"ascii-art/renderer"
)

const (
	// Exit codes for different error scenarios
	exitCodeUsageError  = 1
	exitCodeBannerError = 2
	exitCodeRenderError = 3

	// Default banner style
	defaultBanner = "standard"
)

func main() {
	// Parse command-line arguments
	text, banner, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeUsageError)
	}

	// Get banner file path
	bannerPath, err := GetBannerPath(banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(exitCodeUsageError)
	}

	// Load character map using parser
	charMap, err := parser.BuildCharacterMap(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner file: %v\n", err)
		os.Exit(exitCodeBannerError)
	}

	// Render the text using renderer
	result, err := renderer.RenderText(text, charMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering text: %v\n", err)
		os.Exit(exitCodeRenderError)
	}

	// Print result to stdout
	fmt.Print(result)
}

// ParseArgs parses command-line arguments and returns text and banner name
func ParseArgs(args []string) (text string, banner string, err error) {
	// args[0] is program name, so we need at least 2 elements
	if len(args) < 2 {
		return "", "", errors.New("usage: go run . \"text\" [banner]")
	}

	// Too many arguments
	if len(args) > 3 {
		return "", "", errors.New("too many arguments\nusage: go run . \"text\" [banner]")
	}

	// Get text (args[1])
	text = args[1]

	// Get banner (args[2] if provided, otherwise default to "standard")
	if len(args) == 3 {
		banner = args[2]
	} else {
		banner = defaultBanner
	}

	return text, banner, nil
}

// GetBannerPath converts banner name to file path
func GetBannerPath(banner string) (string, error) {
	// Map of valid banner names to file paths
	bannerPaths := map[string]string{
		"standard":   "testdata/standard.txt",
		"shadow":     "testdata/shadow.txt",
		"thinkertoy": "testdata/thinkertoy.txt",
	}

	// Check if banner is valid
	path, exists := bannerPaths[banner]
	if !exists {
		return "", fmt.Errorf("invalid banner name: %q\nValid options: standard, shadow, thinkertoy", banner)
	}

	return path, nil
}
