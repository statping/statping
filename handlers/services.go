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

func RenderServiceChartsHandler(w http.ResponseWriter, r *http.Request) {
	services := core.CoreApp.Services()
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Cache-Control", "max-age=60")
	ExecuteJSResponse(w, r, "charts.js", services)
}

func ServicesHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "services.html", core.CoreApp.DbServices)
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

	service := &core.Service{Service: &types.Service{
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
	}}
	_, err := service.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error starting %v check routine. %v", service.Name, err))
	}

	go service.CheckQueue(true)
	core.OnNewService(service)

	ExecuteResponse(w, r, "services.html", core.CoreApp.DbServices)
}

func ServicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	if serv == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	service := serv
	service.Delete()
	ExecuteResponse(w, r, "services.html", core.CoreApp.DbServices)
}

func ServicesViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	if serv == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ExecuteResponse(w, r, "service.html", serv)
}

func ServicesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	service := serv
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
	serviceUpdate := &core.Service{Service: &types.Service{
		Id:             service.Id,
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
	}}
	serviceUpdate.Update()
	core.CoreApp.SelectAllServices()

	serv = core.SelectService(serviceUpdate.Id)
	ExecuteResponse(w, r, "service.html", serv)
}

func ServicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	serv := core.SelectService(utils.StringInt(vars["id"]))
	service := serv
	core.DeleteFailures(service)
	core.CoreApp.SelectAllServices()
	ExecuteResponse(w, r, "services.html", core.CoreApp.DbServices)
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
	fmt.Println(checkin.Create())
	ExecuteResponse(w, r, "service.html", service)
}
