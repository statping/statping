package handlers

import (
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core/integrations"
	"net/http"
)

func apiAllIntegrationsHandler(w http.ResponseWriter, r *http.Request) {
	integrations := integrations.Integrations
	returnJson(integrations, w, r)
}

func apiIntegrationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	intg := vars["name"]
	r.ParseForm()
	for k, v := range r.PostForm {
		log.Info(k, v)
	}

	integration, err := integrations.Find(intg)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	list, err := integration.List()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(list, w, r)
}
