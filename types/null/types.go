package null

import (
	"database/sql"
)

// NewNullString returns a sql.NullString for JSON parsing
func NewNullString(s string) NullString {
	return NullString{sql.NullString{s, true}}
}

// NewNullBool returns a sql.NullBool for JSON parsing
func NewNullBool(s bool) NullBool {
	return NullBool{sql.NullBool{s, true}}
}

// NewNullInt64 returns a sql.NullInt64 for JSON parsing
func NewNullInt64(s int64) NullInt64 {
	return NullInt64{sql.NullInt64{s, true}}
}

// NewNullFloat64 returns a sql.NullFloat64 for JSON parsing
func NewNullFloat64(s float64) NullFloat64 {
	return NullFloat64{sql.NullFloat64{s, true}}
}

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}
