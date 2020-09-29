package utils

import (
	"github.com/hako/durafmt"
	"time"
)

// Now returns the UTC timestamp
func Now() time.Time {
	return time.Now().UTC()
}

type Duration struct {
	time.Duration
}

func (d Duration) Human() string {
	return durafmt.Parse(d.Duration).LimitFirstN(2).String()
}

// FormatDuration converts a time.Duration into a string
func FormatDuration(d time.Duration) string {
	return durafmt.ParseShort(d).LimitFirstN(3).String()
}
