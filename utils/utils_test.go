package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInitLogs(t *testing.T) {
	assert.Nil(t, InitLogs())
}

func TestIntString(t *testing.T) {
	assert.Equal(t, "1", IntString(1))
}

func TestStringInt(t *testing.T) {
	assert.Equal(t, int64(1), StringInt("1"))
}

func TestTimestamp_Ago(t *testing.T) {
	now := Timestamp(time.Now())
	assert.Equal(t, "Just now", now.Ago())
}

func TestUnderScoreString(t *testing.T) {
	assert.Equal(t, "this_is_a_test", UnderScoreString("this is a test"))
}

func TestHashPassword(t *testing.T) {
	assert.Equal(t, 60, len(HashPassword("password123")))
}

func TestNewSHA1Hash(t *testing.T) {
	assert.NotEmpty(t, NewSHA1Hash(5))
}

func TestRandomString(t *testing.T) {
	assert.NotEmpty(t, RandomString(5))
}

func TestSha256(t *testing.T) {
	assert.Equal(t, "dc724af18fbdd4e59189f5fe768a5f8311527050", Sha256([]byte("testing")))
}
