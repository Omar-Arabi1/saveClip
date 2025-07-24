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
	
	err = os.WriteFile("clipboard_contents.json", marshaledData, 0644)
	
	if err != nil {
		return err
	}
	
	return err
}