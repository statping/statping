package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"strconv"
)

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
	ExecuteResponse(w, r, "services.html", core.CoreApp.Services)
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
	checkType := r.PostForm.Get("check_type")
	postData := r.PostForm.Get("post_data")

	service := &core.Service{
		Name:           name,
		Domain:         domain,
		Method:         method,
		Expected:       expected,
		ExpectedStatus: status,
		Interval:       interval,
		Type:           checkType,
		Port:           port,
		PostData:       postData,
	}
	_, err := service.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error starting %v check routine. %v", service.Name, err))
	}

	go service.CheckQueue()
	core.OnNewService(service)

	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func ServicesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	service.Delete()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func ServicesViewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))
	ExecuteResponse(w, r, "service.html", service)
}

func ServicesBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))

	var badge []byte
	if service.Online {
		badge = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="104" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="104" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#4c1" d="M54 0h50v20H54z"/><path fill="url(#b)" d="M0 0h104v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">` + service.Name + `</text><text x="28" y="14">` + service.Name + `</text><text x="78" y="15" fill="#010101" fill-opacity=".3">online</text><text x="78" y="14">online</text></g></svg>`)
	} else {
		badge = []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="99" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="99" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#e05d44" d="M54 0h45v20H54z"/><path fill="url(#b)" d="M0 0h99v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">` + service.Name + `</text><text x="28" y="14">` + service.Name + `</text><text x="75.5" y="15" fill="#010101" fill-opacity=".3">offline</text><text x="75.5" y="14">offline</text></g></svg>`)
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	w.Write(badge)

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
	checkType := r.PostForm.Get("check_type")
	postData := r.PostForm.Get("post_data")
	serviceUpdate := &core.Service{
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
	}
	service = service.Update(serviceUpdate)
	ExecuteResponse(w, r, "service.html", service)
}

func ServicesDeleteFailuresHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	service := core.SelectService(utils.StringInt(vars["id"]))

	service.DeleteFailures()
	core.CoreApp.Services, _ = core.SelectAllServices()
	http.Redirect(w, r, "/services", http.StatusSeeOther)
}

func CheckinCreateUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	vars := mux.Vars(r)
	interval := utils.StringInt(r.PostForm.Get("interval"))
	service := core.SelectService(utils.StringInt(vars["id"]))
	checkin := &core.Checkin{
		Service:  service.Id,
		Interval: interval,
		Api:      utils.NewSHA1Hash(18),
	}
	checkin.Create()
	fmt.Println(checkin.Create())
	ExecuteResponse(w, r, "service.html", service)
}
