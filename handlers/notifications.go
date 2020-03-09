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
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/types/notifications"
	"github.com/hunterlong/statping/types/null"
	"github.com/hunterlong/statping/utils"
	"net/http"
)

func apiNotifiersHandler(w http.ResponseWriter, r *http.Request) {
	all := notifications.All()
	returnJson(all, w, r)
}

func apiNotifierGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notifier, err := notifications.Find(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	returnJson(notifier, w, r)
}

func apiNotifierUpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	notifer, err := notifications.Find(vars["notifier"])
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}

	n := notifer.Select()

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&n)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	err = n.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	notifications.OnSave(n.Method)
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

	notifier, err := notifications.Find(method)
	if err != nil {
		log.Errorln(fmt.Sprintf("issue saving notifier %v: %v", method, err))
		ExecuteResponse(w, r, "settings.gohtml", core.App, "settings")
		return
	}

	n := notifier.Select()

	if host != "" {
		n.Host = host
	}
	if port != 0 {
		n.Port = port
	}
	if username != "" {
		n.Username = username
	}
	if password != "" && password != "##########" {
		n.Password = password
	}
	if var1 != "" {
		n.Var1 = var1
	}
	if var2 != "" {
		n.Var2 = var2
	}
	if apiKey != "" {
		n.ApiKey = apiKey
	}
	if apiSecret != "" {
		n.ApiSecret = apiSecret
	}
	if limits != 0 {
		n.Limits = limits
	}
	n.Enabled = null.NewNullBool(enabled == "on")

	//err = notifications.OnTest(notifier)
	//if err == nil {
	//	w.Write([]byte("ok"))
	//} else {
	//	w.Write([]byte(err.Error()))
	//}
}
