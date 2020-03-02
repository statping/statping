package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestFixedTime(t *testing.T) {

	timeVal, err := time.Parse("2006-01-02T15:04:05Z", "2020-05-22T06:02:13Z")
	require.Nil(t, err)

	examples := []struct {
		Time     time.Time
		Duration time.Duration
		Expected string
	}{{
		timeVal,
		time.Second,
		"2020-05-22T06:02:13Z",
	}, {
		timeVal,
		time.Minute,
		"2020-05-22T06:02:00Z",
	}, {
		timeVal,
		time.Hour,
		"2020-05-22T06:00:00Z",
	}, {
		timeVal,
		Day,
		"2020-05-22T00:00:00Z",
	}, {
		timeVal.Add(2 * Month),
		Month,
		"2020-07-01T00:00:00Z",
	}, {
		timeVal.Add(2 * Year),
		Year,
		"2022-01-01T00:00:00Z",
	}}

	for _, e := range examples {
		assert.Equal(t, e.Expected, FixedTime(e.Time, e.Duration), fmt.Sprintf("reformating for: %v %v", e.Time, e.Duration))
	}

}
