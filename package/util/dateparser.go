package util

import (
	"library/package/log"
	"time"
)

func DateParser(dateString string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		log.Errorf("error parsing date %v", dateString)
	}
	return parsedDate, nil
}

// PurseTime converts string time int golang time in a given format
func PurseTime(format string, timeString string) time.Time {
	newTime := time.Time{}
	newTime, _ = time.Parse(format, timeString)
	return newTime
}
