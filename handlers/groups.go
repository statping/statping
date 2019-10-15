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

func groupViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var group *core.Group
	id := vars["id"]
	group = core.SelectGroup(utils.ToInt(id))

	if group == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ExecuteResponse(w, r, "group.gohtml", group, nil)
}

// apiAllGroupHandler will show all the groups
func apiAllGroupHandler(w http.ResponseWriter, r *http.Request) {
	auth, admin := IsUser(r), IsAdmin(r)
	groups := core.SelectGroups(admin, auth)
	returnJson(groups, w, r)
}

// apiGroupHandler will show a single group
func apiGroupHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := core.SelectGroup(utils.ToInt(vars["id"]))
	if group == nil {
		sendErrorJson(errors.New("group not found"), w, r)
		return
	}
	returnJson(group, w, r)
}

// apiGroupUpdateHandler will update a group
func apiGroupUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := core.SelectGroup(utils.ToInt(vars["id"]))
	if group == nil {
		sendErrorJson(errors.New("group not found"), w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&group)
	_, err := group.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(group, "update", w, r)
}

// apiCreateGroupHandler accepts a POST method to create new groups
func apiCreateGroupHandler(w http.ResponseWriter, r *http.Request) {
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

// apiGroupDeleteHandler accepts a DELETE method to delete groups
func apiGroupDeleteHandler(w http.ResponseWriter, r *http.Request) {
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

type groupOrder struct {
	Id    int64 `json:"group"`
	Order int   `json:"order"`
}

func apiGroupReorderHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var newOrder []*groupOrder
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newOrder)
	for _, g := range newOrder {
		group := core.SelectGroup(g.Id)
		group.Order = g.Order
		group.Update()
	}
	returnJson(newOrder, w, r)
}
