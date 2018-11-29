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
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"time"
)

func renderServiceChartsHandler(w http.ResponseWriter, r *http.Request) {
	services := core.CoreApp.Services
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=60")

	end := time.Now().UTC()
	start := time.Now().Add((-24 * 7) * time.Hour).UTC()
	var srvs []*core.Service
	for _, s := range services {
		srvs = append(srvs, s.(*core.Service))
	}
	out := struct {
		Services []*core.Service
		Start    int64
		End      int64
	}{srvs, start.Unix(), end.Unix()}

	executeJSResponse(w, r, "charts.js", out)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "services.html", core.CoreApp.Services, nil)
}

type serviceOrder struct {
	Id    int64 `json:"service"`
	Order int   `json:"order"`
}

func reorderServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var newOrder []*serviceOrder
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newOrder); err != nil {
		utils.Log(3, fmt.Sprint("error decoding reordering services: %v", err.Error()))
	}
	for _, s := range newOrder {
		service := core.SelectService(s.Id)
		service.Order = s.Order
		if err := service.Update(false); err != nil {
			utils.Log(3, fmt.Sprint("error reordering services: %v", err.Error()))
		}
	}
	w.Write([]byte("ok"))
	w.WriteHeader(http.StatusOK)
}

func servicesViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fields := parseGet(r)
	r.ParseForm()

	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	group := r.Form.Get("group")
	serv := core.SelectService(utils.ToInt(vars["id"]))
	if serv == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	end := time.Now().UTC()
	start := end.Add((-24 * 7) * time.Hour).UTC()

	if startField != 0 {
		start = time.Unix(startField, 0).UTC()
	}
	if endField != 0 {
		end = time.Unix(endField, 0).UTC()
	}
	if group == "" {
		group = "hour"
	}

	data := core.GraphDataRaw(serv, start, end, group, "latency")

	out := struct {
		Service   *core.Service
		Start     string
		End       string
		StartUnix int64
		EndUnix   int64
		Data      string
	}{serv, start.Format(utils.FlatpickrReadable), end.Format(utils.FlatpickrReadable), start.Unix(), end.Unix(), data.ToString()}

	ExecuteResponse(w, r, "service.html", out, nil)
}

func apiServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	servicer := core.SelectServicer(utils.ToInt(vars["id"]))
	if servicer == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	service := servicer.Select()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service)
}

func apiCreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(r) {
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
	if !isAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectServicer(utils.ToInt(vars["id"]))
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

func apiServiceDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}
	fields := parseGet(r)
	grouping := fields.Get("group")
	if grouping == "" {
		grouping = "hour"
	}
	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))

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
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	fields := parseGet(r)
	grouping := fields.Get("group")
	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	obj := core.GraphDataRaw(service, time.Unix(startField, 0), time.Unix(endField, 0), grouping, "ping_time")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func apiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
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
	if !isAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	services := core.Services()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
}

func servicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	service.DeleteFailures()
	ExecuteResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}
