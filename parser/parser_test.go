// Package parser provides tests for the banner parser functionality.
// Tests verify correct loading and parsing of ASCII art banner files
// in standard, shadow, and thinkertoy formats.
package parser

import (
	"testing"
)

// TestLoadBannerSpaceChar verifies that the space character loads correctly
// from the shadow banner file.
func TestLoadBannerSpaceChar(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	spaceChar := ' '
	expected := []string{
		"",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}

	actual, ok := banner[spaceChar]
	if !ok {
		t.Errorf("banner does not contain space character")
	}

	if len(actual) != len(expected) {
		t.Errorf("expected %d lines for space, got %d", len(expected), len(actual))
	}

	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// TestLoadBannerMissingFile verifies that LoadBanner returns an error
// when the specified banner file does not exist.
func TestLoadBannerMissingFile(t *testing.T) {
	_, err := LoadBanner("../testdata/nope.txt")
	if err == nil {
		t.Errorf("expected error for missing file, got nil")
	}
}

// TestLoadBannerExclamationChar verifies that the exclamation mark
// loads correctly from the shadow banner file.
func TestLoadBannerExclamationChar(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner Failed: %v", err)
	}
	char := '!'
	expected := []string{
		"",
		"   ",
		"_| ",
		"_| ",
		"_| ",
		"   ",
		"_| ",
		"   ",
	}
	actual, ok := banner[char]
	if !ok {
		t.Errorf("banner does not contain '!' character")
	}
	if len(actual) != len(expected) {
		t.Fatalf("expected %d lines for '!', got %d", len(expected), len(actual))
	}
	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// TestLoadBannerStandardSpace verifies that the space character loads
// correctly from the standard banner file.
func TestLoadBannerStandardSpace(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("loadBanner failed: %v", err)
	}
	space := ' '

	actual, ok := banner[space]
	if !ok {
		t.Fatalf("banner does not contain space character")
	}
	if len(actual) != 8 {
		t.Fatalf("expected 8 lines for space, got %d", len(actual))
	}
	expected := []string{
		"",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// TestLoadBannerShadowA verifies that the letter 'A' loads correctly
// from the shadow banner file.
func TestLoadBannerShadowA(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	ch := 'A'
	actual, ok := banner[ch]
	if !ok {
		t.Fatalf("banner does not contain 'A' character")
	}
	if len(actual) != 8 {
		t.Fatalf("expected 8 lines for 'A', got %d", len(actual))
	}

	expected := []string{
		"",
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"_|_|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
	}

	for i, line := range actual {
		if line != expected[i] {
			t.Errorf("line %d: expected %q, got %q", i, expected[i], line)
		}
	}
}

// TestLoadBannerEmptyFile verifies that LoadBanner returns an error
// when given an empty banner file (0 lines).
func TestLoadBannerEmptyFile(t *testing.T) {
	_, err := LoadBanner("../testdata/empty.txt")
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}

// TestLoadBannerCorruptedFile verifies that LoadBanner returns an error
// when given a corrupted banner file with too few lines.
func TestLoadBannerCorruptedFile(t *testing.T) {
	_, err := LoadBanner("../testdata/corrupted.txt")
	if err == nil {
		t.Error("expected error for corrupted file, got nil")
	}
}

// TestLoadBannerOversizedFile verifies that LoadBanner returns an error
// when given an oversized banner file with too many lines.
func TestLoadBannerOversizedFile(t *testing.T) {
	_, err := LoadBanner("../testdata/oversized.txt")
	if err == nil {
		t.Error("expected error for oversized file, got nil")
	}
}

// TestLoadBannerThinkertoy verifies that the thinkertoy banner file
// loads correctly with all 95 ASCII characters (32-126).
func TestLoadBannerThinkertoy(t *testing.T) {
	banner, err := LoadBanner("../testdata/thinkertoy.txt")
	if err != nil {
		t.Fatalf("thinkertoy failed: %v", err)
	}
	if len(banner) != totalChars {
		t.Errorf("expected %d chars, got %d", totalChars, len(banner))
	}
}

// TestLoadBannerNumbers verifies that all digits 0-9 load correctly
// from the standard banner file with exactly 8 lines each.
func TestLoadBannerNumbers(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	for r := '0'; r <= '9'; r++ {
		lines, ok := banner[r]
		if !ok {
			t.Errorf("missing digit %c", r)
			continue
		}
		if len(lines) != linesPerGlyph {
			t.Errorf("digit %c has %d lines, expected %d",
				r, len(lines), linesPerGlyph)
		}
	}
}

// TestLoadBannerCompleteCharacterSet verifies that all 95 printable
// ASCII characters (32-126) load correctly with exactly 8 lines each.
func TestLoadBannerCompleteCharacterSet(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	if len(banner) != totalChars {
		t.Fatalf("expected %d chars, got %d", totalChars, len(banner))
	}

	for r := firstPrintable; r <= lastPrintable; r++ {
		lines, ok := banner[r]
		if !ok {
			t.Errorf("missing char %c (ASCII %d)", r, r)
			continue
		}
		if len(lines) != linesPerGlyph {
			t.Errorf("char %c (ASCII %d) has %d lines, expected %d",
				r, r, len(lines), linesPerGlyph)
		}
	}
}

// TestLoadBannerAllSpecialCharacters verifies that all special characters
// (punctuation, symbols, etc.) load correctly from the banner file.
func TestLoadBannerAllSpecialCharacters(t *testing.T) {
	banner, err := LoadBanner("../testdata/standard.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	// Test all ASCII special characters
	specials := `!"#$%&'()*+,-./:;<=>?@[\]^_{|}~` + "`"

	for _, ch := range specials {
		lines, ok := banner[ch]
		if !ok {
			t.Errorf("missing special character %q (ASCII %d)", ch, ch)
			continue
		}
		if len(lines) != linesPerGlyph {
			t.Errorf("special char %q has %d lines, expected %d",
				ch, len(lines), linesPerGlyph)
		}
	}
}
