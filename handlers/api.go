package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(core.CoreApp)
}

func ApiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	checkin := core.FindCheckin(vars["api"])
	//checkin.Receivehit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkin)
}

func ApiServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	json.NewEncoder(w).Encode(service)
}

func ApiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	var s *types.Service
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&s)
	service := serv.ToService()
	core.UpdateService(service)
	json.NewEncoder(w).Encode(s)
}

func ApiAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	services, _ := core.SelectAllServices()
	json.NewEncoder(w).Encode(services)
}

func ApiUserHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(r)
	user, _ := core.SelectUser(utils.StringInt(vars["id"]))
	json.NewEncoder(w).Encode(user)
}

func ApiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	users, _ := core.SelectAllUsers()
	json.NewEncoder(w).Encode(users)
}

func isAPIAuthorized(r *http.Request) bool {
	if IsAuthenticated(r) {
		return true
	}
	if isAuthorized(r) {
		return true
	}
	return false
}
