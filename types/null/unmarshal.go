package null

import "encoding/json"

// Unmarshaler for NullInt64
func (i *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &i.Int64)
	i.Valid = (err == nil)
	return err
}

// Unmarshaler for NullFloat64
func (f *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &f.Float64)
	f.Valid = (err == nil)
	return err
}

// Unmarshaler for NullBool
func (bb *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &bb.Bool)
	bb.Valid = (err == nil)
	return err
}

// Unmarshaler for NullString
func (s *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &s.String)
	s.Valid = (err == nil)
	return err
}
