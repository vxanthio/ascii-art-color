package parser

import (
	"testing"
)

// testing spaces in shadow file
func TestLoadBannerSpaceChar(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner failed: %v", err)
	}

	space := ' '
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

	actual, ok := banner[space]
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

// missing file test
func TestLoadBannerMissingFile(t *testing.T) {
	_, err := LoadBanner("../testdata/nope.txt")
	if err == nil {
		t.Errorf("expected error for missing file, got nil")
	}
}

// testing exclamation
func TestLoadBannerExclamationChar(t *testing.T) {
	banner, err := LoadBanner("../testdata/shadow.txt")
	if err != nil {
		t.Fatalf("LoadBanner Failed: %v", err)
	}
	ex := '!'
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
	actual, ok := banner[ex]
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

// testing spaces in standard file
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

// testing 'A' in shadow file
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
