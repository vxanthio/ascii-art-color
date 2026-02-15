package main

import (
	"testing"
)

func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{"./ascii-art"}

	_, _, err := ParseArgs(args)

	if err == nil {
		t.Error("Expected error for no arguments, got nil")
	}

	expectedMsg := "usage: go run . \"text\" [banner]"
	if err.Error() != expectedMsg {
		t.Errorf("Expected error message: %q, got: %q", expectedMsg, err.Error())
	}
}

func TestParseArgs_TextOnly(t *testing.T) {
	args := []string{"./ascii-art", "Hello"}

	text, banner, err := ParseArgs(args)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if text != "Hello" {
		t.Errorf("Expected text: 'Hello', got: %q", text)
	}

	if banner != "standard" {
		t.Errorf("Expected banner: 'standard', got: %q", banner)
	}
}

func TestParseArgs_TextAndBanner(t *testing.T) {
	args := []string{"./ascii-art", "Hello", "shadow"}

	text, banner, err := ParseArgs(args)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if text != "Hello" {
		t.Errorf("Expected text: 'Hello', got: %q", text)
	}

	if banner != "shadow" {
		t.Errorf("Expected banner: 'shadow', got: %q", banner)
	}
}

func TestParseArgs_TooManyArguments(t *testing.T) {
	args := []string{"./ascii-art", "Hello", "shadow", "extra"}

	_, _, err := ParseArgs(args)

	if err == nil {
		t.Error("Expected error for too many arguments, got nil")
	}
}

func TestParseArgs_AllBannerTypes(t *testing.T) {
	testCases := []struct {
		args           []string
		expectedBanner string
	}{
		{[]string{"prog", "Hi", "standard"}, "standard"},
		{[]string{"prog", "Hi", "shadow"}, "shadow"},
		{[]string{"prog", "Hi", "thinkertoy"}, "thinkertoy"},
	}

	for _, tc := range testCases {
		_, banner, err := ParseArgs(tc.args)

		if err != nil {
			t.Errorf("Args %v: expected no error, got: %v", tc.args, err)
		}

		if banner != tc.expectedBanner {
			t.Errorf("Args %v: expected banner %q, got: %q",
				tc.args, tc.expectedBanner, banner)
		}
	}
}

func TestParseArgs_EmptyStringText(t *testing.T) {
	args := []string{"./ascii-art", ""}

	text, banner, err := ParseArgs(args)

	if err != nil {
		t.Errorf("Expected no error for empty string, got: %v", err)
	}

	if text != "" {
		t.Errorf("Expected empty text, got: %q", text)
	}

	if banner != "standard" {
		t.Errorf("Expected banner: 'standard', got: %q", banner)
	}
}

func TestHasColorFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{"no args", []string{"prog"}, false},
		{"text only", []string{"prog", "hello"}, false},
		{"text and banner", []string{"prog", "hello", "shadow"}, false},
		{"color flag present", []string{"prog", "--color=red", "hello"}, true},
		{"color flag with substring", []string{"prog", "--color=red", "sub", "hello"}, true},
		{"wrong format dash", []string{"prog", "-color=red", "hello"}, false},
		{"wrong format colon", []string{"prog", "--color:red", "hello"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasColorFlag(tt.args)
			if got != tt.want {
				t.Errorf("hasColorFlag(%v) = %v, want %v", tt.args, got, tt.want)
			}
		})
	}
}

func TestExtractColorArgs(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantColor string
		wantSub   string
		wantText  string
		wantBnr   string
		wantErr   bool
	}{
		{
			name:      "flag and text only",
			args:      []string{"prog", "--color=red", "hello"},
			wantColor: "red", wantSub: "", wantText: "hello", wantBnr: "standard",
		},
		{
			name:      "flag text and banner",
			args:      []string{"prog", "--color=red", "hello", "shadow"},
			wantColor: "red", wantSub: "", wantText: "hello", wantBnr: "shadow",
		},
		{
			name:      "flag substring and text",
			args:      []string{"prog", "--color=red", "sub", "hello"},
			wantColor: "red", wantSub: "sub", wantText: "hello", wantBnr: "standard",
		},
		{
			name:      "flag substring text and banner",
			args:      []string{"prog", "--color=red", "sub", "hello", "thinkertoy"},
			wantColor: "red", wantSub: "sub", wantText: "hello", wantBnr: "thinkertoy",
		},
		{
			name:      "audit case orange GuYs",
			args:      []string{"prog", "--color=orange", "GuYs", "HeY GuYs"},
			wantColor: "orange", wantSub: "GuYs", wantText: "HeY GuYs", wantBnr: "standard",
		},
		{
			name:      "audit case blue B",
			args:      []string{"prog", "--color=blue", "B", "RGB()"},
			wantColor: "blue", wantSub: "B", wantText: "RGB()", wantBnr: "standard",
		},
		{
			name:      "hex color",
			args:      []string{"prog", "--color=#ff0000", "hello"},
			wantColor: "#ff0000", wantSub: "", wantText: "hello", wantBnr: "standard",
		},
		{
			name:      "rgb color",
			args:      []string{"prog", "--color=rgb(255,0,0)", "hello"},
			wantColor: "rgb(255,0,0)", wantSub: "", wantText: "hello", wantBnr: "standard",
		},
		{
			name:      "newline in text",
			args:      []string{"prog", "--color=red", "hello\\nworld"},
			wantColor: "red", wantSub: "", wantText: "hello\nworld", wantBnr: "standard",
		},
		{
			name:    "missing text after flag",
			args:    []string{"prog", "--color=red"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colorSpec, sub, text, bnr, err := extractColorArgs(tt.args)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if colorSpec != tt.wantColor {
				t.Errorf("color = %q, want %q", colorSpec, tt.wantColor)
			}
			if sub != tt.wantSub {
				t.Errorf("substring = %q, want %q", sub, tt.wantSub)
			}
			if text != tt.wantText {
				t.Errorf("text = %q, want %q", text, tt.wantText)
			}
			if bnr != tt.wantBnr {
				t.Errorf("banner = %q, want %q", bnr, tt.wantBnr)
			}
		})
	}
}

func TestGetBannerPath_ValidBanners(t *testing.T) {
	testCases := []struct {
		banner       string
		expectedPath string
	}{
		{"standard", "testdata/standard.txt"},
		{"shadow", "testdata/shadow.txt"},
		{"thinkertoy", "testdata/thinkertoy.txt"},
	}

	for _, tc := range testCases {
		path, err := GetBannerPath(tc.banner)

		if err != nil {
			t.Errorf("Banner %q: expected no error, got: %v", tc.banner, err)
		}

		if path != tc.expectedPath {
			t.Errorf("Banner %q: expected path %q, got: %q",
				tc.banner, tc.expectedPath, path)
		}
	}
}

func TestGetBannerPath_InvalidBanner(t *testing.T) {
	banner := "invalid"

	_, err := GetBannerPath(banner)

	if err == nil {
		t.Error("Expected error for invalid banner, got nil")
	}
}
