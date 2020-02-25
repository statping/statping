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
	"fmt"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	*database.UserObj
}

// ReturnUser returns *core.User based off a *types.User
func uwrap(u *database.UserObj) *User {
	return &User{u}
}

// CountUsers returns the amount of users
func CountUsers() int64 {
	var amount int64
	Database(&User{}).Count(&amount)
	return amount
}

// SelectUser returns the User based on the User's ID.
func SelectUser(id int64) (*User, error) {
	user, err := database.User(id)
	if err != nil {
		return nil, err
	}
	return uwrap(user), err
}

// SelectUsername returns the User based on the User's username
func SelectUsername(username string) (*User, error) {
	user, err := database.UserByUsername(username)
	if err != nil {
		return nil, err
	}
	return uwrap(user), err
}

// Delete will remove the User record from the database
func (u *User) Delete() error {
	return database.Delete(u)
}

// Update will update the User's record in database
func (u *User) Update() error {
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	return database.Update(u)
}

// Create will insert a new User into the database
func (u *User) Create() (int64, error) {
	u.CreatedAt = time.Now().UTC()
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)

	user, err := database.Create(u)
	if err != nil {
		return 0, err
	}
	if user.Id == 0 {
		log.Errorln(fmt.Sprintf("Failed to create User %v. %v", u.Username, err))
		return 0, err
	}
	return u.Id, err
}

// SelectAllUsers returns all users
func SelectAllUsers() ([]*User, error) {
	var users []*User
	err := database.AllUsers(&users)
	if err != nil {
		log.Errorln(fmt.Sprintf("Failed to load all users. %v", err))
		return nil, err
	}
	return users, err
}

// AuthUser will return the User and a boolean if authentication was correct.
// AuthUser accepts username, and password as a string
func AuthUser(username, password string) (*User, bool) {
	user, err := SelectUsername(username)
	if err != nil {
		log.Warnln(fmt.Errorf("user %v not found", username))
		return nil, false
	}
	if CheckHash(password, user.Password) {
		user.UpdatedAt = time.Now().UTC()
		user.Update()
		return user, true
	}
	return nil, false
}

// CheckHash returns true if the password matches with a hashed bcrypt password
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
