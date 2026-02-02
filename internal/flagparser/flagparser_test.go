// Package flagparser_test contains unit tests for the flagparser package.
//
// These tests verify that ParseArgs correctly validates command-line arguments,
// including argument count, flag syntax, flag position, and supported color formats.
package flagparser_test

import (
	"ascii-art-color/internal/flagparser"
	"testing"
)

// TestParseArgs_NoArguments verifies that providing no arguments returns a usage error.
func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{"program"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when no arguments are provided")
	}

}

// TestParseArgs_TooManyArgs verifies that exceeding the maximum number of arguments returns an error.
func TestParseArgs_TooManyArgs(t *testing.T) {
	args := []string{"program", "banner", "--color=red", "substring", "some text", "EXTRA"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when too many arguments are provided")
	}
}

// TestParseArgs_InvalidColorPrefix verifies that flags not starting with '--' are rejected.
func TestParseArgs_InvalidColorPrefix(t *testing.T) {
	args := []string{"program", "-color:black", "some text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error for invalid flag format")
	}
}

// TestParseArgs_FormatColor verifies that a color flag without '=' is rejected.
func TestParseArgs_FormatColor(t *testing.T) {
	args := []string{"program", "--color:red", "some text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("exprected usage error for invalid --color format")
	}
}

// TestParseArgs_SingleStringAllowed verifies that a single string argument is valid.
func TestParseArgs_SingleStringAllowed(t *testing.T) {
	args := []string{
		"program",
		"text",
	}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid single string input: %v", err)
	}
}

// TestParseArgs_FlagAndStringAllowed verifies that a color flag followed by a string is valid.
func TestParseArgs_FlagAndStringAllowed(t *testing.T) {
	args := []string{"program", "--color=red", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected  error for valid color flag and string:%v", err)
	}
}

// TestParseArgs_ColorSubstringAllowed verifies that a color flag, string, and substring are valid.
func TestParseArgs_ColorSubstringAllowed(t *testing.T) {
	args := []string{"program", "--color=red", "text", "substring"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid color flag,string and substring:%v", err)
	}
}

// TestParseArgs_MissingString verifies that providing a color flag without a string returns an error.
func TestParseArgs_MissingString(t *testing.T) {
	args := []string{"program", "--color=red"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when color flag is provided without a string")
	}
}

// TestParseArgs_MultipleFlags verifies that multiple --color flags are rejected.
func TestParseArgs_MultipleFlags(t *testing.T) {
	args := []string{"program", "--color=red", "--color=blue", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when multiple color flags are provided")
	}
}

// TestParseArgs_InvalidPositionForColorFlag verifies that the color flag must be the second argument.
func TestParseArgs_InvalidPositionForColorFlag(t *testing.T) {
	args := []string{"program", "text", "--color=red"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected usage error when color flag is in an invalid position")
	}
}

// TestParseArgs_SubstringMissingWhileStringExists verifies that the substring argument is optional.
func TestParseArgs_SubstringMissingWhileStringExists(t *testing.T) {
	args := []string{"program", "--color=red", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error when substring is optional")
	}
}

// TestParseArgs_ValidRGBColor verifies that a valid RGB color format is accepted.
func TestParseArgs_ValidRGBColor(t *testing.T) {
	args := []string{"program", "--color=rgb(255,0,0)", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid RGB color:%v", err)
	}
}

// TestParseArgs_ValidHexColor verifies that a valid HEX color format is accepted.
func TestParseArgs_ValidHexColor(t *testing.T) {
	args := []string{"program", "--color=#ff0000", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error for valid HEX color:%v", err)
	}
}

// TestParseArgs_InvalidRGB_OutOfRange verifies that RGB values outside 0–255 are rejected.
func TestParseArgs_InvalidRGB_OutOfRange(t *testing.T) {
	args := []string{"program", "--color=rgb(300,0,0)", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error for RGB value out of range")
	}
}

// TestParseArgs_InvalidHexLength verifies that HEX colors with invalid length are rejected.
func TestParseArgs_InvalidHexLength(t *testing.T) {
	args := []string{"program", "--color=#123", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error for invalid HEX length")
	}
}

// TestParseArgs_InvalidColorName verifies that unsupported color names are rejected.
func TestParseArgs_InvalidColorName(t *testing.T) {
	args := []string{"program", "--color=pink", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error for unsupported color name")
	}
}
