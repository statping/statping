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
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	if !IsAuthenticated(r) {
		err := core.ErrorResponse{}
		ExecuteResponse(w, r, "login.html", err, nil)
	} else {
		ExecuteResponse(w, r, "dashboard.html", core.CoreApp, nil)
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
		ExecuteResponse(w, r, "login.html", err, nil)
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
	ExecuteResponse(w, r, "help.html", nil, nil)
}

func LogsHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	utils.LockLines.Lock()
	logs := make([]string, 0)
	len := len(utils.LastLines)
	// We need string log lines from end to start.
	for i := len - 1; i >= 0; i-- {
		logs = append(logs, utils.LastLines[i].FormatForHtml()+"\r\n")
	}
	utils.LockLines.Unlock()
	ExecuteResponse(w, r, "logs.html", logs, nil)
}

func LogsLineHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if lastLine := utils.GetLastLine(); lastLine != nil {
		w.Write([]byte(lastLine.FormatForHtml()))
	}
}
