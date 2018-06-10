package main

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

func SelectUser(username string) User {
	var user User
	rows, err := db.Query("SELECT id, username, password FROM users WHERE username=$1", username)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}
	}
	return user
}

func (u *User) Create() int {
	password := HashPassword(u.Password)
	var lastInsertId int
	db.QueryRow("INSERT INTO users(username,password,created_at) VALUES($1,$2,NOW()) returning id;", u.Username, password).Scan(&lastInsertId)
	return lastInsertId
}

func SelectAllUsers() []User {
	var users []User
	rows, err := db.Query("SELECT id, username, password FROM users ORDER BY id ASC")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

func AuthUser(username, password string) (User, bool) {
	var user User
	var auth bool
	user = SelectUser(username)
	if CheckHash(password, user.Password) {
		auth = true
	}
	return user, auth
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
