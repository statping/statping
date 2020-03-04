package users

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/utils"
	"time"
)

func Find(id int64) (*User, error) {
	var user *User
	db := database.DB().Model(&User{}).Where("id = ?", id).Find(&user)
	return user, db.Error()
}

func FindByUsername(username string) (*User, error) {
	var user *User
	db := database.DB().Model(&User{}).Where("username = ?", username).Find(&user)
	return user, db.Error()
}

func All() []*User {
	var users []*User
	database.DB().Model(&User{}).Find(&users)
	return users
}

func (u *User) Create() error {
	u.CreatedAt = time.Now().UTC()
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)

	db := database.DB().Create(&u)
	return db.Error()
}

func (u *User) Update() error {
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	db := database.DB().Update(&u)
	return db.Error()
}

func (u *User) Delete() error {
	db := database.DB().Delete(&u)
	return db.Error()
}
