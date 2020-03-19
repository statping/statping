package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/statping/statping/utils"
	"net/http"
)

func DecodeJSON(r *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&obj)
	if err != nil {
		return err
	}
	return nil
}

func GetID(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		return 0, errors.New("no id specified in request")
	}
	return utils.ToInt(vars["id"]), nil
}
