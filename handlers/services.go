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
	"sort"
	"time"
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

func apiServiceTimeDataHandler(w http.ResponseWriter, r *http.Request) {
	service, err := serviceByID(r)
	if err != nil {
		sendErrorJson(errors.New("service data not found"), w, r)
		return
	}

	allFailures := service.AllFailures()
	allHits := service.AllHits()

	tMap := make(map[time.Time]bool)
	for _, v := range allHits.List() {
		tMap[v.CreatedAt] = true
	}
	for _, v := range allFailures.List() {
		tMap[v.CreatedAt] = false
	}

	var servs []ser
	for t, v := range tMap {
		s := ser{
			Time:   t,
			Online: v,
		}
		servs = append(servs, s)
	}

	sort.Sort(ByTime(servs))

	var allTimes []series
	online := servs[0].Online
	thisTime := servs[0].Time
	for i := 0; i < len(servs); i++ {
		v := servs[i]

		if v.Online != online {
			s := series{
				Start:    thisTime,
				End:      v.Time,
				Duration: v.Time.Sub(thisTime).Milliseconds(),
				Online:   online,
			}
			allTimes = append(allTimes, s)
			thisTime = v.Time
			online = v.Online
		}
	}

	first := servs[0].Time
	last := servs[len(servs)-1].Time

	if !service.Online {
		s := series{
			Start:    allTimes[len(allTimes)-1].End,
			End:      utils.Now(),
			Duration: utils.Now().Sub(last).Milliseconds(),
			Online:   service.Online,
		}
		allTimes = append(allTimes, s)
	} else {
		l := allTimes[len(allTimes)-1]
		allTimes[len(allTimes)-1] = series{
			Start:    l.Start,
			End:      utils.Now(),
			Duration: utils.Now().Sub(l.Start).Milliseconds(),
			Online:   true,
		}
	}

	jj := uptimeSeries{
		Start:    first,
		End:      last,
		Uptime:   addDurations(allTimes, true),
		Downtime: addDurations(allTimes, false),
		Series:   allTimes,
	}

	returnJson(jj, w, r)
}

type ByTime []ser

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func addDurations(s []series, on bool) int64 {
	var dur int64
	for _, v := range s {
		if v.Online == on {
			dur += v.Duration
		}
	}
	return dur
}

type ser struct {
	Time   time.Time
	Online bool
}

func findNextFailure(m map[time.Time]bool, after time.Time, online bool) time.Time {
	for k, v := range m {
		if k.After(after) && v == online {
			return k
		}
	}
	return time.Time{}
}

//func calculateDuration(m map[time.Time]bool, on bool) time.Duration {
//	var t time.Duration
//	for t, v := range m {
//		if v == on {
//			t.
//		}
//	}
//}

type uptimeSeries struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Uptime   int64     `json:"uptime"`
	Downtime int64     `json:"downtime"`
	Series   []series  `json:"series"`
}

type series struct {
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Duration int64     `json:"duration"`
	Online   bool      `json:"online"`
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
