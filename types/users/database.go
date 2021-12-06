package users

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
)

var (
	db  database.Database
	log = utils.Log.WithField("type", "user")
)

func SetDB(database database.Database) {
	db = database.Model(&User{})
}

func (u *User) AfterFind() {
	metrics.Query("user", "find")
}

func (u *User) AfterCreate() {
	metrics.Query("user", "create")
}

func (u *User) AfterUpdate() {
	metrics.Query("user", "update")
}

func (u *User) AfterDelete() {
	metrics.Query("user", "delete")
}

func Find(id int64) (*User, error) {
	var user User
	q := db.Where("id = ?", id).Find(&user)
	return &user, q.Error()
}

func FindByUsername(username string) (*User, error) {
	var user User
	q := db.Where("username = ?", username).Find(&user)
	return &user, q.Error()
}

func FindByAPIKey(key string) (*User, error) {
	var user User
	q := db.Where("api_key = ?", key).Find(&user)
	return &user, q.Error()
}

func All() []*User {
	var users []*User
	db.Find(&users)
	return users
}

func (u *User) Create() error {
	q := db.Create(u)
	if db.Error() == nil {
		log.Warnf("User #%d (%s) has been created", u.Id, u.Username)
	}
	return q.Error()
}

func (u *User) Update() error {
	q := db.Update(u)
	return q.Error()
}

func (u *User) Delete() error {
	q := db.Delete(u)
	if db.Error() == nil {
		log.Warnf("User #%d (%s) has been deleted", u.Id, u.Username)
	}
	return q.Error()
}
