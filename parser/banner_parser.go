package parser

import (
	"bufio"
	"os"
)

// Banner represents the ASCII-art data for all supported characters.
type Banner map[rune][]string

// LoadBanner reads a banner file (e.g. standard.txt) and returns its parsed representation.
func LoadBanner(path string) (Banner, error) {
	lines, err := readLines(path)
	if err != nil {
		return nil, err // caller decides how to handle I/O errors
	}
	return buildBanner(lines), nil
}

// readLines opens the file at the given path and returns all its lines as a slice of strings.
func readLines(path string) ([]string, error) {
	// Open the file for reading.
	file, err := os.Open(path)
	if err != nil {
		return nil, err // propagate any error (e.g. file not found)
	}
	defer file.Close() // ensure the file is closed when we return

	// Create a scanner to read the file line by line.
	scanner := bufio.NewScanner(file)

	var lines []string
	// Scan each line and append its text to the slice.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// Check for any scanning error (I/O errors, etc.).
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	// Return all collected lines.
	return lines, nil
}

// firstPrintable and lastPrintable define the ASCII range this banner supports.
const (
	firstPrintable = rune(32)  // space
	lastPrintable  = rune(126) // ~
)

// buildBanner groups the given lines into 8-line glyphs for each printable ASCII rune.
// It assumes the input format is: 8 lines of glyph + 1 blank separator line per character.
func buildBanner(lines []string) Banner {
	banner := make(Banner)
	runeCode := firstPrintable // current rune we are filling (starts at space)
	i := 0
	// Each character uses 8 lines plus 1 separator line.
	for i+8 <= len(lines) && runeCode <= lastPrintable {
		block := lines[i : i+8]  // 8 lines for this rune
		banner[runeCode] = block // store glyph in the map
		runeCode++               // move to next rune
		i += 9                   // skip 8 glyph lines + 1 separator line
	}
	return banner
}
