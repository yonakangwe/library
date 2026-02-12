package time_parser

import (
	"library/package/log"
	"library/package/util"
	"fmt"
	"time"
)

func TimeParser(format, timestring string) (time.Time, error) {
	timeToReturn, err := time.Parse(format, timestring)
	if util.IsError(err) {
		log.Errorf("error formating date: %v", err)
	}
	return timeToReturn, err
}

func humanizeDuration(d time.Duration) string {
	seconds := int(d.Seconds())
	minutes := int(d.Minutes())
	hours := int(d.Hours())
	days := int(d.Hours() / 24)
	weeks := int(d.Hours() / (24 * 7))
	months := int(d.Hours() / (24 * 30))
	years := int(d.Hours() / (24 * 365))

	switch {
	case seconds < 60:
		return fmt.Sprintf("%d seconds ago", seconds)
	case minutes < 60:
		return fmt.Sprintf("%d minutes ago", minutes)
	case hours < 24:
		return fmt.Sprintf("%d hours ago", hours)
	case days < 7:
		return fmt.Sprintf("%d days ago", days)
	case weeks < 4:
		return fmt.Sprintf("%d weeks ago", weeks)
	case months < 12:
		return fmt.Sprintf("%d months ago", months)
	default:
		return fmt.Sprintf("%d years ago", years)
	}
}

func TimeDuration(timestamp time.Time) string {
	layout := "2006-01-02 15:04:05.000 -0700"
	t, err := time.Parse(layout, timestamp.Format("2006-01-02 15:04:05.000 -0700"))
	if err != nil {
		log.Errorf("Error parsing time:", err)
		return ""
	}

	now := time.Now()
	duration := now.Sub(t)
	return humanizeDuration(duration)
}
