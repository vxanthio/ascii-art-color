// Package renderer contains unit tests for the ASCII renderer.
// These tests verify that RendererASCII correctly converts input strings
// into their ASCII-art representations using a provided banner.
package renderer

import (
	"testing"
)

// TestEmptyInput verifies that an empty input string
// produces no output.
func TestEmptyInput(t *testing.T) {
	input := ""
	banner := map[rune][]string{}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)

	}
	if input != output {
		t.Errorf("expected:\n%q\ngot:\n%q", input, output)

	}
}

// TestSingleCharacter verifies rendering of a single character
// with a banner height of 8 lines
func TestSingleCharacter(t *testing.T) {
	input := "A"
	expected := `A1
A2
A3
A4
A5
A6
A7
A8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)

	}
}

// TestMultipleCharacters verifies that multiple characters
// are rendered horizontally on the same ASCII-art rows.
func TestMultipleCharacters(t *testing.T) {
	input := "AB"
	expected := `A1B1
A2B2
A3B3
A4B4
A5B5
A6B6
A7B7
A8B8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'B': {"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)

	}

}

// TestSpaceBetweenCharacters verifies that spaces between characters
// are correctly rendered using the space entry in the banner.
func TestSpaceBetweenCharacters(t *testing.T) {
	input := "A A"
	expected := `A1  A1
A2  A2
A3  A3
A4  A4
A5  A5
A6  A6
A7  A7
A8  A8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		' ': {"  ", "  ", "  ", "  ", "  ", "  ", "  ", "  "},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)

	}

}

// TestNumbersBetweenCharacters verifies that numeric characters
// are rendered correctly when mixed with letters.
func TestNumbersBetweenCharacters(t *testing.T) {
	input := "A1A"
	expected := `A11A1
A21A2
A31A3
A41A4
A51A5
A61A6
A71A7
A81A8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'1': {"1", "1", "1", "1", "1", "1", "1", "1"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)

	}
}

// TestSpecialCharacters verifies rendering of special characters
// that exist in the banner map.
func TestSpecialCharacters(t *testing.T) {
	input := "{}"
	expected := `{}
{}
{}
{}
{}
{}
{}
{}`
	banner := map[rune][]string{
		'{': {"{", "{", "{", "{", "{", "{", "{", "{"},
		'}': {"}", "}", "}", "}", "}", "}", "}", "}"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)

	}
}

// TestNewlineBetweenCharacters verifies that a newline in the input
// separates the output into multiple ASCII-art blocks.
func TestNewlineBetweenCharacters(t *testing.T) {
	input := "A\nB"
	expected := `A1
A2
A3
A4
A5
A6
A7
A8
B1
B2
B3
B4
B5
B6
B7
B8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'B': {"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)

	}

}

// TestTrailingNewline verifies that a trailing newline
// does not produce an extra empty ASCII block.
func TestTrailingNewline(t *testing.T) {
	input := "A\n"
	expected := `A1
A2
A3
A4
A5
A6
A7
A8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)
	}
}

// TestConsecutiveNewlines verifies that consecutive newlines
// create visible separation between ASCII-art blocks,
// matching the behavior shown in the official examples.
func TestConsecutiveNewlines(t *testing.T) {
	input := "A\n\nB"
	expected := `A1
A2
A3
A4
A5
A6
A7
A8

B1
B2
B3
B4
B5
B6
B7
B8`
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
		'B': {"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8"},
	}
	output, err := RendererASCII(input, banner)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if expected != output {
		t.Errorf("expected:\n%q\ngot:\n%q", expected, output)
	}
}
func TestMissigCharaster(t *testing.T) {
	input := "AB"
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A5", "A6", "A7", "A8"},
	}
	output, err := RendererASCII(input, banner)
	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != "" {
		t.Errorf("expected empty output on error, got %q", output)

	}

}
func TestCorruptedBanner(t *testing.T) {
	input := "A"
	banner := map[rune][]string{
		'A': {"A1", "A2", "A3", "A4", "A6", "A7", "A8"},
	}
	output, err := RendererASCII(input, banner)
	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != "" {
		t.Errorf("expected empty output on error, got %q", output)
	}
}
func TestEmptyBanner(t *testing.T) {
	input := "A"
	banner := map[rune][]string{}
	output, err := RendererASCII(input, banner)
	if err == nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if output != "" {
		t.Errorf("expected empty output on error, got %q", output)
	}
}
