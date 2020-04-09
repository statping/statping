package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net"
	"net/http"
)

func apiAllCheckinsHandler(w http.ResponseWriter, r *http.Request) {
	chks := checkins.All()
	returnJson(chks, w, r)
}

func apiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin, err := checkins.FindByAPI(vars["api"])
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	returnJson(checkin, w, r)
}

func checkinCreateHandler(w http.ResponseWriter, r *http.Request) {
	var checkin *checkins.Checkin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&checkin)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	service, err := services.Find(checkin.ServiceId)
	if err != nil {
		sendErrorJson(fmt.Errorf("missing service_id field"), w, r)
		return
	}
	checkin.ServiceId = service.Id
	if err := checkin.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(checkin, "create", w, r)
}

func checkinHitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin, err := checkins.FindByAPI(vars["api"])
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %s was not found", vars["api"]), w, r)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	hit := &checkins.CheckinHit{
		Checkin:   checkin.Id,
		From:      ip,
		CreatedAt: utils.Now(),
	}
	log.Infof("Checking %s was requested", checkin.Name)

	err = hit.Create()
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	checkin.Failing = false
	checkin.LastHitTime = utils.Now()
	sendJsonAction(hit.Id, "update", w, r)
}

func checkinDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin, err := checkins.FindByAPI(vars["api"])
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}

	if err := checkin.Delete(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(checkin, "delete", w, r)
}
