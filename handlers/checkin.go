package handlers

import (
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net"
	"net/http"
)

func findCheckin(r *http.Request) (*checkins.Checkin, string, error) {
	vars := mux.Vars(r)
	id := vars["api"]
	if id == "" {
		return nil, "", errors.IDMissing
	}
	checkin, err := checkins.FindByAPI(id)
	if err != nil {
		return nil, id, errors.Missing(checkins.Checkin{}, id)
	}
	return checkin, id, nil
}

func apiAllCheckinsHandler(w http.ResponseWriter, r *http.Request) {
	returnJson(checkins.All(), w, r)
}

func apiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	checkin, _, err := findCheckin(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(checkin, w, r)
}

func checkinCreateHandler(w http.ResponseWriter, r *http.Request) {
	var checkin *checkins.Checkin
	if err := DecodeJSON(r, &checkin); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	service, err := services.Find(checkin.ServiceId)
	if err != nil {
		sendErrorJson(err, w, r)
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
	checkin, _, err := findCheckin(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	log.Infof("Checking %s was requested", checkin.Name)

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	if last := checkin.LastHit(); last == nil {
		checkin.Start()
	}

	hit := &checkins.CheckinHit{
		Checkin:   checkin.Id,
		From:      ip,
		CreatedAt: utils.Now(),
	}

	if err := hit.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	checkin.Failing = false
	checkin.LastHitTime = utils.Now()

	sendJsonAction(hit.Id, "update", w, r)
}

func checkinDeleteHandler(w http.ResponseWriter, r *http.Request) {
	checkin, _, err := findCheckin(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := checkin.Delete(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(checkin, "delete", w, r)
}
