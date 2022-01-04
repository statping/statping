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

// UnmarshalYAML for NullInt64
func (i *NullInt64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val int64
	if err := unmarshal(&val); err != nil {
		return err
	}
	*i = NewNullInt64(val)
	return nil
}

// UnmarshalYAML for NullFloat64
func (f *NullFloat64) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val float64
	if err := unmarshal(&val); err != nil {
		return err
	}
	*f = NewNullFloat64(val)
	return nil
}

// UnmarshalYAML for NullBool
func (bb *NullBool) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val bool
	if err := unmarshal(&val); err != nil {
		return err
	}
	*bb = NewNullBool(val)
	return nil
}

// UnmarshalYAML for NullFloat64
func (s *NullString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var val string
	if err := unmarshal(&val); err != nil {
		return err
	}
	*s = NewNullString(val)
	return nil
}
