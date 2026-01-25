package color

import (
	"fmt"
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
	return RGB{}, fmt.Errorf("unknown color nameL %q", spec)
}
