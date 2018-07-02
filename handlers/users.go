package handlers

import (
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
	res.One(&user)
	return user
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	users, _ := core.SelectAllUsers()
	ExecuteResponse(w, r, "users.html", users)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	email := r.PostForm.Get("email")
	admin := r.PostForm.Get("admin")

	user := &core.User{
		Username: username,
		Password: password,
		Email:    email,
		Admin:    (admin == "on"),
	}
	_, err := user.Create()
	if err != nil {
		utils.Log(2, err)
	}
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func UsersDeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
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
	user.Delete()
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
