package color

import (
	"fmt"
	"strconv"
	"strings"
)

// RGB represents a 24-bit color.
type RGB struct {
	R, G, B uint8
}

var namedColors = map[string]RGB{
	"black":   {0, 0, 0},
	"red":     {255, 0, 0},
	"green":   {0, 255, 0},
	"yellow":  {255, 255, 0},
	"blue":    {0, 0, 255},
	"magenta": {255, 0, 255},
	"cyan":    {0, 255, 255},
	"white":   {255, 255, 255},
	"orange":  {255, 165, 0},
	"purple":  {128, 0, 128},
	"pink":    {255, 192, 203},
	"brown":   {165, 42, 42},
	"gray":    {128, 128, 128},
}

// Parse converts a color specification string into RGB.
func Parse(spec string) (RGB, error) {
	lower := strings.ToLower(spec)

	if color, ok := namedColors[lower]; ok {
		return color, nil
	}
	if len(spec) == 7 && spec[0] == '#' {
		return parseHex(spec)
	}
	if strings.HasPrefix(lower, "rgb(") {
		return parseRGB(spec)
	}
	return RGB{}, fmt.Errorf("unknown color format %q", spec)
}

func parseHex(hex string) (RGB, error) {
	if len(hex) != 7 || hex[0] != '#' {
		return RGB{}, fmt.Errorf("invalid hex format: %q", hex)
	}

	r, err := strconv.ParseUint(hex[1:3], 16, 8)
	if err != nil {
		return RGB{}, fmt.Errorf("invalid red hex: %v", err)
	}

	g, err := strconv.ParseUint(hex[3:5], 16, 8)
	if err != nil {
		return RGB{}, fmt.Errorf("invalid green hex: %v", err)
	}
	b, err := strconv.ParseUint(hex[5:7], 16, 8)
	if err != nil {
		return RGB{}, fmt.Errorf("invalid blue hex: %v", err)
	}

	return RGB{uint8(r), uint8(g), uint8(b)}, nil
}

func parseRGB(rgbStr string) (RGB, error) {
	content := strings.TrimPrefix(rgbStr, "rgb(")
	content = strings.TrimSuffix(content, ")")

	parts := strings.Split(content, ",")
	if len(parts) != 3 {
		return RGB{}, fmt.Errorf("rgb() requires exactly 3 components, got %d", len(parts))
	}
	var r, g, b uint8
	for i, part := range parts {
		valStr := strings.TrimSpace(part)
		val, err := strconv.ParseUint(valStr, 10, 8)
		if err != nil {
			return RGB{}, fmt.Errorf("invalid rgb() component %q: %v", valStr, err)
		}
		if val > 255 {
			return RGB{}, fmt.Errorf("rgb() component %q out of range (0-255)", valStr)
		}
		switch i {
		case 0:
			r = uint8(val)
		case 1:
			g = uint8(val)
		case 2:
			b = uint8(val)
		}
	}

	return RGB{r, g, b}, nil
}
