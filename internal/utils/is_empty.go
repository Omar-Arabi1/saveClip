package utils

import (
	"errors"
	"strings"
)

func IsEmpty(text string) error {
	var err error

	if strings.TrimSpace(text) == "" {
		err = errors.New("invalid label")
	}
	return err
}
