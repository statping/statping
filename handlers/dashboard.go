// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"bytes"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"strconv"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	if !IsUser(r) {
		err := core.ErrorResponse{}
		ExecuteResponse(w, r, "login.gohtml", err, nil)
	} else {
		ExecuteResponse(w, r, "dashboard.gohtml", core.CoreApp, nil)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if sessionStore == nil {
		resetCookies()
	}
	session, _ := sessionStore.Get(r, cookieKey)
	form := parseForm(r)
	username := form.Get("username")
	password := form.Get("password")
	user, auth := core.AuthUser(username, password)
	if auth {
		session.Values["authenticated"] = true
		session.Values["user_id"] = user.Id
		session.Values["admin"] = user.Admin.Bool
		session.Save(r, w)
		utils.Log.Infoln(fmt.Sprintf("User %v logged in from IP %v", user.Username, r.RemoteAddr))
		http.Redirect(w, r, basePath+"dashboard", http.StatusSeeOther)
	} else {
		err := core.ErrorResponse{Error: "Incorrect login information submitted, try again."}
		ExecuteResponse(w, r, "login.gohtml", err, nil)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, cookieKey)
	session.Values["authenticated"] = false
	session.Values["admin"] = false
	session.Values["user_id"] = 0
	session.Save(r, w)
	http.Redirect(w, r, basePath, http.StatusSeeOther)
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	if !IsUser(r) {
		http.Redirect(w, r, basePath, http.StatusSeeOther)
		return
	}
	help := source.HelpMarkdown()
	ExecuteResponse(w, r, "help.gohtml", help, nil)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
	utils.LockLines.Lock()
	logs := make([]string, 0)
	length := len(utils.LastLines)
	// We need string log lines from end to start.
	for i := length - 1; i >= 0; i-- {
		logs = append(logs, utils.LastLines[i].FormatForHtml()+"\r\n")
	}
	utils.LockLines.Unlock()
	ExecuteResponse(w, r, "logs.gohtml", logs, nil)
}

func logsLineHandler(w http.ResponseWriter, r *http.Request) {
	if lastLine := utils.GetLastLine(); lastLine != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(lastLine.FormatForHtml()))
	}
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	var notifiers []*notifier.Notification
	for _, v := range core.CoreApp.Notifications {
		notifier := v.(notifier.Notifier)
		notifiers = append(notifiers, notifier.Select())
	}

	export, _ := core.ExportSettings()

	mime := http.DetectContentType(export)
	fileSize := len(string(export))

	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename=export.json")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, "export.json", utils.Now(), bytes.NewReader(export))

}
