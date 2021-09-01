package actor

import (
	"errors"
	"strconv"
)

func CheckNotEmpty(input string) error {
	if input == "" {
		return errors.New("input should not be empty")
	}
	return nil
}

func CheckIsAPositiveNumber(input string) error {
	if n, err := strconv.Atoi(input); err != nil {
		return err
	} else if n < 0 {
		return errors.New("the number cant be native")
	}
	return nil
}
