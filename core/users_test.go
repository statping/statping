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

func TestSelectUsername(t *testing.T) {
	user, err := SelectUsername("hunter")
	assert.Nil(t, err)
	assert.Equal(t, "test@email.com", user.Email)
	assert.Equal(t, int64(1), user.Id)
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

func TestCreateUser2(t *testing.T) {
	user := &types.User{
		Username: "hunterlong",
		Password: "password123",
		Email:    "user@email.com",
		Admin:    true,
	}
	userId, err := CreateUser(user)
	assert.Nil(t, err)
	assert.NotZero(t, userId)
}

func TestSelectAllUsersAgain(t *testing.T) {
	users, err := SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestAuthUser(t *testing.T) {
	user, auth := AuthUser("hunterlong", "password123")
	assert.True(t, auth)
	assert.NotNil(t, user)
	assert.Equal(t, "user@email.com", user.Email)
	assert.Equal(t, int64(2), user.Id)
	assert.True(t, user.Admin)
}

func TestFailedAuthUser(t *testing.T) {
	user, auth := AuthUser("hunter", "wrongpassword")
	assert.False(t, auth)
	assert.Nil(t, user)
}

func TestCheckPassword(t *testing.T) {
	user, err := SelectUser(2)
	assert.Nil(t, err)
	pass := CheckHash("password123", user.Password)
	assert.True(t, pass)
}

func TestDeleteUser(t *testing.T) {
	user, err := SelectUser(2)
	assert.Nil(t, err)
	err = DeleteUser(user)
	assert.Nil(t, err)
}
