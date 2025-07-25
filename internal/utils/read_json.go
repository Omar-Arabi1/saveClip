package utils

import (
	"encoding/json"
	"os"

	"github.com/Omar-Arabi1/saveClip/internal/models"
)

func ReadJson() ([]models.Clip, error) {
	clipsPath := CreateDataFile()
	data, err := os.ReadFile(clipsPath)

	if err != nil {
		return []models.Clip{}, err
	}

	var clips []models.Clip

	err = json.Unmarshal(data, &clips)

	if err != nil {
		return []models.Clip{}, err
	}

	return clips, err
}
