package handlers

import (
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/utils"
	"net/http"
)

func findIncident(r *http.Request) (*incidents.Incident, int64, error) {
	vars := mux.Vars(r)
	if utils.NotNumber(vars["id"]) {
		return nil, 0, errors.NotNumber
	}
	id := utils.ToInt(vars["id"])
	if id == 0 {
		return nil, id, errors.IDMissing
	}
	checkin, err := incidents.Find(id)
	if err != nil {
		return nil, id, errors.Missing(&incidents.Incident{}, id)
	}
	return checkin, id, nil
}

func apiServiceIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(service.Incidents, w, r)
}

func apiIncidentUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	incid, _, err := findIncident(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(incid.Updates, w, r)
}

func apiCreateIncidentUpdateHandler(w http.ResponseWriter, r *http.Request) {
	incid, _, err := findIncident(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	var update *incidents.IncidentUpdate
	if err := DecodeJSON(r, &update); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	update.IncidentId = incid.Id

	if err := update.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(update, "create", w, r)
}

func apiCreateIncidentHandler(w http.ResponseWriter, r *http.Request) {
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	var incident *incidents.Incident
	if err := DecodeJSON(r, &incident); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	incident.ServiceId = service.Id
	if err := incident.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(incident, "create", w, r)
}

func apiIncidentUpdateHandler(w http.ResponseWriter, r *http.Request) {
	incident, _, err := findIncident(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := DecodeJSON(r, &incident); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(incident.Updates, "update", w, r)
}

func apiDeleteIncidentHandler(w http.ResponseWriter, r *http.Request) {
	incident, _, err := findIncident(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := incident.Delete(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(incident, "delete", w, r)
}

func apiDeleteIncidentUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	update, err := incidents.FindUpdate(utils.ToInt(vars["uid"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := update.Delete(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(update, "delete", w, r)
}
