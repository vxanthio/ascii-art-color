package flagparser

import (
	"errors"
	"strings"
)

func ParseArgs(args []string) error {
	if len(args) < 2 {
		return errors.New("error")
	}
	if len(args) > 5 {
		return errors.New("error")
	}
	if err := validateColorFlag(args); err != nil {
		return err
	}

	return nil

}
func validateColorFlag(args []string) error {
	isItAFlag := strings.HasPrefix(args[1], "-")
	if isItAFlag == true {

		firstTwoLetters := strings.HasPrefix(args[1], "--")

		if firstTwoLetters == false {
			return errors.New("error")
		}

		equalExistance := strings.Contains(args[1], "=")
		if equalExistance == false {
			return errors.New("error")
		}

	}
	return nil
}
