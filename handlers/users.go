// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"strconv"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := core.SelectAllUsers()
	ExecuteResponse(w, r, "users.gohtml", users, nil)
}

func usersEditHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, _ := core.SelectUser(int64(id))
	ExecuteResponse(w, r, "user.gohtml", user, nil)
}

func apiUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := core.SelectUser(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	user.Password = ""
	returnJson(user, w, r)
}

func apiUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := core.SelectUser(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(fmt.Errorf("user #%v was not found", vars["id"]), w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	if user.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}
	err = user.Update()
	if err != nil {
		sendErrorJson(fmt.Errorf("issue updating user #%v: %v", user.Id, err), w, r)
		return
	}
	sendJsonAction(user, "update", w, r)
}

func apiUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	users := core.CountUsers()
	if users == 1 {
		sendErrorJson(errors.New("cannot delete the last user"), w, r)
		return
	}
	user, err := core.SelectUser(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = user.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(user, "delete", w, r)
}

func apiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := core.SelectAllUsers()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(users, w, r)
}

func apiCreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user *types.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	newUser := core.ReturnUser(user)
	_, err = newUser.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(newUser, "create", w, r)
}
