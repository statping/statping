package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"os"
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

	file, err := os.Open("./statup.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 62)
	stat, err := os.Stat("./statup.log")
	start := stat.Size() - 62
	_, err = file.ReadAt(buf, start)
	if err == nil {
		fmt.Printf("%s\n", buf)
	}

	ExecuteResponse(w, r, "logs.html", nil)
}

func LogsLineHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(utils.LastLine.(string)))
}

type backups struct {
	Core types.Core `json:"core"`
}

func BackupCreateHandler(w http.ResponseWriter, r *http.Request) {
	//if !IsAuthenticated(r) {
	//	http.Redirect(w, r, "/", http.StatusSeeOther)
	//	return
	//}

	backup := backups{
		Core: *core.CoreApp.ToCore(),
	}

	out, err := json.Marshal(backup)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

	json.NewEncoder(w).Encode(backup)

}
