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
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

func apiAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !IsReadAuthenticated(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	groups := core.SelectGroups()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func apiGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !IsReadAuthenticated(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	group := core.SelectGroup(utils.ToInt(vars["id"]))
	if group == nil {
		sendErrorJson(errors.New("group not found"), w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}

func apiCreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if !IsFullAuthenticated(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	var group *core.Group
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&group)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	_, err = group.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(group, "create", w, r)
}

func apiGroupDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !IsFullAuthenticated(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	group := core.SelectGroup(utils.ToInt(vars["id"]))
	if group == nil {
		sendErrorJson(errors.New("group not found"), w, r)
		return
	}
	err := group.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(group, "delete", w, r)
}
