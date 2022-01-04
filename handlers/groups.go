package handlers

import (
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/utils"
	"net/http"
)

func findGroup(r *http.Request) (*groups.Group, error) {
	vars := mux.Vars(r)
	if utils.NotNumber(vars["id"]) {
		return nil, errors.NotNumber
	}
	id := utils.ToInt(vars["id"])
	if id == 0 {
		return nil, errors.IDMissing
	}
	g, err := groups.Find(id)
	if err != nil {
		return nil, err
	}
	if !g.Public.Bool && !IsReadAuthenticated(r) {
		return nil, errors.NotAuthenticated
	}
	return g, nil
}

// apiAllGroupHandler will show all the groups
func apiAllGroupHandler(r *http.Request) interface{} {
	auth, admin := IsUser(r), IsAdmin(r)
	return groups.SelectGroups(admin, auth)
}

// apiGroupHandler will show a single group
func apiGroupHandler(w http.ResponseWriter, r *http.Request) {
	group, err := findGroup(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(group, w, r)
}

// apiGroupUpdateHandler will update a group
func apiGroupUpdateHandler(w http.ResponseWriter, r *http.Request) {
	group, err := findGroup(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := DecodeJSON(r, &group); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := group.Update(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(group, "update", w, r)
}

// apiCreateGroupHandler accepts a POST method to create new groups
func apiCreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group *groups.Group
	if err := DecodeJSON(r, &group); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := group.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(group, "create", w, r)
}

// apiGroupDeleteHandler accepts a DELETE method to delete groups
func apiGroupDeleteHandler(w http.ResponseWriter, r *http.Request) {
	group, err := findGroup(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := group.Delete(); err != nil {
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

	if err := DecodeJSON(r, &newOrder); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	for _, g := range newOrder {
		group, err := groups.Find(g.Id)
		if err != nil {
			sendErrorJson(err, w, r)
			return
		}
		group.Order = g.Order
		if err := group.Update(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}
	returnJson(newOrder, w, r)
}
