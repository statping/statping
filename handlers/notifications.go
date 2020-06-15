package handlers

import (
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/services"
	"net/http"
	"sort"
)

func apiNotifiersHandler(w http.ResponseWriter, r *http.Request) {
	var notifs []notifications.Notification
	notifiers := services.AllNotifiers()
	for _, n := range notifiers {
		notif := n.Select()
		notifer, _ := notifications.Find(notif.Method)
		notif.UpdateFields(notifer)
		notifs = append(notifs, *notif)
	}
	sort.Sort(notifications.NotificationOrder(notifs))
	returnJson(notifs, w, r)
}

func apiNotifierGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notif := services.FindNotifier(vars["notifier"])
	notifer, err := notifications.Find(notif.Method)
	if err != nil {
		sendErrorJson(err, w, r)
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

	err = notifer.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	//notifications.OnSave(notifer.Method)
	sendJsonAction(vars["notifier"], "update", w, r)
}

func testNotificationHandler(w http.ResponseWriter, r *http.Request) {
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

	notif := services.ReturnNotifier(notifer.Method)
	out, err := notif.OnTest()

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
