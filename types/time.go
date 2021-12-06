package types

import (
	"time"
)

const (
	TIME_NANO  = "2006-01-02T15:04:05Z"
	TIME       = "2006-01-02 15:04:05"
	CHART_TIME = "2006-01-02T15:04:05.999999-07:00"
	TIME_DAY   = "2006-01-02"
)

var (
	Second = time.Second
	Minute = time.Minute
	Hour   = time.Hour
	Day    = Hour * 24
	Week   = Day * 7
	Month  = Week * 4
	Year   = Day * 365
)

func FixedTime(t time.Time, d time.Duration) string {
	return t.Format(durationStr(d))
}

func durationStr(d time.Duration) string {

	switch m := d.Seconds(); {

	case m >= Month.Seconds():
		return "2006-01-01T00:00:00Z"

	case m >= Week.Seconds():
		return "2006-01-02T00:00:00Z"

	case m >= Day.Seconds():
		return "2006-01-02T00:00:00Z"

	case m >= Hour.Seconds():
		return "2006-01-02T15:00:00Z"

	case m >= Minute.Seconds():
		return "2006-01-02T15:04:00Z"

	default:
		return "2006-01-02T15:04:05Z"
	}
}
