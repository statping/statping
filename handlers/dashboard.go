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
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"strconv"
	"time"
)

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println()
	if !IsAuthenticated(r) {
		err := core.ErrorResponse{}
		executeResponse(w, r, "login.html", err, nil)
	} else {
		executeResponse(w, r, "dashboard.html", core.CoreApp, nil)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if sessionStore == nil {
		resetCookies()
	}
	session, _ := sessionStore.Get(r, cookieKey)
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
		executeResponse(w, r, "login.html", err, nil)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, cookieKey)
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func helpHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	help := source.HelpMarkdown()
	executeResponse(w, r, "help.html", help, nil)
}

func logsHandler(w http.ResponseWriter, r *http.Request) {
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
	executeResponse(w, r, "logs.html", logs, nil)
}

func logsLineHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if lastLine := utils.GetLastLine(); lastLine != nil {
		w.Write([]byte(lastLine.FormatForHtml()))
	}
}

type exportData struct {
	Core      *core.Core         `json:"core"`
	Notifiers types.AllNotifiers `json:"notifiers"`
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	if !IsAuthenticated(r) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var notifiers []*notifier.Notification
	for _, v := range core.CoreApp.Notifications {
		notifier := v.(notifier.Notifier)
		notifiers = append(notifiers, notifier.Select())
	}

	data := exportData{core.CoreApp, notifiers}

	export, _ := json.Marshal(data)

	mime := http.DetectContentType(export)
	fileSize := len(string(export))

	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename=export.json")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	http.ServeContent(w, r, "export.json", time.Now(), bytes.NewReader(export))

}
