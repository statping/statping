// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"strconv"
	"time"
)

func renderServiceChartsHandler(w http.ResponseWriter, r *http.Request) {
	services := core.CoreApp.Services
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=60")

	end := utils.Now().UTC()
	start := utils.Now().Add((-24 * 7) * time.Hour).UTC()
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
	data := map[string]interface{}{
		"Services": core.CoreApp.Services,
	}
	ExecuteResponse(w, r, "services.gohtml", data, nil)
}

type serviceOrder struct {
	Id    int64 `json:"service"`
	Order int   `json:"order"`
}

func reorderServiceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var newOrder []*serviceOrder
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newOrder)
	for _, s := range newOrder {
		service := core.SelectService(s.Id)
		service.Order = s.Order
		service.Update(false)
	}
	returnJson(newOrder, w, r)
}

func servicesViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fields := parseGet(r)
	r.ParseForm()

	var serv *core.Service
	id := vars["id"]
	if _, err := strconv.Atoi(id); err == nil {
		serv = core.SelectService(utils.ToInt(id))
	} else {
		serv = core.SelectServiceLink(id)
	}
	if serv == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	group := r.Form.Get("group")

	end := utils.Now().UTC()
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

	ExecuteResponse(w, r, "service.gohtml", out, nil)
}

func apiServiceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	servicer := core.SelectService(utils.ToInt(vars["id"]))
	if servicer == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	returnJson(servicer, w, r)
}

func apiCreateServiceHandler(w http.ResponseWriter, r *http.Request) {
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
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
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

func apiServiceRunningHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	if service.IsRunning() {
		service.Close()
	} else {
		service.Start()
	}
	sendJsonAction(service, "running", w, r)
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

	start := time.Unix(startField, 0)
	end := time.Unix(endField, 0)

	obj := core.GraphDataRaw(service, start, end, grouping, "latency")
	returnJson(obj, w, r)
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

	start := time.Unix(startField, 0)
	end := time.Unix(endField, 0)

	obj := core.GraphDataRaw(service, start, end, grouping, "ping_time")
	returnJson(obj, w, r)
}

type dataXy struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type dataXyMonth struct {
	Date string    `json:"date"`
	Data []*dataXy `json:"data"`
}

func apiServiceHeatmapHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	var monthOutput []*dataXyMonth

	start := service.CreatedAt
	//now := utils.Now()

	sY, sM, _ := start.Date()

	var date time.Time

	month := int(sM)
	maxMonth := 12

	for year := int(sY); year <= utils.Now().Year(); year++ {

		if year == utils.Now().Year() {
			maxMonth = int(utils.Now().Month())
		}

		for m := month; m <= maxMonth; m++ {

			var output []*dataXy

			for day := 1; day <= 31; day++ {
				date = time.Date(year, time.Month(m), day, 0, 0, 0, 0, time.UTC)
				failures, _ := service.TotalFailuresOnDate(date)
				output = append(output, &dataXy{day, int(failures)})
			}

			thisDate := fmt.Sprintf("%v-%v-01 00:00:00", year, m)
			monthOutput = append(monthOutput, &dataXyMonth{thisDate, output})
		}

		month = 1

	}
	returnJson(monthOutput, w, r)
}

func apiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
	services := core.Services()
	returnJson(services, w, r)
}

func servicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	service.DeleteFailures()
	sendJsonAction(service, "delete_failures", w, r)
}

func apiServiceFailuresHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	servicer := core.SelectService(utils.ToInt(vars["id"]))
	if servicer == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	returnJson(servicer.AllFailures(), w, r)
}

func apiServiceHitsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	servicer := core.SelectService(utils.ToInt(vars["id"]))
	if servicer == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}

	hits, err := servicer.Hits()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	returnJson(hits, w, r)
}

func createServiceHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteResponse(w, r, "service_create.gohtml", core.CoreApp, nil)
}
