package handlers

import (
	"github.com/gorilla/mux"
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/errors"
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

func findService(r *http.Request) (*services.Service, error) {
	vars := mux.Vars(r)
	id := utils.ToInt(vars["id"])
	servicer, err := services.Find(id)
	if err != nil {
		return nil, err
	}
	if !servicer.Public.Bool && !IsReadAuthenticated(r) {
		return nil, errors.NotAuthenticated
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
			sendErrorJson(err, w, r)
			return
		}
		service.Order = s.Order
		service.Update()
	}
	returnJson(newOrder, w, r)
}

func apiServiceHandler(r *http.Request) interface{} {
	srv, err := findService(r)
	if err != nil {
		return err
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
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	if err := DecodeJSON(r, &service); err != nil {
		sendErrorJson(err, w, r)
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

func apiServiceDataHandler(w http.ResponseWriter, r *http.Request) {
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
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
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
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
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
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

func apiServiceTimeDataHandler(w http.ResponseWriter, r *http.Request) {
	service, err := findService(r)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	groupHits, err := database.ParseQueries(r, service.AllHits())
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	groupFailures, err := database.ParseQueries(r, service.AllFailures())
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	var allFailures []*failures.Failure
	var allHits []*hits.Hit

	if err := groupHits.Find(&allHits); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	if err := groupFailures.Find(&allFailures); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	uptimeData, err := service.UptimeData(allHits, allFailures)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(uptimeData, w, r)
}

func apiServiceDeleteHandler(w http.ResponseWriter, r *http.Request) {
	service, err := findService(r)
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
	var srvs []services.Service
	for _, v := range services.AllInOrder() {
		if !v.Public.Bool && !IsUser(r) {
			continue
		}
		srvs = append(srvs, v)
	}
	return srvs
}

func servicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	service, err := findService(r)
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
	service, err := findService(r)
	if err != nil {
		return err
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
	service, err := findService(r)
	if err != nil {
		return err
	}
	var hts []*hits.Hit
	query, err := database.ParseQueries(r, service.AllHits())
	if err != nil {
		return err
	}
	query.Find(&hts)
	return hts
}
