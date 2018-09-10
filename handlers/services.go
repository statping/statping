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
	"net/http"
	"strconv"
)

type Service struct {
	*types.Service
}

func RenderServiceChartHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=60")
	ExecuteJSResponse(w, r, "charts.js", []*core.Service{service})
}

func RenderServiceChartsHandler(w http.ResponseWriter, r *http.Request) {
	services := core.CoreApp.Services
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=60")
	ExecuteJSResponse(w, r, "charts.js", services)
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
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

func ReorderServiceHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var newOrder []*serviceOrder
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newOrder)
	for _, s := range newOrder {
		fmt.Println("updating: ", s.Id, " to be order_id: ", s.Order)
		service := core.SelectService(s.Id)
		service.Order = s.Order
		service.Update(false)
	}
	w.WriteHeader(http.StatusOK)
}

func CreateServiceHandler(w http.ResponseWriter, r *http.Request) {
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
	_, err := service.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error starting %v check routine. %v", service.Name, err))
	}
	//notifiers.OnNewService(core.ReturnService(service.Service))
	ExecuteResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}

func ServicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
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
	ExecuteResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}

func ServicesViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	if serv == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ExecuteResponse(w, r, "service.html", serv, nil)
}

func ServicesUpdateHandler(w http.ResponseWriter, r *http.Request) {
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
	ExecuteResponse(w, r, "service.html", service, "/services")
}

func ServicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	service.DeleteFailures()
	ExecuteResponse(w, r, "services.html", core.CoreApp.Services, "/services")
}

func CheckinCreateUpdateHandler(w http.ResponseWriter, r *http.Request) {
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
	ExecuteResponse(w, r, "service.html", service, "/services")
}
