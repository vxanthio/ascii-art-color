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
