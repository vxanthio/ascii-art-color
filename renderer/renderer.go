package renderer

import "strings"

func rendererASCII(input string, banner map[rune][]string) string {
	result := ""
	parts := strings.Split(input, "\n")
	if len(parts) > 0 {
		if parts[len(parts)-1] == "" {
			parts = parts[0 : len(parts)-1]
		}
	}
	if len(parts) == 1 && parts[0] == "" {
		return result
	}
	for p, line := range parts {

		for i := 0; i < 8; i++ {
			for _, ch := range line {
				rows := banner[ch]

				result += rows[i]

			}
			if i != 7 {
				result += "\n"
			}

		}
		if p != len(parts)-1 {
			result += "\n"
		}
	}
	return result
}
