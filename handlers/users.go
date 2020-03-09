// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
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
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"net/http"
)

func getUser(r *http.Request) (*users.User, int64, error) {
	vars := mux.Vars(r)
	num := utils.ToInt(vars["id"])
	user, err := users.Find(num)
	if err != nil {
		return nil, num, err
	}
	return user, num, nil
}

func apiUserHandler(w http.ResponseWriter, r *http.Request) {
	user, _, err := getUser(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	user.Password = ""
	returnJson(user, w, r)
}

func apiUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	user, id, err := getUser(r)
	if err != nil {
		sendErrorJson(fmt.Errorf("user #%d was not found", id), w, r)
		return
	}

	err = DecodeJSON(r, &user)
	if err != nil {
		sendErrorJson(fmt.Errorf("user #%d was not found", id), w, r)
		return
	}

	if user.Password != "" {
		user.Password = utils.HashPassword(user.Password)
	}

	err = user.Update()
	if err != nil {
		sendErrorJson(fmt.Errorf("issue updating user #%d: %s", user.Id, err), w, r)
		return
	}
	sendJsonAction(user, "update", w, r)
}

func apiUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	allUsers := users.All()
	if len(allUsers) == 1 {
		sendErrorJson(errors.New("cannot delete the last user"), w, r)
		return
	}
	user, _, err := getUser(r)
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
	allUsers := users.All()
	returnJson(allUsers, w, r)
}

func apiCreateUsersHandler(w http.ResponseWriter, r *http.Request) {
	var user *users.User
	err := DecodeJSON(r, &user)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	err = user.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(user, "create", w, r)
}
