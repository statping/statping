package groups

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var example = &Group{
	Id:     1,
	Name:   "Example Group",
	Public: null.NewNullBool(true),
	Order:  1,
}

var s1 = &services.Service{
	Name:    "Example Service",
	Public:  null.NewNullBool(true),
	Order:   1,
	GroupId: 1,
}

var s2 = &services.Service{
	Name:    "Example Service 2",
	Public:  null.NewNullBool(true),
	Order:   2,
	GroupId: 1,
}

func TestInit(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.CreateTable(&Group{}, &services.Service{})
	db.Create(&example)
	db.Create(&s1)
	db.Create(&s2)
	SetDB(db)
	services.SetDB(db)
}

func TestFind(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Example Group", item.Name)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
}

func TestCreate(t *testing.T) {
	example := &Group{
		Name:   "Example 2",
		Public: null.NewNullBool(false),
		Order:  3,
	}
	err := example.Create()
	require.Nil(t, err)
	assert.NotZero(t, example.Id)
	assert.Equal(t, "Example 2", example.Name)
	assert.NotZero(t, example.CreatedAt)
}

func TestUpdate(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	item.Name = "Updated"
	item.Order = 1
	err = item.Update()
	require.Nil(t, err)
	assert.Equal(t, "Updated", item.Name)
}

func TestSelectGroups(t *testing.T) {
	groups := SelectGroups(true, true)
	assert.Len(t, groups, 2)

	groups = SelectGroups(false, false)
	assert.Len(t, groups, 1)

	groups = SelectGroups(false, true)
	assert.Len(t, groups, 2)

	assert.Equal(t, "Updated", groups[0].Name)
	assert.Equal(t, "Example 2", groups[1].Name)
}

func TestDelete(t *testing.T) {
	all := All()
	assert.Len(t, all, 2)

	item, err := Find(1)
	require.Nil(t, err)

	err = item.Delete()
	require.Nil(t, err)

	all = All()
	assert.Len(t, all, 1)
}

func TestClose(t *testing.T) {
	assert.Nil(t, db.Close())
}
