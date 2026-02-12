package util

import (
	"library/package/log"
	"time"
)

func TimeParser(format, timestring string) (time.Time, error) {
	timeToReturn, err := time.Parse(format, timestring)
	if IsError(err) {
		log.Errorf("error formating date: %v", err)
	}
	return timeToReturn, err
}
