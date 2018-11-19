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
	Status string      `json:"status"`
	Object string      `json:"type,omitempty"`
	Method string      `json:"method,omitempty"`
	Error  string      `json:"error,omitempty"`
	Id     int64       `json:"id,omitempty"`
	Output interface{} `json:"output,omitempty"`
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
	if grouping == "" {
		grouping = "hour"
	}
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
	sendJsonAction(newService, "create", w, r)
}

func apiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !isAPIAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectServicer(utils.StringInt(vars["id"]))
	if service.Select() == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&service)
	err := service.Update(true)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	go service.Check(true)

	sendJsonAction(service, "update", w, r)
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
	sendJsonAction(service, "delete", w, r)
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
	notifer, not, err := notifier.SelectNotifier(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&notifer)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	_, err = notifier.Update(not, notifer)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifier.OnSave(notifer.Method)
	sendJsonAction(notifer, "update", w, r)
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

func sendJsonAction(obj interface{}, method string, w http.ResponseWriter, r *http.Request) {
	var objName string
	var objId int64
	switch v := obj.(type) {
	case types.ServiceInterface:
		objName = "service"
		objId = v.Select().Id
	case *notifier.Notification:
		objName = "notifier"
		objId = v.Id
	case *core.Core, *types.Core:
		objName = "core"
	case *types.User:
		objName = "user"
		objId = v.Id
	case *core.User:
		objName = "user"
		objId = v.Id
	case *types.Message:
		objName = "message"
		objId = v.Id
	case *core.Message:
		objName = "message"
		objId = v.Id
	case *types.Checkin:
		objName = "checkin"
		objId = v.Id
	default:
		objName = "missing"
	}

	output := apiResponse{
		Object: objName,
		Method: method,
		Id:     objId,
		Status: "success",
		Output: obj,
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
	utils.Http(r)
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
