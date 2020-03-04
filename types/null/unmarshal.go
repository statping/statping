package null

import "encoding/json"

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
