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
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"net/http"
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
	if Store == nil {
		resetCookies()
	}
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
	ExecuteResponse(w, r, "logs.html", nil)
}

func LogsLineHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	switch v := utils.LastLine.(type) {
	case string:
		w.Write([]byte(v))
	case error:
		w.Write([]byte(v.Error()))
	case []byte:
		w.Write(v)
	}
}
