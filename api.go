package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(core)
}

func ApiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin := FindCheckin(vars["api"])
	checkin.Receivehit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkin)
}

func ApiServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, _ := SelectService(StringInt(vars["id"]))
	json.NewEncoder(w).Encode(service)
}

func ApiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, _ := SelectService(StringInt(vars["id"]))

	var s Service
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&s)

	json.NewEncoder(w).Encode(service)
}

func ApiAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	services, _ := SelectAllServices()
	json.NewEncoder(w).Encode(services)
}

func ApiUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, _ := SelectUser(StringInt(vars["id"]))
	json.NewEncoder(w).Encode(user)
}

func ApiAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := SelectAllUsers()
	json.NewEncoder(w).Encode(users)
}
