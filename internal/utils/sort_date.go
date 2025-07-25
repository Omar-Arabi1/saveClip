package utils

import (
	"sort"

	"github.com/Omar-Arabi1/saveClip/internal/models"
)

func SortDate(data []models.Clip, newest bool) {
	if newest {
		sort.SliceStable(data, func(i, j int) bool { return data[i].CreationDate > data[j].CreationDate })
	} else {
		sort.SliceStable(data, func(i, j int) bool { return data[i].CreationDate < data[j].CreationDate })
	}
}
