package users

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
	"gorm.io/gorm"
)

var (
	db  *database.Database
	log = utils.Log.WithField("type", "user")
)

func SetDB(dbz *database.Database) {
	db = database.Wrap(dbz.Model(&User{}))
}

func (u *User) AfterFind(*gorm.DB) error {
	metrics.Query("user", "find")
	return nil
}

func (u *User) AfterCreate(*gorm.DB) error {
	metrics.Query("user", "create")
	return nil
}

func (u *User) AfterUpdate(*gorm.DB) error {
	metrics.Query("user", "update")
	return nil
}

func (u *User) AfterDelete(*gorm.DB) error {
	metrics.Query("user", "delete")
	return nil
}

func Find(id int64) (*User, error) {
	var user User
	q := db.Where("id = ?", id).Find(&user)
	return &user, q.Error
}

func FindByUsername(username string) (*User, error) {
	var user User
	q := db.Where("username = ?", username).Find(&user)
	return &user, q.Error
}

func FindByAPIKey(key string) (*User, error) {
	var user User
	q := db.Where("api_key = ?", key).Find(&user)
	return &user, q.Error
}

func All() []*User {
	var users []*User
	db.Find(&users)
	return users
}

func (u *User) Create() error {
	q := db.Create(u)
	if db.Error == nil {
		log.Warnf("User #%d (%s) has been created", u.Id, u.Username)
	}
	return q.Error
}

func (u *User) Update() error {
	q := db.Save(u)
	return q.Error
}

func (u *User) Delete() error {
	q := db.Delete(u)
	if db.Error == nil {
		log.Warnf("User #%d (%s) has been deleted", u.Id, u.Username)
	}
	return q.Error
}
