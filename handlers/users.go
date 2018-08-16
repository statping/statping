// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"strconv"
)

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
		utils.Log(2, "cannot delete the only user in the system")
		http.Redirect(w, r, "/users", http.StatusSeeOther)
		return
	}
	core.DeleteUser(user)
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
