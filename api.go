package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"math/rand"
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
	service := SelectService(StringInt(vars["id"]))
	json.NewEncoder(w).Encode(service)
}

func ApiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := SelectService(StringInt(vars["id"]))
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

func NewSHA1Hash(n ...int) string {
	noRandomCharacters := 32
	if len(n) > 0 {
		noRandomCharacters = n[0]
	}
	randString := RandomString(noRandomCharacters)
	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

var characterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString generates a random string of n length
func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = characterRunes[rand.Intn(len(characterRunes))]
	}
	return string(b)
}
