// Package parser provides functionality for loading and parsing ASCII art banner files.
// Supports standard ASCII printable characters (32-126) from banner files in the format:
// 8 lines per character + 1 separator line, 855 total lines for 95 characters.
package parser

import (
	"bufio"
	"fmt"
	"os"
)

const (
	firstPrintable rune = 32  // ASCII 32 (space)
	lastPrintable  rune = 126 // ASCII 126 (tilde)
	totalChars          = 95
	expectedLines       = 855 // 95 chars × 9 lines (8 glyph + 1 separator)
	linesPerGlyph       = 8
	linesPerChar        = 9 // 8 glyph + 1 separator
)

// Banner represents the ASCII-art data for all supported characters.
type Banner map[rune][]string

// LoadBanner reads a banner file (e.g. standard.txt) and returns its parsed representation.
func LoadBanner(path string) (Banner, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file %q: %w", path, err)
	}
	banner, err := buildBanner(lines)
	if err != nil {
		return nil, fmt.Errorf("failed to parse banner %q: %w", path, err)
	}
	return banner, nil
}

// readLines opens the file at the given path and returns all its lines as a slice of strings.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path) // #nosec G304 -- trusted banner files
	if err != nil {
		return nil, err
	}
	defer file.Close() //nolint:errcheck // read-only file

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

func buildBanner(lines []string) (Banner, error) {
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty banner file")
	}
	if len(lines) != expectedLines {
		return nil, fmt.Errorf("invalid format: expected %d lines, got %d",
			expectedLines, len(lines))
	}

	banner := make(Banner)
	runeCode := firstPrintable
	i := 0

	for i+linesPerGlyph <= len(lines) && runeCode <= lastPrintable {
		block := lines[i : i+linesPerGlyph]
		banner[runeCode] = block
		runeCode++
		i += linesPerChar
	}

	if len(banner) != totalChars {
		return nil, fmt.Errorf("incomplete banner: got %d chars, expected %d",
			len(banner), totalChars)
	}
	return banner, nil
}
