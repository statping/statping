package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(core.CoreApp)
}

func ApiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin := core.FindCheckin(vars["api"])
	checkin.Receivehit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkin)
}

func ApiServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	json.NewEncoder(w).Encode(service)
}

func ApiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	var s core.Service
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&s)
	json.NewEncoder(w).Encode(service)
}

func ApiAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, _ := core.SelectAllServices()
	json.NewEncoder(w).Encode(services)
}

func ApiUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, _ := core.SelectUser(utils.StringInt(vars["id"]))
	json.NewEncoder(w).Encode(user)
}

func ApiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := core.SelectAllUsers()
	json.NewEncoder(w).Encode(users)
}
