package null

import "encoding/json"

// MarshalJSON for NullInt64
func (i NullInt64) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(i.Int64)
}

// MarshalJSON for NullFloat64
func (f NullFloat64) MarshalJSON() ([]byte, error) {
	if !f.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(f.Float64)
}

// MarshalJSON for NullBool
func (bb NullBool) MarshalJSON() ([]byte, error) {
	if !bb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(bb.Bool)
}

// MarshalJSON for NullString
func (s NullString) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(s.String)
}
