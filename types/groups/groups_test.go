package groups

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGroup_Create(t *testing.T) {
	group := &Group{
		Name: "Testing",
	}
	err := group.Create()
	assert.Nil(t, err)
	assert.NotZero(t, group.Id)
}

func TestGroup_Services(t *testing.T) {
	group, err := Find(1)
	require.Nil(t, err)
	assert.NotEmpty(t, group.Services())
}

func TestSelectGroups(t *testing.T) {
	grs := SelectGroups(true, false)
	assert.Equal(t, int(3), len(grs))
	grs = SelectGroups(true, true)
	assert.Equal(t, int(5), len(grs))
}
