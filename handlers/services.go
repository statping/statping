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
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/jinzhu/now"
	"net/http"
	"strconv"
	"time"
)

type Service struct {
	*types.Service
}

func renderServiceChartHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fields := parseGet(r)
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=30")

	startField := fields.Get("start")
	endField := fields.Get("end")

	end := now.EndOfDay().UTC()
	start := now.BeginningOfDay().UTC()

	if startField != "" {
		start = time.Unix(utils.StringInt(startField), 0)
		start = now.New(start).BeginningOfDay().UTC()
	}
	if endField != "" {
		end = time.Unix(utils.StringInt(endField), 0)
		end = now.New(end).EndOfDay().UTC()
	}

	fmt.Println("start: ", start.String(), "end: ", end.String())

	service := core.SelectService(utils.StringInt(vars["id"]))
	data := core.GraphDataRaw(service, start, end).ToString()

	out := struct {
		Services []*core.Service
		Data     []string
	}{[]*core.Service{service}, []string{data}}

	executeJSResponse(w, r, "charts.js", out)
}

func renderServiceChartsHandler(w http.ResponseWriter, r *http.Request) {
	services := core.CoreApp.Services
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=60")

	var data []string
	end := now.EndOfDay().UTC()
	start := now.BeginningOfDay().UTC()

	for _, s := range services {
		d := core.GraphDataRaw(s, start, end).ToString()
		data = append(data, d)
	}

	out := struct {
		Services []types.ServiceInterface
		Data     []string
	}{services, data}

	executeJSResponse(w, r, "charts.js", out)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	executeResponse(w, r, "services.html", core.CoreApp.Services, nil)
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
	decoder.Decode(&newOrder)
	for _, s := range newOrder {
		service := core.SelectService(s.Id)
		service.Order = s.Order
		service.Update(false)
	}
	w.WriteHeader(http.StatusOK)
}

func createServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	name := r.PostForm.Get("name")
	domain := r.PostForm.Get("domain")
	method := r.PostForm.Get("method")
	expected := r.PostForm.Get("expected")
	status, _ := strconv.Atoi(r.PostForm.Get("expected_status"))
	interval, _ := strconv.Atoi(r.PostForm.Get("interval"))
	port, _ := strconv.Atoi(r.PostForm.Get("port"))
	timeout, _ := strconv.Atoi(r.PostForm.Get("timeout"))
	checkType := r.PostForm.Get("check_type")
	postData := r.PostForm.Get("post_data")
	order, _ := strconv.Atoi(r.PostForm.Get("order"))

	if checkType == "http" && status == 0 {
		status = 200
	}

	service := core.ReturnService(&types.Service{
		Name:           name,
		Domain:         domain,
		Method:         method,
		Expected:       expected,
		ExpectedStatus: status,
		Interval:       interval,
		Type:           checkType,
		Port:           port,
		PostData:       postData,
		Timeout:        timeout,
		Order:          order,
	})
	_, err := service.Create(true)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error starting %v check routine. %v", service.Name, err))
	}
	//notifiers.OnNewService(core.ReturnService(service.Service))
	executeResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}

func servicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	if service == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	service.Delete()
	executeResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}

func servicesViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fields := parseGet(r)
	startField := utils.StringInt(fields.Get("start"))
	endField := utils.StringInt(fields.Get("end"))
	serv := core.SelectService(utils.StringInt(vars["id"]))
	if serv == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	end := time.Now()
	start := end.Add((-24 * 7) * time.Hour)

	if startField != 0 {
		start = time.Unix(startField, 0)
	}
	if endField != 0 {
		end = time.Unix(endField, 0)
	}

	data := core.GraphDataRaw(serv, start, end)

	out := struct {
		Service *core.Service
		Start   int64
		End     int64
		Data    string
	}{serv, start.Unix(), end.Unix(), data.ToString()}

	executeResponse(w, r, "service.html", out, nil)
}

func servicesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	r.ParseForm()
	name := r.PostForm.Get("name")
	domain := r.PostForm.Get("domain")
	method := r.PostForm.Get("method")
	expected := r.PostForm.Get("expected")
	status, _ := strconv.Atoi(r.PostForm.Get("expected_status"))
	interval, _ := strconv.Atoi(r.PostForm.Get("interval"))
	port, _ := strconv.Atoi(r.PostForm.Get("port"))
	timeout, _ := strconv.Atoi(r.PostForm.Get("timeout"))
	checkType := r.PostForm.Get("check_type")
	postData := r.PostForm.Get("post_data")
	order, _ := strconv.Atoi(r.PostForm.Get("order"))

	service.Name = name
	service.Domain = domain
	service.Method = method
	service.ExpectedStatus = status
	service.Expected = expected
	service.Interval = interval
	service.Type = checkType
	service.Port = port
	service.PostData = postData
	service.Timeout = timeout
	service.Order = order

	service.Update(true)
	service.Check(true)
	executeResponse(w, r, "service.html", service, "/services")
}

func servicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	service.DeleteFailures()
	executeResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}

func checkinCreateUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	interval := utils.StringInt(r.PostForm.Get("interval"))
	serv := core.SelectService(utils.StringInt(vars["id"]))
	service := serv
	checkin := &types.Checkin{
		Service:  service.Id,
		Interval: interval,
		Api:      utils.NewSHA1Hash(18),
	}
	checkin.Create()
	executeResponse(w, r, "service.html", service, "/services")
}
