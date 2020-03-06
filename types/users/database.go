package users

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
	"github.com/prometheus/common/log"
	"time"
)

func DB() database.Database {
	return database.DB().Model(&User{})
}

func Find(id int64) (*User, error) {
	var user User
	db := DB().Where("id = ?", id).Find(&user)
	return &user, db.Error()
}

func FindByUsername(username string) (*User, error) {
	var user User
	db := DB().Where("username = ?", username).Find(&user)
	return &user, db.Error()
}

func All() []*User {
	var users []*User
	DB().Find(&users)
	return users
}

func (u *User) Create() error {
	u.CreatedAt = time.Now().UTC()
	u.Password = utils.HashPassword(u.Password)
	if u.ApiKey == "" || u.ApiSecret == "" {
		u.ApiKey = utils.NewSHA1Hash(16)
		u.ApiSecret = utils.NewSHA1Hash(16)
	}
	db := DB().Create(u)
	if db.Error() == nil {
		log.Warnf("User #%d (%s) has been created", u.Id, u.Username)
	}
	return db.Error()
}

func (u *User) Update() error {
	//u.ApiKey = utils.NewSHA1Hash(5)
	//u.ApiSecret = utils.NewSHA1Hash(10)
	db := DB().Update(u)
	return db.Error()
}

func (u *User) Delete() error {
	db := DB().Delete(u)
	if db.Error() == nil {
		log.Warnf("User #%d (%s) has been deleted")
	}
	return db.Error()
}
