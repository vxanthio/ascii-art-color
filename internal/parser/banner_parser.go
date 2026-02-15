// Package parser provides functionality for loading and parsing ASCII art banner files.
//
// The parser reads banner files containing ASCII art representations for printable
// characters (range 32-126). Each banner file follows a strict format: 8 lines per
// character definition plus 1 separator line, totaling 855 lines for 95 characters.
//
// Responsibilities of this package:
//   - Read banner files from disk
//   - Validate banner file format
//   - Parse character definitions into usable data structures
//
// Any malformed banner file or invalid format results in an error.
package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
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

// LoadBanner reads a banner file from the provided filesystem and returns its parsed
// representation as a Banner map.
//
// The function reads the specified banner file, validates its format (855 lines total),
// and constructs a map associating each printable ASCII character (32-126) with its
// 8-line ASCII art representation.
//
// Parameters:
//   - fsys: The filesystem to read from (can be embed.FS, os.DirFS, or any fs.FS).
//   - path: The file path within the filesystem (e.g., "testdata/standard.txt").
//
// Returns:
//   - A Banner map containing all character definitions.
//   - An error if the file cannot be read or the format is invalid.
func LoadBanner(fsys fs.FS, path string) (Banner, error) {
	lines, err := readLines(fsys, path)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file %q: %w", path, err)
	}
	banner, err := buildBanner(lines)
	if err != nil {
		return nil, fmt.Errorf("failed to parse banner %q: %w", path, err)
	}
	return banner, nil
}

// readLines reads all lines from a file in the provided filesystem.
//
// The function uses fs.ReadFile to read the file content, then scans it line by line.
// This approach works with both embedded filesystems and disk-based filesystems.
// It handles both reading errors and scanner errors appropriately.
//
// Parameters:
//   - fsys: The filesystem to read from (can be embed.FS, os.DirFS, or any fs.FS).
//   - path: The file path within the filesystem.
//
// Returns:
//   - A slice containing all lines from the file.
//   - An error if the file cannot be opened or read.
func readLines(fsys fs.FS, path string) ([]string, error) {
	data, err := fs.ReadFile(fsys, path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, err
	}
	return lines, nil
}

// buildBanner constructs a Banner map from the raw lines read from a banner file.
//
// It validates the format (855 total lines, 95 characters with 8 lines each plus
// 1 separator) and creates a mapping from each printable ASCII character (32-126)
// to its 8-line ASCII art representation.
//
// Parameters:
//   - lines: The raw lines from a banner file.
//
// Returns:
//   - A Banner map containing all character definitions.
//   - An error if the format is invalid or incomplete.
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
	i := 1

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

// CharWidths returns the column width of each character in text based on the
// provided Banner glyph data. Each width corresponds to len(glyph[0]) for the
// character's ASCII art representation. Unknown characters get width 0.
//
// Parameters:
//   - text: The input string whose character widths are needed.
//   - banner: The loaded Banner map containing glyph data.
//
// Returns:
//   - A slice of integers with one width per character in text.
func CharWidths(text string, banner Banner) []int {
	widths := make([]int, len(text))
	for i, char := range text {
		glyph := banner[char]
		if glyph == nil {
			continue
		}
		widths[i] = len(glyph[0])
	}
	return widths
}
