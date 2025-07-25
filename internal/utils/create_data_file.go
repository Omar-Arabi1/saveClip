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

	clipsPath := filepath.Join(homeDir, "clips.json")
	clipsPathHidden := SetHidden(clipsPath)

	_, err = os.Stat(clipsPathHidden)

	if !errors.Is(err, os.ErrNotExist) {
		return clipsPathHidden
	}

	var file *os.File
	file, err = os.Create(clipsPathHidden)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	os.WriteFile(clipsPathHidden, []byte("[]"), 0664)
	return clipsPathHidden
}
