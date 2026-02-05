// Package flagparser validates command-line arguments for the ascii-art-color program.
//
// Its responsibility is limited to validating input format and structure.
// It does NOT perform rendering or color validation logic.
//
// It ensures:
//   - the correct number of arguments is provided
//   - the --color flag (if present) appears in the correct position
//   - only one --color flag is used
//   - the --color flag contains a non-empty value
//
// Any invalid input results in a usage error.
package flagparser

import (
	"errors"
	"strings"
)

// Argument count boundaries according to the project specification.
const (
	minimumArgs = 2
	maximumArgs = 5
)

// errUsage is the single user-facing error returned for any invalid CLI input.
// This keeps command-line output consistent and predictable.
var errUsage = errors.New("Usage: go run . [OPTION] [STRING]")

// ParseArgs validates the provided command-line arguments.
func ParseArgs(args []string) error {
	colorFlagCount := 0

	// Validate argument count boundaries.
	if len(args) < minimumArgs || len(args) > maximumArgs {
		return errUsage
	}

	// Any flag-like argument must be the --color flag.
	if strings.HasPrefix(args[1], "-") && !strings.HasPrefix(args[1], "--color=") {
		return errUsage
	}

	// Scan arguments to detect the --color flag and enforce its position.
	for i := 1; i < len(args); i++ {
		if strings.HasPrefix(args[i], "--color=") {
			colorFlagCount++

			// Only one --color flag is allowed.
			if colorFlagCount > 1 {
				return errUsage
			}

			// The --color flag must appear as the second argument.
			if i != 1 {
				return errUsage
			}
		}
	}

	// If --color is provided, a string to color must follow.
	if strings.HasPrefix(args[1], "--color=") && len(args) < 3 {
		return errUsage
	}

	// Only validate that the color value is non-empty.
	if strings.HasPrefix(args[1], "--color=") {
		_, color, found := strings.Cut(args[1], "=")
		if !found || color == "" {
			return errUsage
		}

	}
	return nil

}
