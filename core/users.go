// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
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
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	*types.User
}

// ReturnUser returns *core.User based off a *types.User
func ReturnUser(u *types.User) *User {
	return &User{u}
}

// SelectUser returns the User based on the user's ID.
func SelectUser(id int64) (*User, error) {
	var user User
	err := usersDB().First(&user, id)
	return &user, err.Error
}

// SelectUser returns the User based on the user's username
func SelectUsername(username string) (*User, error) {
	var user User
	res := usersDB().Where("username = ?", username)
	err := res.First(&user)
	return &user, err.Error
}

// Delete will remove the user record from the database
func (u *User) Delete() error {
	return usersDB().Delete(u).Error
}

// Update will update the user's record in database
func (u *User) Update() error {
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	return usersDB().Update(u).Error
}

// Create will insert a new user into the database
func (u *User) Create() (int64, error) {
	u.CreatedAt = time.Now()
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	db := usersDB().Create(u)
	if db.Error != nil {
		return 0, db.Error
	}
	if u.Id == 0 {
		utils.Log(3, fmt.Sprintf("Failed to create user %v. %v", u.Username, db.Error))
		return 0, db.Error
	}
	return u.Id, db.Error
}

// SelectAllUsers returns all users
func SelectAllUsers() ([]*User, error) {
	var users []*User
	db := usersDB().Find(&users)
	if db.Error != nil {
		utils.Log(3, fmt.Sprintf("Failed to load all users. %v", db.Error))
	}
	return users, db.Error
}

// AuthUser will return the User and a boolean if authentication was correct.
// AuthUser accepts username, and password as a string
func AuthUser(username, password string) (*User, bool) {
	user, err := SelectUsername(username)
	if err != nil {
		utils.Log(2, err)
		return nil, false
	}
	if CheckHash(password, user.Password) {
		return user, true
	}
	return nil, false
}

// CheckHash returns true if the password matches with a hashed bcrypt password
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
