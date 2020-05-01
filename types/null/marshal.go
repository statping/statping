package null

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

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

// MarshalYAML for NullInt64
func (i NullInt64) MarshalYAML() (interface{}, error) {
	if !i.Valid {
		return 0, nil
	}
	return yaml.Marshal(i.Int64)
}

// MarshalYAML for NullFloat64
func (f NullFloat64) MarshalYAML() (interface{}, error) {
	if !f.Valid {
		return 0.0, nil
	}
	return yaml.Marshal(f.Float64)
}

// MarshalYAML for NullBool
func (bb NullBool) MarshalYAML() (interface{}, error) {
	if !bb.Valid {
		return false, nil
	}
	return yaml.Marshal(bb.Bool)
}

// MarshalYAML for NullString
func (s NullString) MarshalYAML() (interface{}, error) {
	if !s.Valid {
		return "", nil
	}
	return yaml.Marshal(s.String)
}
