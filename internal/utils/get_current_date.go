package utils

import (
	"fmt"
	"time"
)

func GetCurrentDate() string {
	var currentTime time.Time = time.Now()

	var currentYear int = currentTime.Year()
	var currentMonth int = int(currentTime.Month())
	var currentDay int = currentTime.Day()
	
	var currentTimeFormatted string = fmt.Sprintf("%d-%d-%d", currentYear, currentMonth, currentDay)

	return currentTimeFormatted
}
