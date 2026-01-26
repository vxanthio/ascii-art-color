package color

import "testing"

func TestParseNamedColor(t *testing.T) {
	got, err := Parse("red")
	if err != nil {
		t.Fatalf("Parse(\"red\") returned error: %v", err)
	}

	want := RGB{R: 255, G: 0, B: 0}
	if got != want {
		t.Fatalf("Parse(\"red\") = %#v, want %#v", got, want)
	}
}

func TestParseNamedColors(t *testing.T) {
	tests := []struct {
		name    string
		spec    string
		want    RGB
		wantErr bool
	}{
		{"red", "red", RGB{255, 0, 0}, false},
		{"green", "green", RGB{0, 255, 0}, false},
		{"blue", "blue", RGB{0, 0, 255}, false},
		{"case insensitive", "RED", RGB{255, 0, 0}, false},
		{"unknown color", "blurple", RGB{}, true},
		{"hex red", "#ff0000", RGB{255, 0, 0}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.spec)
			if (err != nil) != tt.wantErr {
				t.Fatalf(`Parse(%q) error = %v, wantErr %t`, tt.spec, err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Fatalf(`Parse(%q) = %#v, want %#v`, tt.spec, got, tt.want)
			}
		})
	}
}
