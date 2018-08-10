package handlers

import (
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"time"
	)

type dashboard struct {
	Services        []*core.Service
	Core            *core.Core
	CountOnline     int
	CountServices   int
	Count24Failures uint64
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		err := core.ErrorResponse{}
		ExecuteResponse(w, r, "login.html", err)
	} else {
		fails := core.CountFailures()
		out := dashboard{core.CoreApp.Services, core.CoreApp, core.CountOnline(), len(core.CoreApp.Services), fails}
		ExecuteResponse(w, r, "dashboard.html", out)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, COOKIE_KEY)
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	user, auth := core.AuthUser(username, password)
	if auth {
		session.Values["authenticated"] = true
		session.Values["user_id"] = user.Id
		session.Save(r, w)
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		err := core.ErrorResponse{Error: "Incorrect login information submitted, try again."}
		ExecuteResponse(w, r, "login.html", err)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, COOKIE_KEY)
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "help.html", nil)
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	logs := make([]string, 0)
	if len(utils.LastLines) > 0 {
		for _, line := range utils.LastLines {
			logs = append([]string{compileLogLine(line.Log, &line.Date, "")}, logs...)
		}
	} else {
		logs = append(logs, "No log lines.")
	}
	ExecuteResponse(w, r, "logs.html", logs)
}

func LogsLineHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	line := compileLogLine(utils.LastLine.Log, &utils.LastLine.Date, "")
	w.Write([]byte(line))
}

func compileLogLine(line interface{}, date *time.Time, suffix string) (res string) {
	switch v := line.(type) {
	case string:
		res = string(v)
	case error:
		res = v.Error()
	case []byte:
		res = string(v)
	default:
		res = ""
	}

	if date != nil {
		res = date.Format(time.RFC3339)+" - "+res
	}

	if suffix != "" {
		res = (res)+suffix
	}

	return
}
