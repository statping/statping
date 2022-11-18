package handlers

import (
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/statping-ng/statping-ng/types/errors"
	"github.com/statping-ng/statping-ng/types/failures"
	"github.com/statping-ng/statping-ng/types/notifications"
	"github.com/statping-ng/statping-ng/types/services"
)

func apiAllNotifiersHandler(r *http.Request) interface{} {
	var notifs []notifications.Notification
	for _, n := range services.AllNotifiers() {
		notif := n.Select()
		no, err := notifications.Find(notif.Method)
		if err != nil {
			log.Error(err)
		}
		notif.UpdateFields(no)
		notifs = append(notifs, *notif)
	}
	sort.Sort(notifications.NotificationOrder(notifs))
	return notifs
}

func apiNotifierGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notifer := services.FindNotifier(vars["notifier"])
	if notifer == nil {
		sendErrorJson(errors.New("could not find notifier"), w, r)
		return
	}
	returnJson(notifer, w, r)
}

func apiNotifierUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notifer, err := notifications.Find(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := DecodeJSON(r, &notifer); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	log.Infof("Updating %s Notifier", notifer.Title)

	if err := notifer.Update(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	notif := services.ReturnNotifier(notifer.Method)
	if err := notif.Valid(notifer.Values()); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if _, err := notif.OnSave(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(vars["notifier"], "update", w, r)
}

type testNotificationReq struct {
	Method       string                     `json:"method"`
	Notification notifications.Notification `json:"notifier"`
}

func testNotificationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n := services.FindNotifier(vars["notifier"])
	if n == nil {
		sendErrorJson(errors.New("unknown notifier"), w, r)
		return
	}

	var req testNotificationReq
	if err := DecodeJSON(r, &req); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	notif := services.ReturnNotifier(n.Method)

	var out string
	var err error
	if req.Method == "success" {
		out, err = notif.OnSuccess(services.Example(true))
	} else {
		out, err = notif.OnFailure(services.Example(false), failures.Example())
	}

	resp := &notifierTestResp{
		Success:  err == nil,
		Response: out,
		Error:    err,
	}
	returnJson(resp, w, r)
}

type notifierTestResp struct {
	Success  bool   `json:"success"`
	Response string `json:"response,omitempty"`
	Error    error  `json:"error,omitempty"`
}
