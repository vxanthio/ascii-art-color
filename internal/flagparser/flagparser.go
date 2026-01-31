package flagparser

import (
	"errors"
	"strconv"
	"strings"
)

const (
	minimumArgs = 2
	maximumArgs = 5
)

func ParseArgs(args []string) error {
	count := 0
	if len(args) < minimumArgs {
		return errors.New("error")
	}
	if len(args) > maximumArgs {
		return errors.New("error")
	}
	if err := validateColorFlag(args); err != nil {
		return err
	}
	for i, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			count++
			if i != 1 {
				return errors.New("error")
			}
			if count > 1 {
				return errors.New("error")
			}
		}
	}
	if strings.HasPrefix(args[1], "--color=") && len(args) < 3 {
		return errors.New("error")
	}
	if strings.HasPrefix(args[1], "--color=") {
		checkColorInTheFlag := strings.Split(args[1], "=")
		if len(checkColorInTheFlag) > 1 && checkColorInTheFlag[1] == "" {
			return errors.New("error")
		}
		color := checkColorInTheFlag[1]
		if color != "" {
			if err := validColors(color); err != nil {
				return err
			}
		}
	}
	return nil

}
func validateColorFlag(args []string) error {
	isItAFlag := strings.HasPrefix(args[1], "-")
	if isItAFlag {

		firstTwoLetters := strings.HasPrefix(args[1], "--")

		if !firstTwoLetters {
			return errors.New("error")
		}

		hasEqual := strings.Contains(args[1], "=")
		if !hasEqual {
			return errors.New("error")
		}

	}
	return nil
}
func validColors(color string) error {
	allowedColors := map[string]bool{
		"red":     true,
		"green":   true,
		"yellow":  true,
		"orange":  true,
		"blue":    true,
		"magenta": true,
	}
	_, exists := allowedColors[color]
	if !exists {
		if strings.HasPrefix(color, "rgb(") {
			inner := strings.TrimSuffix(strings.TrimPrefix(color, "rgb("), ")")
			separatedText := strings.Split(inner, ",")
			if len(separatedText) != 3 {
				return errors.New("error")
			}
			for i := 0; i < len(separatedText); i++ {
				digits, err := strconv.Atoi(separatedText[i])
				if err != nil {
					return err
				}
				if digits < 0 || digits > 255 {

					return errors.New("error")
				}
			}
		}
	}

	return nil
}
