package flagparser_test

import (
	"ascii-art-color/internal/flagparser"
	"testing"
)

func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{"program"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("Error was expected")
	}

}
func TestParseArgs_TooManyArgs(t *testing.T) {
	args := []string{"program",
		"banner",
		"--color=red",
		"substring",
		"some text",
		"EXTRA"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("Error too many args")
	}
}
func TestParseArgs_InvalidColorPrefix(t *testing.T) {
	args := []string{"program",
		"-color:black",
		"some text",
	}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error invalid color flag prefix")
	}
}
func TestParseArgs_FormatColor(t *testing.T) {
	args := []string{
		"program",
		"--color:red",
		"some text",
	}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("exprected error invalid color format")
	}
}
func TestParseArgs_SingleStringAllowed(t *testing.T) {
	args := []string{
		"program",
		"text",
	}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
func TestParseArgs_FlagAndStringAllowed(t *testing.T) {
	args := []string{"program", "--color=red", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error:%v", err)
	}
}

func TestParseArgs_ColorSubstringAllowed(t *testing.T) {
	args := []string{"program", "--color=red", "text", "substring"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error:%v", err)
	}
}
func TestParseArgs_MissingString(t *testing.T) {
	args := []string{"program", "--color=red"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("Usage: go run . [OPTION] [STRING]")
	}
}
func TestParseArgs_MultipleFlags(t *testing.T) {
	args := []string{"program", "--color=red", "--color=blue", "text"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error")
	}
}
func TestParseArgs_InvalidPositionForColorFlag(t *testing.T) {
	args := []string{"program", "text", "--color=red"}
	err := flagparser.ParseArgs(args)
	if err == nil {
		t.Errorf("expected error")
	}
}
func TestParseArgs_SubstringMissingWhileStringExists(t *testing.T) {
	args := []string{"program", "--color=red", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error")
	}
}
func TestParseArgs_ValidRGBColor(t *testing.T) {
	args := []string{"program", "--color=rgb(255,0,0)", "text"}
	err := flagparser.ParseArgs(args)
	if err != nil {
		t.Errorf("unexpected error")
	}
}
