package core

import (
	"github.com/hunterlong/statup/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := &types.User{
		Username: "hunter",
		Password: "password123",
		Email:    "test@email.com",
		Admin:    true,
	}
	userId, err := CreateUser(user)
	assert.Nil(t, err)
	assert.NotZero(t, userId)
}

func TestSelectAllUsers(t *testing.T) {
	users, err := SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(users))
}

func TestSelectUser(t *testing.T) {
	user, err := SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "test@email.com", user.Email)
	assert.True(t, user.Admin)
}

func TestUpdateUser(t *testing.T) {
	user, err := SelectUser(1)
	assert.Nil(t, err)

	user.Username = "updated"

	err = UpdateUser(user)
	assert.Nil(t, err)

	updatedUser, err := SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "updated", updatedUser.Username)
}
