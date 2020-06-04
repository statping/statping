package incidents

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var example = &Incident{
	Title:       "Example",
	Description: "No description",
	ServiceId:   1,
}

var update1 = &IncidentUpdate{
	IncidentId: 1,
	Message:    "First one here",
	Type:       "update",
}

var update2 = &IncidentUpdate{
	IncidentId: 1,
	Message:    "Second one here",
	Type:       "update",
}

func TestInit(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.AutoMigrate(&Incident{}, &IncidentUpdate{})
	db.Create(&example)
	SetDB(db)
}

func TestFind(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Example", item.Title)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
}

func TestCreate(t *testing.T) {
	example := &Incident{
		Title: "Example 2",
	}
	err := example.Create()
	require.Nil(t, err)
	assert.NotZero(t, example.Id)
	assert.Equal(t, "Example 2", example.Title)
	assert.NotZero(t, example.CreatedAt)
}

func TestUpdate(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	item.Title = "Updated"
	err = item.Update()
	require.Nil(t, err)
	assert.Equal(t, "Updated", item.Title)
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
