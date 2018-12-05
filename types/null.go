// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package types

import (
	"database/sql"
	"encoding/json"
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

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// MarshalJSON for NullFloat64
func (ni *NullFloat64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Float64)
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// Unmarshaler for NullInt64
func (nf *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Int64)
	nf.Valid = (err == nil)
	return err
}

// Unmarshaler for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return err
}

// Unmarshaler for NullBool
func (nf *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Bool)
	nf.Valid = (err == nil)
	return err
}

// Unmarshaler for NullString
func (nf *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.String)
	nf.Valid = (err == nil)
	return err
}
