package null

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
