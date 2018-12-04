// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/hunterlong/statping/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := ReturnUser(&types.User{
		Username: "hunter",
		Password: "password123",
		Email:    "test@email.com",
		Admin:    types.NewNullBool(true),
	})
	userId, err := user.Create()
	assert.Nil(t, err)
	assert.NotZero(t, userId)
}

func TestSelectAllUsers(t *testing.T) {
	users, err := SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestSelectUser(t *testing.T) {
	user, err := SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "info@betatude.com", user.Email)
	assert.True(t, user.Admin.Bool)
}

func TestSelectUsername(t *testing.T) {
	user, err := SelectUsername("hunter")
	assert.Nil(t, err)
	assert.Equal(t, "test@email.com", user.Email)
	assert.Equal(t, int64(3), user.Id)
	assert.True(t, user.Admin.Bool)
}

func TestUpdateUser(t *testing.T) {
	user, err := SelectUser(1)
	assert.Nil(t, err)
	user.Username = "updated"
	err = user.Update()
	assert.Nil(t, err)
	updatedUser, err := SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "updated", updatedUser.Username)
}

func TestCreateUser2(t *testing.T) {
	user := ReturnUser(&types.User{
		Username: "hunterlong",
		Password: "password123",
		Email:    "User@email.com",
		Admin:    types.NewNullBool(true),
	})
	userId, err := user.Create()
	assert.Nil(t, err)
	assert.NotZero(t, userId)
}

func TestSelectAllUsersAgain(t *testing.T) {
	users, err := SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func TestAuthUser(t *testing.T) {
	user, auth := AuthUser("hunterlong", "password123")
	assert.True(t, auth)
	assert.NotNil(t, user)
	assert.Equal(t, "User@email.com", user.Email)
	assert.Equal(t, int64(4), user.Id)
	assert.True(t, user.Admin.Bool)
}

func TestFailedAuthUser(t *testing.T) {
	user, auth := AuthUser("hunterlong", "wrongpassword")
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
	err = user.Delete()
	assert.Nil(t, err)
}

func TestDbConfig_Close(t *testing.T) {
	err := DbSession.Close()
	assert.Nil(t, err)
}
