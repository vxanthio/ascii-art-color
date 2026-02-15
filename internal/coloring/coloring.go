// Package coloring provides utilities for applying ANSI color codes to ASCII art.
//
// The package is responsible for mapping character indexes in the original
// plain text to column offsets in the rendered ASCII art, allowing substrings
// in the output to be colorized accurately.
package coloring

import (
	"strings"
)

// Reset is the ANSI escape sequence used to reset terminal coloring back
// to the default style after a colored segment.
const Reset = "\033[0m"

// ApplyColor applies ANSI color codes to matching substrings in rendered ASCII art.
//
// It determines which characters in the input text should be colored, maps those
// character positions to column ranges in the ASCII art using charWidths, and
// inserts colorCode and Reset escape sequences at the appropriate boundaries.
//
// Parameters:
//   - asciiArt: rendered ASCII art lines to be colorized
//   - text: original plain text used to generate the ASCII art
//   - substring: substring to colorize; if empty, the entire text is colored
//   - colorCode: ANSI escape sequence that starts the coloring
//   - charWidths: column widths corresponding to each character in text
//
// Returns:
//   - A new slice of strings containing the colored ASCII art
func ApplyColor(
	asciiArt []string,
	text string,
	substring string,
	colorCode string,
	charWidths []int,
) []string {
	if len(asciiArt) == 0 || len(charWidths) == 0 || len(text) == 0 {
		return asciiArt
	}

	positions := findPositions(text, substring)
	result := make([]string, len(asciiArt))

	for i, line := range asciiArt {
		result[i] = colorLine(line, positions, charWidths, colorCode)
	}

	return result
}

// colorLine applies ANSI color codes to a single line of ASCII art.
//
// It uses the boolean positions slice to determine where coloring should
// start and end, based on character boundaries defined by charWidths.
// This function assumes that positions corresponds to indexes in the
// original text, not byte offsets in the ASCII art.
//
// Parameters:
//   - line: The ASCII art line to colorize.
//   - positions: Boolean slice marking which characters should be colored.
//   - charWidths: Column widths for each character in the original text.
//   - colorCode: ANSI escape sequence for the desired color.
//
// Returns:
//   - The colorized line with ANSI color codes inserted.
func colorLine(
	line string,
	positions []bool,
	charWidths []int,
	colorCode string,
) string {
	var builder strings.Builder
	offset := 0

	for idx, width := range charWidths {
		if offset >= len(line) {
			break
		}

		end := offset + width
		if end > len(line) {
			end = len(line)
		}

		isStart := positions[idx] && (idx == 0 || !positions[idx-1])
		isEnd := positions[idx] && (idx == len(positions)-1 || !positions[idx+1])

		if isStart {
			builder.WriteString(colorCode)
		}

		builder.WriteString(line[offset:end])

		if isEnd {
			builder.WriteString(Reset)
		}

		offset = end
	}

	if offset < len(line) {
		builder.WriteString(line[offset:])
	}

	return builder.String()
}

// findPositions returns a boolean slice indicating which character indexes
// in text are part of a substring match.
//
// Each index set to true represents a character that should be colorized.
// If substring is empty, all positions in text are marked true, indicating
// that the entire text should be colored.
//
// Parameters:
//   - text: The text to search for substring matches.
//   - substring: The substring to find; if empty, all positions are marked true.
//
// Returns:
//   - A boolean slice with the same length as text, with true for matched positions.
func findPositions(text string, substring string) []bool {
	positions := make([]bool, len(text))

	if len(substring) == 0 {
		for i := range positions {
			positions[i] = true
		}
		return positions
	}

	for i := 0; i <= len(text)-len(substring); i++ {
		match := true

		for p := 0; p < len(substring); p++ {
			if text[i+p] != substring[p] {
				match = false
				break
			}
		}

		if match {
			for p := 0; p < len(substring); p++ {
				positions[i+p] = true
			}
		}
	}

	return positions
}
