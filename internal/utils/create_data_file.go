package utils

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func CreateDataFile() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)
	}

	clipsPath := filepath.Join(homeDir, ".clips.json")

	_, err = os.Stat(clipsPath)

	if !errors.Is(err, os.ErrNotExist) {
		return clipsPath
	}

	var file *os.File
	file, err = os.Create(clipsPath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	os.WriteFile(clipsPath, []byte("[]"), 0664)
	return clipsPath
}
