package database

import (
	"fmt"
	"time"
)

type TimeGroup interface {
}

func (d *Database) ParseTime(t string) (time.Time, error) {
	switch d.Dialector.Name() {
	case "mysql":
		return time.Parse("2006-01-02T15:04:05Z", t)
	case "postgres":
		return time.Parse("2006-01-02T15:04:05Z", t)
	default:
		return time.Parse("2006-01-02 15:04:05", t)
	}
}

// FormatTime returns the timestamp in the same format as the DATETIME column in database
func (d *Database) FormatTime(t time.Time) string {
	switch d.Dialector.Name() {
	case "postgres":
		return t.Format("2006-01-02 15:04:05.999999999")
	default:
		return t.Format("2006-01-02 15:04:05")
	}
}

// SelectByTime returns an SQL query that will group "created_at" column by x seconds and returns as "timeframe"
func (d *Database) SelectByTime(increment time.Duration) string {
	seconds := int64(increment.Seconds())
	switch d.Dialector.Name() {
	case "mysql":
		return fmt.Sprintf("FROM_UNIXTIME(FLOOR(UNIX_TIMESTAMP(created_at) / %d) * %d) AS timeframe", seconds, seconds)
	case "postgres":
		return fmt.Sprintf("date_trunc('minute', created_at) - (CAST(EXTRACT(MINUTE FROM created_at) AS integer) %% %d) * interval '1 minute' AS timeframe", seconds)
	default:
		return fmt.Sprintf("datetime((strftime('%%s', created_at) / %d) * %d, 'unixepoch') as timeframe", seconds, seconds)
	}
}

func (d *Database) Since(ago time.Time) *Database {
	return Wrap(d.Where("created_at > ?", d.FormatTime(ago)))
}

func (d *Database) Between(t1 time.Time, t2 time.Time) *Database {
	return Wrap(d.Where("created_at BETWEEN ? AND ?", d.FormatTime(t1), d.FormatTime(t2)))
}
