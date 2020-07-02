package users

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var example = &User{
	Username: "example_user",
	Email:    "Description here",
	Password: "password123",
	Admin:    null.NewNullBool(true),
}

func TestInit(t *testing.T) {
	err := utils.InitLogs()
	require.Nil(t, err)
	db, err := database.OpenTester()
	require.Nil(t, err)
	db.CreateTable(&User{})
	db.Create(&example)
	SetDB(db)
}

func TestFind(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "example_user", item.Username)
	assert.NotEmpty(t, item.ApiKey)
	assert.NotEqual(t, "password123", item.Password)
	assert.True(t, item.Admin.Bool)
}

func TestFindByUsername(t *testing.T) {
	item, err := FindByUsername("example_user")
	require.Nil(t, err)
	assert.Equal(t, "example_user", item.Username)
	assert.NotEmpty(t, item.ApiKey)
	assert.NotEqual(t, "password123", item.Password)
	assert.True(t, item.Admin.Bool)
}

func TestAll(t *testing.T) {
	items := All()
	assert.Len(t, items, 1)
}

func TestCreate(t *testing.T) {
	example := &User{
		Username: "exampleuser2",
		Password: "password12345",
		Email:    "info@yahoo.com",
	}
	err := example.Create()
	require.Nil(t, err)
	assert.NotZero(t, example.Id)
	assert.Equal(t, "exampleuser2", example.Username)
	assert.NotEqual(t, "password12345", example.Password)
	assert.NotZero(t, example.CreatedAt)
	assert.NotEmpty(t, example.ApiKey)
}

func TestUpdate(t *testing.T) {
	item, err := Find(1)
	require.Nil(t, err)
	item.Username = "updated_user"
	err = item.Update()
	require.Nil(t, err)
	assert.Equal(t, "updated_user", item.Username)
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
