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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"strconv"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	users, _ := core.SelectAllUsers()
	executeResponse(w, r, "users.html", users, nil)
}

func usersEditHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := core.SelectUser(int64(id))
	executeResponse(w, r, "user.html", user, nil)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	vars := mux.Vars(r)
	id := utils.StringInt(vars["id"])
	user, err := core.SelectUser(id)
	if err != nil {
		utils.Log(3, fmt.Sprintf("user error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Username = r.PostForm.Get("username")
	user.Email = r.PostForm.Get("email")
	isAdmin := r.PostForm.Get("admin") == "on"
	user.Admin = types.NewNullBool(isAdmin)
	password := r.PostForm.Get("password")
	if password != "##########" {
		user.Password = utils.HashPassword(password)
	}
	err = user.Update()
	if err != nil {
		utils.Log(3, err)
	}

	fmt.Println(user.Id)
	fmt.Println(user.Password)
	fmt.Println(user.Admin.Bool)

	users, _ := core.SelectAllUsers()
	executeResponse(w, r, "users.html", users, "/users")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	email := r.PostForm.Get("email")
	admin := r.PostForm.Get("admin")

	user := core.ReturnUser(&types.User{
		Username: username,
		Password: password,
		Email:    email,
		Admin:    types.NewNullBool(admin == "on"),
	})
	_, err := user.Create()
	if err != nil {
		utils.Log(3, err)
	}
	//notifiers.OnNewUser(user)
	executeResponse(w, r, "users.html", user, "/users")
}

func usersDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
	user.Delete()
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
