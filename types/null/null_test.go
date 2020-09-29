package null

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJSONMarshal(t *testing.T) {
	tests := []struct {
		Input        interface{}
		ExpectedJSON string
	}{
		{
			Input:        NewNullBool(true),
			ExpectedJSON: `true`,
		},
		{
			Input:        NewNullBool(false),
			ExpectedJSON: `false`,
		},
		{
			Input:        NewNullFloat64(0.994),
			ExpectedJSON: `0.994`,
		},
		{
			Input:        NewNullFloat64(0),
			ExpectedJSON: `0`,
		},
		{
			Input:        NewNullInt64(42),
			ExpectedJSON: `42`,
		},
		{
			Input:        NewNullInt64(0),
			ExpectedJSON: `0`,
		},
		{
			Input:        NewNullString("test"),
			ExpectedJSON: `"test"`,
		},
		{
			Input:        NewNullString(""),
			ExpectedJSON: `""`,
		},
	}

	for _, test := range tests {
		str, err := json.Marshal(test.Input)
		require.Nil(t, err)
		assert.Equal(t, test.ExpectedJSON, string(str))
	}
}

func TestNewNullBool(t *testing.T) {
	val := NewNullBool(true)
	assert.True(t, val.Bool)

	val = NewNullBool(false)
	assert.False(t, val.Bool)
}

func TestNewNullInt64(t *testing.T) {
	val := NewNullInt64(29)
	assert.Equal(t, int64(29), val.Int64)
}

func TestNewNullString(t *testing.T) {
	val := NewNullString("statping.com")
	assert.Equal(t, "statping.com", val.String)
}

func TestNewNullFloat64(t *testing.T) {
	val := NewNullFloat64(42.222)
	assert.Equal(t, float64(42.222), val.Float64)
}
