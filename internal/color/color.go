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
