package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/http"
)

func apiNotifiersHandler(w http.ResponseWriter, r *http.Request) {
	notifiers := services.AllNotifiers()
	returnJson(notifiers, w, r)
}

func apiNotifierGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notif := services.FindNotifier(vars["notifier"])
	notifer, err := notifications.Find(notif.Method)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notif = notif.UpdateFields(notifer)
	returnJson(notif, w, r)
}

func apiNotifierUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notifer, err := notifications.Find(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	var notif *notifications.Notification
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&notif)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifer = notifer.UpdateFields(notif)
	err = notifer.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	//notifications.OnSave(notifer.Method)
	sendJsonAction(vars["notifier"], "update", w, r)
}

func testNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	form := parseForm(r)
	vars := mux.Vars(r)
	method := vars["method"]
	enabled := form.Get("enable")
	host := form.Get("host")
	port := int(utils.ToInt(form.Get("port")))
	username := form.Get("username")
	password := form.Get("password")
	var1 := form.Get("var1")
	var2 := form.Get("var2")
	apiKey := form.Get("api_key")
	apiSecret := form.Get("api_secret")
	limits := int(utils.ToInt(form.Get("limits")))

	notifier, err := notifications.Find(method)
	if err != nil {
		log.Errorln(fmt.Sprintf("issue saving notifier %v: %v", method, err))
		sendErrorJson(err, w, r)
		return
	}

	n := notifier

	if host != "" {
		n.Host = host
	}
	if port != 0 {
		n.Port = port
	}
	if username != "" {
		n.Username = username
	}
	if password != "" && password != "##########" {
		n.Password = password
	}
	if var1 != "" {
		n.Var1 = var1
	}
	if var2 != "" {
		n.Var2 = var2
	}
	if apiKey != "" {
		n.ApiKey = apiKey
	}
	if apiSecret != "" {
		n.ApiSecret = apiSecret
	}
	if limits != 0 {
		n.Limits = limits
	}
	n.Enabled = null.NewNullBool(enabled == "on")

	//err = notifications.OnTest(notifier)
	//if err == nil {
	//	w.Write([]byte("ok"))
	//} else {
	//	w.Write([]byte(err.Error()))
	//}
}
