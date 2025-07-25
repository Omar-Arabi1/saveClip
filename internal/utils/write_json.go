package utils

import (
	"encoding/json"
	"os"

	"github.com/Omar-Arabi1/saveClip/internal/models"
)

func WriteJson(data []models.Clip) error {
	marshaledData, err := json.MarshalIndent(data, "", "  ")
	
	if err != nil {
		return err
	}
	
	clipsPath := CreateDataFile()
	err = os.WriteFile(clipsPath, marshaledData, 0644)
	
	if err != nil {
		return err
	}
	
	return err
}