package utils

import (
	"errors"

	"github.com/Omar-Arabi1/saveClip/internal/models"
)

func AreClipsEmpty(clips []models.Clip) error {
	var err error
	
	if len(clips) == 0 {
		err = errors.New("your list is empty use the save command to add something to it")
	}
	return err
}