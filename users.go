package main

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Id        int64     `db:"id,omitempty" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"-"`
	Email     string    `db:"email" json:"-"`
	ApiKey    string    `db:"api_key" json:"api_key"`
	ApiSecret string    `db:"api_secret" json:"-"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func SelectUser(id int64) (*User, error) {
	var user User
	col := dbSession.Collection("users")
	res := col.Find("id", id)
	err := res.One(&user)
	return &user, err
}

func SelectUsername(username string) (*User, error) {
	var user User
	col := dbSession.Collection("users")
	res := col.Find("username", username)
	err := res.One(&user)
	return &user, err
}

func (u *User) Create() (int64, error) {
	u.CreatedAt = time.Now()
	password := HashPassword(u.Password)
	u.Password = password
	u.ApiKey = NewSHA1Hash(5)
	u.ApiSecret = NewSHA1Hash(10)
	col := dbSession.Collection("users")
	uuid, err := col.Insert(u)
	if uuid == nil {
		return 0, err
	}
	return uuid.(int64), err
}

func SelectAllUsers() ([]User, error) {
	var users []User
	col := dbSession.Collection("users").Find()
	err := col.All(&users)
	return users, err
}

func AuthUser(username, password string) (*User, bool) {
	var auth bool
	user, err := SelectUsername(username)
	if err != nil {
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
