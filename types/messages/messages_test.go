package messages

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var example = &Message{
	Title:       "Example Message",
	Description: "Description here",
	StartOn:     utils.Now().Add(10 * time.Minute),
	EndOn:       utils.Now().Add(15 * time.Minute),
	ServiceId:   1,
}

func TestInit(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.CreateTable(&Message{})
	db.Create(&example)
	SetDB(db)
}

func TestFind(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Example Message", item.Title)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
}

func TestCreate(t *testing.T) {
	example := &Message{
		Title:       "Example 2",
		Description: "New Message here",
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
