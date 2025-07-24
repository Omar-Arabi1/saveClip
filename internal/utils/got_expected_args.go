package utils

import (
	"errors"
	"fmt"
)

func GotExpectedArgs(args []string, expectedArgsNum int) error {
	var err error

	if len(args) != expectedArgsNum {
		errMessage := fmt.Sprintf("expected %d arguments got %d", expectedArgsNum, len(args))
		err = errors.New(errMessage)
	}
	return err
}
