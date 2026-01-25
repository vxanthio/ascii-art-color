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
		name string
		spec string
		want RGB
	}{
		{"red", "red", RGB{255, 0, 0}},
		{"green", "green", RGB{0, 255, 0}},
		{"blue", "blue", RGB{0, 0, 255}},
		{"case insensitive", "RED", RGB{255, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.spec)
			if err != nil {
				t.Fatalf(`Parse("%s") error = %v`, tt.spec, err)
			}
			if got != tt.want {
				t.Fatalf(`Parse("%s") = %#v, want %#v`, tt.spec, got, tt.want)
			}
		})
	}
}
