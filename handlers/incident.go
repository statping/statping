package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/utils"
	"net/http"
)

func apiAllIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	inc := incidents.All()
	returnJson(inc, w, r)
}

func apiCreateIncidentHandler(w http.ResponseWriter, r *http.Request) {
	var incident *incidents.Incident
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&incident)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = incident.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(incident, "create", w, r)
}

func apiIncidentUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incident, err := incidents.Find(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&incident)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	updates := incident.Updates()
	sendJsonAction(updates, "update", w, r)
}

func apiDeleteIncidentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incident, err := incidents.Find(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = incident.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(incident, "delete", w, r)
}
