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
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	spaceChar := ' '
	expected := []string{
		"      ",
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
	_, err := LoadBanner("../../cmd/ascii-art/testdata/nope.txt")
	if err == nil {
		t.Errorf("expected error for missing file, got nil")
	}
}

// TestLoadBannerExclamationChar verifies that the exclamation mark
// loads correctly from the shadow banner file.
func TestLoadBannerExclamationChar(t *testing.T) {
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner Failed: %v", err)
	}
	char := '!'
	expected := []string{
		"   ",
		"_| ",
		"_| ",
		"_| ",
		"   ",
		"_| ",
		"   ",
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
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/standard.txt")
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
		"      ",
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
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/shadow.txt")
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
		"         ",
		"  _|_|   ",
		"_|    _| ",
		"_|_|_|_| ",
		"_|    _| ",
		"_|    _| ",
		"         ",
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
	_, err := LoadBanner("../../cmd/ascii-art/testdata/empty.txt")
	if err == nil {
		t.Error("expected error for empty file, got nil")
	}
}

// TestLoadBannerCorruptedFile verifies that LoadBanner returns an error
// when given a corrupted banner file with too few lines.
func TestLoadBannerCorruptedFile(t *testing.T) {
	_, err := LoadBanner("../../cmd/ascii-art/testdata/corrupted.txt")
	if err == nil {
		t.Error("expected error for corrupted file, got nil")
	}
}

// TestLoadBannerOversizedFile verifies that LoadBanner returns an error
// when given an oversized banner file with too many lines.
func TestLoadBannerOversizedFile(t *testing.T) {
	_, err := LoadBanner("../../cmd/ascii-art/testdata/oversized.txt")
	if err == nil {
		t.Error("expected error for oversized file, got nil")
	}
}

// TestLoadBannerThinkertoy verifies that the thinkertoy banner file
// loads correctly with all 95 ASCII characters (32-126).
func TestLoadBannerThinkertoy(t *testing.T) {
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/thinkertoy.txt")
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
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/standard.txt")
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
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/standard.txt")
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

// TestCharWidths verifies that CharWidths returns the correct column width
// for each character in the input text based on the banner glyph data.
func TestCharWidths(t *testing.T) {
	// Build a minimal banner with known glyph widths.
	// All 8 lines of each glyph have the same width (banner format guarantee).
	banner := Banner{
		'H': {"_    _ ", "_|  |_ ", "_|  |_ ", "|_  _| ", " |  |  ", " |  |  ", "       ", "       "},
		'i': {"   ", "   ", " _ ", "| |", "| |", "|_|", "   ", "   "},
		' ': {"      ", "      ", "      ", "      ", "      ", "      ", "      ", "      "},
		'!': {"_ ", "| ", "| ", "| ", "  ", "| ", "  ", "  "},
	}

	tests := []struct {
		name string
		text string
		want []int
	}{
		{
			name: "single char",
			text: "H",
			want: []int{7},
		},
		{
			name: "multiple chars",
			text: "Hi",
			want: []int{7, 3},
		},
		{
			name: "with space",
			text: "H i",
			want: []int{7, 6, 3},
		},
		{
			name: "empty string",
			text: "",
			want: []int{},
		},
		{
			name: "unknown char defaults to zero",
			text: "Z",
			want: []int{0},
		},
		{
			name: "repeated chars same width",
			text: "HHH",
			want: []int{7, 7, 7},
		},
		{
			name: "mixed known and unknown",
			text: "HZi",
			want: []int{7, 0, 3},
		},
		{
			name: "special char",
			text: "Hi!",
			want: []int{7, 3, 2},
		},
		{
			name: "spaces only",
			text: "   ",
			want: []int{6, 6, 6},
		},
		{
			name: "single space",
			text: " ",
			want: []int{6},
		},
		{
			name: "single narrow char",
			text: "!",
			want: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CharWidths(tt.text, banner)
			if len(got) != len(tt.want) {
				t.Fatalf("CharWidths(%q) returned %d widths, want %d", tt.text, len(got), len(tt.want))
			}
			for i, w := range got {
				if w != tt.want[i] {
					t.Errorf("CharWidths(%q)[%d] = %d, want %d", tt.text, i, w, tt.want[i])
				}
			}
		})
	}
}

// TestCharWidths_EmptyBanner verifies CharWidths handles an empty banner map gracefully.
func TestCharWidths_EmptyBanner(t *testing.T) {
	banner := Banner{}
	widths := CharWidths("Hello", banner)

	if len(widths) != 5 {
		t.Fatalf("expected 5 widths, got %d", len(widths))
	}
	for i, w := range widths {
		if w != 0 {
			t.Errorf("widths[%d] = %d, want 0 for empty banner", i, w)
		}
	}
}

// TestCharWidths_RealBanners verifies CharWidths with actual loaded banner data
// from all three banner files (standard, shadow, thinkertoy).
func TestCharWidths_RealBanners(t *testing.T) {
	bannerFiles := []struct {
		name string
		path string
	}{
		{"standard", "../../cmd/ascii-art/testdata/standard.txt"},
		{"shadow", "../../cmd/ascii-art/testdata/shadow.txt"},
		{"thinkertoy", "../../cmd/ascii-art/testdata/thinkertoy.txt"},
	}

	for _, bf := range bannerFiles {
		t.Run(bf.name, func(t *testing.T) {
			banner, err := LoadBanner(bf.path)
			if err != nil {
				t.Fatalf("LoadBanner(%s) failed: %v", bf.name, err)
			}

			text := "Hello World! 123 @#$"
			widths := CharWidths(text, banner)

			if len(widths) != len(text) {
				t.Fatalf("expected %d widths, got %d", len(text), len(widths))
			}

			for i, ch := range text {
				glyph := banner[ch]
				if glyph == nil {
					t.Errorf("char %c missing from %s banner", ch, bf.name)
					continue
				}
				expected := len(glyph[0])
				if widths[i] != expected {
					t.Errorf("width[%d] (%c) = %d, want %d", i, ch, widths[i], expected)
				}
			}
		})
	}
}

// TestCharWidths_ConsistentGlyphLines verifies that all 8 lines of each glyph
// have the same width, which is the assumption CharWidths relies on.
func TestCharWidths_ConsistentGlyphLines(t *testing.T) {
	bannerFiles := []struct {
		name string
		path string
	}{
		{"standard", "../../cmd/ascii-art/testdata/standard.txt"},
		{"shadow", "../../cmd/ascii-art/testdata/shadow.txt"},
		{"thinkertoy", "../../cmd/ascii-art/testdata/thinkertoy.txt"},
	}

	for _, bf := range bannerFiles {
		t.Run(bf.name, func(t *testing.T) {
			banner, err := LoadBanner(bf.path)
			if err != nil {
				t.Fatalf("LoadBanner(%s) failed: %v", bf.name, err)
			}

			for r := firstPrintable; r <= lastPrintable; r++ {
				glyph := banner[r]
				if glyph == nil {
					t.Errorf("char %c missing", r)
					continue
				}
				width := len(glyph[0])
				for lineIdx, line := range glyph {
					if len(line) != width {
						t.Errorf("%s: char %c (ASCII %d) line %d width %d != line 0 width %d",
							bf.name, r, r, lineIdx, len(line), width)
					}
				}
			}
		})
	}
}

// TestLoadBannerAllSpecialCharacters verifies that all special characters
// (punctuation, symbols, etc.) load correctly from the banner file.
func TestLoadBannerAllSpecialCharacters(t *testing.T) {
	banner, err := LoadBanner("../../cmd/ascii-art/testdata/standard.txt")
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
