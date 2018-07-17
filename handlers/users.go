package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"strconv"
)

func SessionUser(r *http.Request) *types.User {
	session, _ := Store.Get(r, COOKIE_KEY)
	if session == nil {
		return nil
	}
	uuid := session.Values["user_id"]
	var user *types.User
	col := core.DbSession.Collection("users")
	res := col.Find("id", uuid)
	err := res.One(&user)
	if err != nil {
		utils.Log(3, fmt.Sprintf("cannot fetch user %v", uuid))
		return nil
	}
	return user
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	users, _ := core.SelectAllUsers()
	ExecuteResponse(w, r, "users.html", users)
}

func UsersEditHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := core.SelectUser(int64(id))
	ExecuteResponse(w, r, "user.html", user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := core.SelectUser(int64(id))

	user.Username = r.PostForm.Get("username")
	user.Email = r.PostForm.Get("email")
	user.Admin = (r.PostForm.Get("admin") == "on")
	password := r.PostForm.Get("password")
	if password != "##########" {
		user.Password = utils.HashPassword(password)
	}
	core.UpdateUser(user)
	users, _ := core.SelectAllUsers()
	ExecuteResponse(w, r, "users.html", users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	email := r.PostForm.Get("email")
	admin := r.PostForm.Get("admin")

	user := &types.User{
		Username: username,
		Password: password,
		Email:    email,
		Admin:    (admin == "on"),
	}
	_, err := core.CreateUser(user)
	if err != nil {
		utils.Log(2, err)
	}
	core.OnNewUser(user)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func UsersDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := core.SelectUser(int64(id))

	users, _ := core.SelectAllUsers()
	if len(users) == 1 {
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}
	core.DeleteUser(user)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
