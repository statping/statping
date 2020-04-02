package handlers

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/failures"
	"github.com/statping/statping/types/hits"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net/http"
)

type serviceOrder struct {
	Id    int64 `json:"service"`
	Order int   `json:"order"`
}

func serviceByID(r *http.Request) (*services.Service, error) {
	vars := mux.Vars(r)
	id := utils.ToInt(vars["id"])
	servicer, err := services.Find(id)
	if err != nil {
		return nil, errors.Errorf("service %d not found", id)
	}
	return servicer, nil
}

func reorderServiceHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var newOrder []*serviceOrder
	if err := DecodeJSON(r, &newOrder); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	for _, s := range newOrder {
		service, err := services.Find(s.Id)
		if err != nil {
			sendErrorJson(errors.Errorf("service %d not found", s.Id), w, r)
			return
		}
		service.Order = s.Order
		service.Update()
	}
	returnJson(newOrder, w, r)
}

func apiServiceHandler(r *http.Request) interface{} {
	srv, err := serviceByID(r)
	if err != nil {
		return err
	}
	user := IsUser(r)
	if !srv.Public.Bool && !user {
		return errors.New("not authenticated")
	}
	srv = srv.UpdateStats()
	return *srv
}

func apiCreateServiceHandler(w http.ResponseWriter, r *http.Request) {
	var service *services.Service
	if err := DecodeJSON(r, &service); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := service.Create(); err != nil {
		sendErrorJson(err, w, r)
		return
	}
	go services.ServiceCheckQueue(service, true)

	sendJsonAction(service, "create", w, r)
}

func apiServiceUpdateHandler(w http.ResponseWriter, r *http.Request) {
	service, err := serviceByID(r)
	if err != nil {
		sendErrorJson(err, w, r, http.StatusNotFound)
		return
	}
	if err := DecodeJSON(r, &service); err != nil {
		sendErrorJson(err, w, r, http.StatusBadRequest)
		return
	}

	err = service.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	go service.CheckService(true)
	sendJsonAction(service, "update", w, r)
}

func apiServiceRunningHandler(w http.ResponseWriter, r *http.Request) {
	service, err := serviceByID(r)
	if err != nil {
		sendErrorJson(err, w, r)
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
	service, err := services.Find(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	groupQuery, err := database.ParseQueries(r, service.AllHits())
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	objs, err := groupQuery.GraphData(database.ByAverage("latency", 1000))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(objs, w, r)
}

func apiServiceFailureDataHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service, err := services.Find(utils.ToInt(vars["id"]))
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	groupQuery, err := database.ParseQueries(r, service.AllFailures())
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	objs, err := groupQuery.GraphData(database.ByCount)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	returnJson(objs, w, r)
}

func apiServicePingDataHandler(w http.ResponseWriter, r *http.Request) {
	service, err := serviceByID(r)
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	groupQuery, err := database.ParseQueries(r, service.AllHits())
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	objs, err := groupQuery.GraphData(database.ByAverage("ping_time", 1000))
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	returnJson(objs, w, r)
}

func apiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	service, err := serviceByID(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = service.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(service, "delete", w, r)
}

func apiAllServicesHandler(r *http.Request) interface{} {
	user := IsUser(r)
	var srvs []services.Service
	for _, v := range services.AllInOrder() {
		if !v.Public.Bool && !user {
			continue
		}
		srvs = append(srvs, v)
	}
	return srvs
}

func servicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	service, err := serviceByID(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = service.DeleteFailures()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(service, "delete_failures", w, r)
}

func apiServiceFailuresHandler(r *http.Request) interface{} {
	vars := mux.Vars(r)
	service, err := services.Find(utils.ToInt(vars["id"]))
	if err != nil {
		return errors.New("service not found")
	}

	var fails []*failures.Failure
	query, err := database.ParseQueries(r, service.AllFailures())
	if err != nil {
		return err
	}
	query.Find(&fails)
	return fails
}

func apiServiceHitsHandler(r *http.Request) interface{} {
	vars := mux.Vars(r)
	service, err := services.Find(utils.ToInt(vars["id"]))
	if err != nil {
		return errors.New("service not found")
	}

	var hts []*hits.Hit
	query, err := database.ParseQueries(r, service.AllHits())
	if err != nil {
		return err
	}
	query.Find(&hts)
	return hts
}
