package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core/integrations"
	"github.com/hunterlong/statping/types"
	"net/http"
)

func apiAllIntegrationsHandler(w http.ResponseWriter, r *http.Request) {
	integrations := integrations.Integrations
	returnJson(integrations, w, r)
}

func apiIntegrationViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	intgr, err := integrations.Find(vars["name"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(intgr.Get(), w, r)
}

func apiIntegrationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	intgr, err := integrations.Find(vars["name"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	var intJson *types.Integration
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&intJson); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	integration := intgr.Get()
	integration.Enabled = intJson.Enabled
	integration.Fields = intJson.Fields

	if err := integrations.Update(integration); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	list, err := intgr.List()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(list, w, r)
}
