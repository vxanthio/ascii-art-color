/*
Package color parses color specifications into ANSI 24-bit terminal escape codes.

Supported formats:
  - Named colors: black, red, green, yellow, blue, magenta, cyan, white,
    orange, purple, pink, brown, gray (case-insensitive)
  - Hex: #RRGGBB (e.g. #ff0000)
  - RGB: rgb(R, G, B) (e.g. rgb(255, 0, 0))

Example:

	rgb, _ := Parse("red")
	fmt.Print(color.ANSI(rgb) + "Hello" + "\033[0m")
*/
package color

import (
	"testing"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name      string
		colorSpec string
		want      RGB
		wantErr   bool
	}{
		// Named colors
		{"named_red", "red", RGB{255, 0, 0}, false},
		{"named_green", "green", RGB{0, 255, 0}, false},
		{"named_blue", "blue", RGB{0, 0, 255}, false},
		{"named_case_insensitive", "RED", RGB{255, 0, 0}, false},
		{"named_unknown", "blurple", RGB{}, true},

		// Extended names
		{"named_orange", "orange", RGB{255, 165, 0}, false},
		{"named_purple", "purple", RGB{128, 0, 128}, false},
		{"named_pink", "pink", RGB{255, 192, 203}, false},
		{"named_brown", "brown", RGB{165, 42, 42}, false},
		{"named_gray", "gray", RGB{128, 128, 128}, false},

		// Hex
		{"hex_red", "#ff0000", RGB{255, 0, 0}, false},
		{"hex_invalid_length_short", "#ff0", RGB{}, true},
		{"hex_invalid_length_long", "#ff000000", RGB{}, true},
		{"hex_invalid_chars", "#gg0000", RGB{}, true},

		// RGB
		{"rgb_red", "rgb(255, 0, 0)", RGB{255, 0, 0}, false},
		{"rgb_spaces", "rgb( 255 , 0 , 0 )", RGB{255, 0, 0}, false},
		{"rgb_invalid_count", "rgb(255)", RGB{}, true},
		{"rgb_out_of_range", "rgb(300, 0, 0)", RGB{}, true},
		{"rgb_non_number", "rgb(a, 0, 0)", RGB{}, true},

		// Empty / whitespace
		{"empty_spec", "", RGB{}, true},
		{"whitespace_spec", "   ", RGB{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.colorSpec)
			if (err != nil) != tt.wantErr {
				t.Fatalf(`Parse(%q) error = %v, wantErr %t`, tt.colorSpec, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf(`Parse(%q) = %#v, want %#v`, tt.colorSpec, got, tt.want)
			}
		})
	}
}

func Test_ANSI(t *testing.T) {
	tests := []struct {
		name string
		rgb  RGB
		want string
	}{
		{"red", RGB{255, 0, 0}, "\033[38;2;255;0;0m"},
		{"green", RGB{0, 255, 0}, "\033[38;2;0;255;0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ANSI(tt.rgb)
			if got != tt.want {
				t.Fatalf("ANSI(%#v) = %q, want %q", tt.rgb, got, tt.want)
			}

		})
	}
}
