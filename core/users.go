package core

import (
	"fmt"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User types.User

func SelectUser(id int64) (*types.User, error) {
	var user *types.User
	col := DbSession.Collection("users")
	res := col.Find("id", id)
	err := res.One(&user)
	return user, err
}

func SelectUsername(username string) (*types.User, error) {
	var user *types.User
	col := DbSession.Collection("users")
	res := col.Find("username", username)
	err := res.One(&user)
	return user, err
}

func DeleteUser(u *types.User) error {
	col := DbSession.Collection("users")
	user := col.Find("id", u.Id)
	return user.Delete()
}

func UpdateUser(u *types.User) error {
	u.CreatedAt = time.Now()
	col := DbSession.Collection("users")
	user := col.Find("id", u.Id)
	return user.Update(u)
}

func CreateUser(u *types.User) (int64, error) {
	u.CreatedAt = time.Now()
	u.Password = utils.HashPassword(u.Password)
	u.ApiKey = utils.NewSHA1Hash(5)
	u.ApiSecret = utils.NewSHA1Hash(10)
	col := DbSession.Collection("users")
	uuid, err := col.Insert(u)
	if uuid == nil {
		utils.Log(3, fmt.Sprintf("Failed to create user %v. %v", u.Username, err))
		return 0, err
	}
	return uuid.(int64), err
}

func SelectAllUsers() ([]User, error) {
	var users []User
	col := DbSession.Collection("users").Find()
	err := col.All(&users)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Failed to load all users. %v", err))
	}
	return users, err
}

func AuthUser(username, password string) (*types.User, bool) {
	var auth bool
	user, err := SelectUsername(username)
	if err != nil {
		utils.Log(2, err)
		return nil, false
	}
	if CheckHash(password, user.Password) {
		auth = true
	}
	return user, auth
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
