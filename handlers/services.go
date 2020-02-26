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
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

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
		service := core.SelectService(s.Id).Model()
		service.Order = s.Order
		database.Update(service)
	}
	returnJson(newOrder, w, r)
}

func apiServiceHandler(r *http.Request) interface{} {
	vars := mux.Vars(r)
	servicer := core.SelectService(utils.ToInt(vars["id"]))
	if servicer == nil {
		return errors.New("service not found")
	}
	return servicer.Service
}

func apiCreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service *types.Service
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	_, err = database.Create(service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(service, "create", w, r)
}

func apiTestServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service *types.Service
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	_, err = database.Create(service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(service, "create", w, r)
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
	err := database.Update(service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	go core.CheckService(service, true)
	sendJsonAction(service, "update", w, r)
}

func apiServiceRunningHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	srv := core.SelectService(utils.ToInt(vars["id"]))
	service := srv.Model()
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
	service, err := database.Service(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	groupQuery := database.ParseQueries(r, service.Hits())

	objs, err := groupQuery.GraphData(database.ByAverage("latency"))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(objs, w, r)
}

func apiServiceFailureDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, err := database.Service(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	groupQuery := database.ParseQueries(r, service.Failures())

	objs, err := groupQuery.GraphData(database.ByCount)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	returnJson(objs, w, r)
}

func apiServicePingDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, err := database.Service(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	groupQuery := database.ParseQueries(r, service.Hits())

	objs, err := groupQuery.GraphData(database.ByAverage("ping_time"))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	returnJson(objs, w, r)
}

type dataXy struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type dataXyMonth struct {
	Date string    `json:"date"`
	Data []*dataXy `json:"data"`
}

func apiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.ToInt(vars["id"]))
	if service == nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	err := database.Delete(service)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(service, "delete", w, r)
}

func apiAllServicesHandler(r *http.Request) interface{} {
	services := core.Services()
	return joinServices(services)
}

func joinServices(srvs []*core.Service) []*types.Service {
	var services []*types.Service
	for _, v := range srvs {
		v.UpdateStats()
		services = append(services, v.Service)
	}
	return services
}

func servicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	srv, err := database.Service(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(errors.New("service not found"), w, r)
		return
	}
	err = srv.Failures().DeleteAll()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(srv.Model(), "delete_failures", w, r)
}

func apiServiceFailuresHandler(r *http.Request) interface{} {
	vars := mux.Vars(r)

	service, err := database.Service(utils.ToInt(vars["id"]))
	if err != nil {
		return errors.New("service not found")
	}

	var fails []*types.Failure
	database.ParseQueries(r, service.Failures()).Find(&fails)
	return fails
}

func apiServiceHitsHandler(r *http.Request) interface{} {
	vars := mux.Vars(r)
	service, err := database.Service(utils.ToInt(vars["id"]))
	if err != nil {
		return errors.New("service not found")
	}

	var hits []types.Hit
	database.ParseQueries(r, service.Hits()).Find(&hits)
	return hits
}

func createServiceHandler(w http.ResponseWriter, r *http.Request) {
	ExecuteResponse(w, r, "service_create.gohtml", core.CoreApp, nil)
}
