// Package main provides the ASCII art generator CLI application.
//
// The application orchestrates the parser and renderer packages to convert text input
// into graphical ASCII art representations. It handles command-line argument parsing,
// banner file selection, and error reporting with appropriate exit codes.
//
// Responsibilities of this package:
//   - Parse command-line arguments
//   - Validate and resolve banner file paths
//   - Coordinate between parser and renderer
//   - Handle errors with appropriate exit codes
//
// Any invalid input, missing files, or rendering errors are reported to stderr.
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"ascii-art-color/internal/color"
	"ascii-art-color/internal/coloring"
	"ascii-art-color/internal/flagparser"
	"ascii-art-color/internal/parser"
	"ascii-art-color/internal/renderer"
)

const (
	// Exit codes for different error scenarios.
	exitCodeUsageError  = 1
	exitCodeBannerError = 2
	exitCodeRenderError = 3
	exitCodeColorError  = 4

	// Default banner style.
	defaultBanner = "standard"
)

func main() {
	if hasColorFlag(os.Args) {
		runColorMode(os.Args)
		return
	}

	text, banner, err := ParseArgs(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeUsageError)
	}

	bannerPath, err := GetBannerPath(banner)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(exitCodeUsageError)
	}

	charMap, err := parser.LoadBanner(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner file: %v\n", err)
		os.Exit(exitCodeBannerError)
	}

	result, err := renderer.RendererASCII(text, charMap)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error rendering text: %v\n", err)
		os.Exit(exitCodeRenderError)
	}

	fmt.Print(result)
}

// runColorMode handles the full color mode pipeline: validation, parsing, rendering,
// and colorizing. It prints the colored ASCII art to stdout and exits with the
// appropriate exit code on error.
//
// For multiline text (containing \n), each line is rendered and colorized separately
// so that character widths and substring positions are computed per line.
//
// Parameters:
//   - args: Command-line arguments including program name.
func runColorMode(args []string) {
	if err := flagparser.ParseArgs(args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeUsageError)
	}

	colorSpec, substring, text, bannerName, err := extractColorArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitCodeUsageError)
	}

	rgb, err := color.Parse(colorSpec)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(exitCodeColorError)
	}

	bannerPath, err := GetBannerPath(bannerName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(exitCodeUsageError)
	}

	charMap, err := parser.LoadBanner(bannerPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner file: %v\n", err)
		os.Exit(exitCodeBannerError)
	}

	colorCode := color.ANSI(rgb)
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		if line == "" {
			if i < len(lines)-1 {
				fmt.Println()
			}
			continue
		}

		art, err := renderer.RendererASCII(line, charMap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error rendering text: %v\n", err)
			os.Exit(exitCodeRenderError)
		}

		artLines := strings.Split(strings.TrimSuffix(art, "\n"), "\n")
		widths := parser.CharWidths(line, charMap)
		colored := coloring.ApplyColor(artLines, line, substring, colorCode, widths)

		for _, cl := range colored {
			fmt.Println(cl)
		}
	}
}

// hasColorFlag checks whether any argument contains the --color= flag.
func hasColorFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			return true
		}
	}
	return false
}

// isValidBanner checks whether a string is a recognized banner name.
func isValidBanner(name string) bool {
	return name == "standard" || name == "shadow" || name == "thinkertoy"
}

// extractColorArgs extracts color spec, substring, text, and banner from color-mode arguments.
//
// The function expects args[1] to be the --color=<value> flag. The remaining arguments
// are interpreted as follows:
//   - 3 args: prog --color=X text (no substring, default banner)
//   - 4 args: prog --color=X text banner (if last arg is valid banner name)
//   - 4 args: prog --color=X substring text (otherwise, default banner)
//   - 5 args: prog --color=X substring text banner
//
// Parameters:
//   - args: Command-line arguments including program name.
//
// Returns:
//   - colorSpec: The color value from the --color= flag.
//   - substring: The substring to color (empty if not provided).
//   - text: The text to render (with escape sequences interpreted).
//   - banner: The banner name to use.
//   - err: An error if extraction fails.
func extractColorArgs(args []string) (colorSpec, substring, text, banner string, err error) {
	_, colorSpec, _ = strings.Cut(args[1], "=")

	remaining := args[2:]

	switch len(remaining) {
	case 0:
		return "", "", "", "", errors.New("missing text argument")
	case 1:
		text = remaining[0]
		banner = defaultBanner
	case 2:
		if isValidBanner(remaining[1]) {
			text = remaining[0]
			banner = remaining[1]
		} else {
			substring = remaining[0]
			text = remaining[1]
			banner = defaultBanner
		}
	case 3:
		substring = remaining[0]
		text = remaining[1]
		banner = remaining[2]
	default:
		return "", "", "", "", errors.New("too many arguments")
	}

	text = strings.ReplaceAll(text, "\\n", "\n")

	return colorSpec, substring, text, banner, nil
}

// ParseArgs parses command-line arguments and extracts text and banner name.
//
// The function validates argument count, extracts the text argument, interprets
// escape sequences (like \n), and determines the banner name (defaulting to "standard"
// if not provided).
//
// Parameters:
//   - args: Command-line arguments slice (args[0] is program name).
//
// Returns:
//   - text: The text to render (with escape sequences interpreted).
//   - banner: The banner name to use.
//   - err: An error if argument validation fails.
func ParseArgs(args []string) (text string, banner string, err error) {
	if len(args) < 2 {
		return "", "", errors.New("usage: go run . \"text\" [banner]")
	}

	if len(args) > 3 {
		return "", "", errors.New("too many arguments\nusage: go run . \"text\" [banner]")
	}

	text = strings.ReplaceAll(args[1], "\\n", "\n")

	if len(args) == 3 {
		banner = args[2]
	} else {
		banner = defaultBanner
	}

	return text, banner, nil
}

// GetBannerPath converts a banner name to its corresponding file path.
//
// The function validates the banner name against a predefined map of valid banners
// (standard, shadow, thinkertoy) and returns the appropriate file path in the testdata
// directory.
//
// Parameters:
//   - banner: The banner name to resolve.
//
// Returns:
//   - The file path to the banner file.
//   - An error if the banner name is invalid.
func GetBannerPath(banner string) (string, error) {
	bannerPaths := map[string]string{
		"standard":   "testdata/standard.txt",
		"shadow":     "testdata/shadow.txt",
		"thinkertoy": "testdata/thinkertoy.txt",
	}

	path, exists := bannerPaths[banner]
	if !exists {
		return "", fmt.Errorf("invalid banner name: %q\nValid options: standard, shadow, thinkertoy", banner)
	}

	return path, nil
}
