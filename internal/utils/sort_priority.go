package utils

import (
	"sort"

	"github.com/Omar-Arabi1/saveClip/internal/models"
)

func SortPriority(data []models.Clip, highest bool) {
	if highest {
		sort.SliceStable(data, func(i, j int) bool { return data[i].Priority > data[j].Priority })
	} else {
		sort.SliceStable(data, func(i, j int) bool { return data[i].Priority < data[j].Priority })
	}
}