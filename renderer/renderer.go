// Package renderer provides functionality for converting input text into ASCII art
// using predefined banner character definitions.
//
// The renderer processes printable ASCII characters (range 32–126) and renders each
// character as an ASCII-art block with a fixed height (bannerHeight).
// Newline characters ('\n') are treated as line separators and preserved as separate
// ASCII-art blocks in the output.
//
// Responsibilities of this package:
//   - Validate input characters
//   - Validate banner integrity
//   - Render ASCII-art output
//
// Any invalid input or malformed banner data results in an error.
package renderer

import (
	"fmt"
	"strings"
)

const bannerHeight = 8

// RendererASCII converts an input string into ASCII art using the provided banner map.
//
// The input may contain printable ASCII characters (codes 32–126) and newline
// characters ('\n'). Newlines are treated as line separators and are not rendered
// as visible characters.
//
// Rendering rules:
//   - Empty input or input consisting only of a single newline returns an empty result.
//   - Consecutive newline characters produce empty output lines.
//   - Each non-empty input line is rendered as a block of bannerHeight ASCII-art rows.
//   - A trailing newline does not produce an extra ASCII-art block.
//
// Validation rules:
//   - Input must contain only printable ASCII characters (excluding '\n').
//   - Banner map must not be empty.
//   - Every character used in input must exist in the banner map.
//   - Each banner entry must contain exactly bannerHeight rows.
//
// Parameters:
//   - input: The text to render as ASCII art.
//   - banner: A map associating each rune with its ASCII-art representation.
//
// Returns:
//   - The rendered ASCII-art string.
//   - An error if input validation or banner validation fails.
func RendererASCII(input string, banner map[rune][]string) (string, error) {
	var result strings.Builder

	// Validate input characters before rendering
	if err := validateInput(input); err != nil {
		return "", err
	}

	// Split input into logical lines based on newline characters
	parts := strings.Split(input, "\n")

	// Remove trailing empty line if input ends with a newline
	if len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	// Handle special case: empty input or input containing only a newline
	if input == "" || input == "\n" {
		return "", nil
	}

	// Banner must not be empty
	if len(banner) == 0 {
		return "", fmt.Errorf("banner is empty")
	}

	for _, line := range parts {
		// Handle empty lines produced by consecutive newline characters
		if line == "" {
			result.WriteString("\n")
			continue
		}

		// Render each row of the ASCII-art block
		for i := 0; i < bannerHeight; i++ {
			for _, ch := range line {
				value, err := validateBannerCharacters(ch, banner)
				if err != nil {
					return "", err
				}
				result.WriteString(value[i])
			}
			result.WriteString("\n")
		}
	}

	output := result.String()

	// Remove the final trailing newline from the output
	if output != "" && output[len(output)-1] == '\n' {
		output = output[:len(output)-1]
	}

	return output, nil
}

// validateBannerCharacters validates that a character exists in the banner map
// and that its ASCII-art representation has the correct height.
//
// Parameters:
//   - ch: The character to validate.
//   - banner: The banner map containing ASCII-art definitions.
//
// Returns:
//   - The ASCII-art rows corresponding to the character.
//   - An error if the character does not exist in the banner
//     or if it does not contain exactly bannerHeight rows.

func validateBannerCharacters(ch rune, banner map[rune][]string) ([]string, error) {

	value, exists := banner[ch]
	if exists == false {
		return []string{}, fmt.Errorf("character %c (ASCII %d) not found in banner", ch, ch)
	}
	if len(value) != bannerHeight {
		return []string{}, fmt.Errorf("invalid character at position  %c (ASCII %d) - must be printable ASCII (32-126) or newline", ch, ch)
	}
	return value, nil
}

// validateInput checks whether the input string contains only valid characters.
//
// Valid characters are printable ASCII characters (codes 32–126) and newline
// characters ('\n'). The function returns an error as soon as an invalid character
// is encountered.
func validateInput(input string) error {
	for _, ch := range input {
		if ch == '\n' {
			continue
		}
		if ch < 32 || ch > 126 {

			return fmt.Errorf("not printable characters")
		}
	}
	return nil
}
