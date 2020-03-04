package notifications

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertNotifierDB(t *testing.T) {
	if skipNewDb {
		t.SkipNow()
	}
	err := InsertNotifierDB()
	assert.Nil(t, err)
}
