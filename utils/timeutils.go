package utils

import (
	"fmt"
	"strconv"
)

// FormatTime formats seconds into time string 'HH:mm:ss'
func FormatTime(seconds int) string {
	var (
		hour   = toTimeString(seconds / 3600)
		minute = toTimeString(seconds % 3600 / 60)
		second = toTimeString(seconds % 60)
	)
	return fmt.Sprintf("%s:%s:%s", hour, minute, second)
}

func toTimeString(value int) string {
	if value < 10 {
		return fmt.Sprintf("0%d", value)
	}
	return strconv.Itoa(value)
}
