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
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
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

// CountUsers returns the amount of users
func CountUsers() int64 {
	var amount int64
	usersDB().Count(&amount)
	return amount
}

// SelectUser returns the User based on the User's ID.
func SelectUser(id int64) (*User, error) {
	var user User
	err := usersDB().Where("id = ?", id).First(&user)
	return &user, err.Error
}

// SelectUsername returns the User based on the User's username
func SelectUsername(username string) (*User, error) {
	var user User
	res := usersDB().Where("username = ?", username)
	err := res.First(&user)
	return &user, err.Error
}

// Delete will remove the User record from the database
func (u *User) Delete() error {
	return usersDB().Delete(u).Error
}

// Update will update the User's record in database
func (u *User) Update() error {
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	return usersDB().Update(u).Error
}

// Create will insert a new User into the database
func (u *User) Create() (int64, error) {
	u.CreatedAt = time.Now().UTC()
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	db := usersDB().Create(u)
	if db.Error != nil {
		return 0, db.Error
	}
	if u.Id == 0 {
		log.Errorln(fmt.Sprintf("Failed to create User %v. %v", u.Username, db.Error))
		return 0, db.Error
	}
	return u.Id, db.Error
}

// SelectAllUsers returns all users
func SelectAllUsers() ([]*User, error) {
	var users []*User
	db := usersDB().Find(&users)
	if db.Error != nil {
		log.Errorln(fmt.Sprintf("Failed to load all users. %v", db.Error))
		return nil, db.Error
	}
	return users, db.Error
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
		return user, true
	}
	return nil, false
}

// CheckHash returns true if the password matches with a hashed bcrypt password
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
