package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/types/integrations"
	"net/http"
)

func findIntegration(r *http.Request) (*integrations.Integration, string, error) {
	vars := mux.Vars(r)
	name := vars["name"]
	intgr, err := integrations.Find(name)
	if err != nil {
		return nil, "", err
	}
	return intgr, name, nil
}

func apiAllIntegrationsHandler(w http.ResponseWriter, r *http.Request) {
	inte := integrations.All()
	returnJson(inte, w, r)
}

func apiIntegrationViewHandler(w http.ResponseWriter, r *http.Request) {
	intgr, _, err := findIntegration(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(intgr, w, r)
}

func apiIntegrationHandler(w http.ResponseWriter, r *http.Request) {
	intgr, _, err := findIntegration(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&intgr); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := intgr.Update(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	returnJson(intgr, w, r)
}
