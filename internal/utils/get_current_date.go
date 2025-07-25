package utils

import (
	"fmt"
	"time"
)

func GetCurrentDate() string {
	var currentTime time.Time = time.Now()

	var currentYear int = currentTime.Year()
	var currentMonth string = fmt.Sprintf("%02d", int(currentTime.Month()))
	var currentDay string = fmt.Sprintf("%02d", currentTime.Day())

	return fmt.Sprintf("%v-%v-%v", currentYear, currentMonth, currentDay)
}
