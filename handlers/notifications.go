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
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

func apiNotifiersHandler(w http.ResponseWriter, r *http.Request) {
	var notifiers []*notifier.Notification
	for _, n := range core.CoreApp.Notifications {
		notif := n.(notifier.Notifier)
		notifiers = append(notifiers, notif.Select())
	}
	returnJson(notifiers, w, r)
}

func apiNotifierGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, notifierObj, err := notifier.SelectNotifier(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(notifierObj, w, r)
}

func apiNotifierUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notifer, not, err := notifier.SelectNotifier(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&notifer)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	_, err = notifier.Update(not, notifer)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifier.OnSave(notifer.Method)
	sendJsonAction(notifer, "update", w, r)
}

func testNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	form := parseForm(r)
	vars := mux.Vars(r)
	method := vars["method"]
	enabled := form.Get("enable")
	host := form.Get("host")
	port := int(utils.ToInt(form.Get("port")))
	username := form.Get("username")
	password := form.Get("password")
	var1 := form.Get("var1")
	var2 := form.Get("var2")
	apiKey := form.Get("api_key")
	apiSecret := form.Get("api_secret")
	limits := int(utils.ToInt(form.Get("limits")))

	fakeNotifer, notif, err := notifier.SelectNotifier(method)
	if err != nil {
		log.Errorln(fmt.Sprintf("issue saving notifier %v: %v", method, err))
		ExecuteResponse(w, r, "settings.gohtml", core.CoreApp, "settings")
		return
	}

	notifer := *fakeNotifer

	if host != "" {
		notifer.Host = host
	}
	if port != 0 {
		notifer.Port = port
	}
	if username != "" {
		notifer.Username = username
	}
	if password != "" && password != "##########" {
		notifer.Password = password
	}
	if var1 != "" {
		notifer.Var1 = var1
	}
	if var2 != "" {
		notifer.Var2 = var2
	}
	if apiKey != "" {
		notifer.ApiKey = apiKey
	}
	if apiSecret != "" {
		notifer.ApiSecret = apiSecret
	}
	if limits != 0 {
		notifer.Limits = limits
	}
	notifer.Enabled = types.NewNullBool(enabled == "on")

	err = notif.(notifier.Tester).OnTest()
	if err == nil {
		w.Write([]byte("ok"))
	} else {
		w.Write([]byte(err.Error()))
	}
}
