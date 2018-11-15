// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"os"
	"time"
)

type apiResponse struct {
	Status string `json:"status"`
	Object string `json:"type,omitempty"`
	Id     int64  `json:"id,omitempty"`
	Method string `json:"method,omitempty"`
	Error  string `json:"error,omitempty"`
}

func apiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(core.CoreApp)
}

func apiRenewHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	var err error
	core.CoreApp.ApiKey = utils.NewSHA1Hash(40)
	core.CoreApp.ApiSecret = utils.NewSHA1Hash(40)
	core.CoreApp, err = core.UpdateCore(core.CoreApp)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func apiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	checkin := core.SelectCheckin(vars["api"])
	//checkin.Receivehit()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkin)
}

func apiServiceDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}
	fields := parseGet(r)
	grouping := fields.Get("group")
	startField := utils.StringInt(fields.Get("start"))
	endField := utils.StringInt(fields.Get("end"))

	if startField == 0 || endField == 0 {
		startField = 0
		endField = 99999999999
	}

	obj := core.GraphDataRaw(service, time.Unix(startField, 0).UTC(), time.Unix(endField, 0).UTC(), grouping, "latency")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func apiServicePingDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	fields := parseGet(r)
	grouping := fields.Get("group")
	startField := utils.StringInt(fields.Get("start"))
	endField := utils.StringInt(fields.Get("end"))
	obj := core.GraphDataRaw(service, time.Unix(startField, 0), time.Unix(endField, 0), grouping, "ping_time")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func apiServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	servicer := core.SelectServicer(utils.StringInt(vars["id"]))
	if servicer == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	service := servicer.Select()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func apiCreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	var service *types.Service
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	newService := core.ReturnService(service)
	_, err = newService.Create(true)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func apiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	var updatedService *types.Service
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&updatedService)
	updatedService.Id = service.Id
	service = core.ReturnService(updatedService)
	err := service.Update(true)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	go service.Check(true)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func apiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	err := service.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	output := apiResponse{
		Object: "service",
		Method: "delete",
		Id:     service.Id,
		Status: "success",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func apiAllServicesHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	services := core.Services()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func apiNotifierGetHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	_, notifierObj, err := notifier.SelectNotifier(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifierObj)
}

func apiNotifierUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	var notification *notifier.Notification
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&notification)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifer, not, err := notifier.SelectNotifier(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifer.Var1 = notification.Var1
	notifer.Var2 = notification.Var2
	notifer.Host = notification.Host
	notifer.Port = notification.Port
	notifer.Password = notification.Password
	notifer.Username = notification.Username
	notifer.ApiKey = notification.ApiKey
	notifer.ApiSecret = notification.ApiSecret
	notifer.Enabled = types.NewNullBool(notification.Enabled.Bool)

	_, err = notifier.Update(not, notifer)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifier.OnSave(notifer.Method)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifer)
}

func apiAllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	messages, err := core.SelectMessages()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func apiMessageCreateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	var message *types.Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&message)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	msg := core.ReturnMessage(message)
	_, err = msg.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func apiMessageGetHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	message, err := core.SelectMessage(utils.StringInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func apiMessageDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	message, err := core.SelectMessage(utils.StringInt(vars["id"]))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = message.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	output := apiResponse{
		Object: "message",
		Method: "delete",
		Id:     message.Id,
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func apiMessageUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	message, err := core.SelectMessage(utils.StringInt(vars["id"]))
	if err != nil {
		sendErrorJson(fmt.Errorf("message #%v was not found", vars["id"]), w, r)
		return
	}
	var messageBody *types.Message
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&messageBody)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	messageBody.Id = message.Id
	message = core.ReturnMessage(messageBody)
	_, err = message.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	output := apiResponse{
		Object: "message",
		Method: "update",
		Id:     message.Id,
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func apiNotifiersHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	var notifiers []*notifier.Notification
	for _, n := range core.CoreApp.Notifications {
		notif := n.(notifier.Notifier)
		notifiers = append(notifiers, notif.Select())
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifiers)
}

func sendErrorJson(err error, w http.ResponseWriter, r *http.Request) {
	output := apiResponse{
		Status: "error",
		Error:  err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func sendUnauthorizedJson(w http.ResponseWriter, r *http.Request) {
	output := apiResponse{
		Status: "error",
		Error:  errors.New("not authorized").Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(output)
}

func isAPIAuthorized(r *http.Request) bool {
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	if IsAuthenticated(r) {
		return true
	}
	if isAuthorized(r) {
		return true
	}
	return false
}
