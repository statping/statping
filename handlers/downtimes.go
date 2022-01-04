package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/downtimes"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/http"
	"time"
)

func findDowntime(r *http.Request) (*downtimes.Downtime, error) {
	vars := mux.Vars(r)
	id := utils.ToInt(vars["id"])
	downtime, err := downtimes.Find(id)
	if err != nil {
		return nil, err
	}

	return downtime, nil
}

func apiAllDowntimesForServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serviceId := utils.ToInt(vars["service_id"])

	ninetyDaysAgo := time.Now().Add(time.Duration(-90*24) * time.Hour)

	downtime, err := downtimes.FindByService(serviceId, ninetyDaysAgo, time.Now())
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(downtime, "fetch", w, r)
}

func apiCreateDowntimeHandler(w http.ResponseWriter, r *http.Request) {
	var downtime *downtimes.Downtime
	if err := DecodeJSON(r, &downtime); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	s, err := services.FindFirstFromDB(downtime.ServiceId)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	downtime.Type = "manual"
	downtime.Id = zeroInt64

	if err := downtime.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if downtime.End == nil {
		updateFields := map[string]interface{}{
			"online":           false,
			"current_downtime": downtime.Id,
			"manual_downtime":  true,
		}

		if err := s.UpdateSpecificFields(updateFields); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	sendJsonAction(downtime, "create", w, r)
}

func apiDowntimeHandler(w http.ResponseWriter, r *http.Request) {
	downtime, err := findDowntime(r)
	if downtime == nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(downtime, "fetch", w, r)
}

func apiPatchDowntimeHandler(w http.ResponseWriter, r *http.Request) {
	downtime, err := findDowntime(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	var req downtimes.Downtime
	if err := DecodeJSON(r, &req); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	s, err := services.FindFirstFromDB(downtime.ServiceId)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if downtime.End != nil && req.End == nil {
		fmt.Errorf("Cannot reopen a downtime!")
	}

	if downtime.End == nil && req.End != nil {
		updateFields := map[string]interface{}{
			"online":           true,
			"failure_counter":  0,
			"current_downtime": 0,
			"manual_downtime":  false,
		}

		if err := s.UpdateSpecificFields(updateFields); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	downtime.Start = req.Start
	downtime.End = req.End
	downtime.Type = "manual"
	downtime.Failures = req.Failures
	downtime.SubStatus = req.SubStatus

	if err := downtime.Update(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(downtime, "update", w, r)
}

func apiDeleteDowntimeHandler(w http.ResponseWriter, r *http.Request) {
	downtime, err := findDowntime(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if downtime.End == nil {
		s2, err := services.FindFirstFromDB(downtime.ServiceId)
		if err != nil {
			fmt.Errorf("Error updating service")
		}
		s2.LastProcessingTime = zeroTime
		s2.Online = true
		s2.FailureCounter = 0
		s2.CurrentDowntime = 0
		s2.ManualDowntime = false

		if err := s2.Update(); err != nil {
			sendErrorJson(err, w, r)
			return
		}
	}

	err = downtime.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(downtime, "delete", w, r)
}
