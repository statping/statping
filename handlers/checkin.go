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
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net"
	"net/http"
)

func apiAllCheckinsHandler(w http.ResponseWriter, r *http.Request) {
	checkins := core.AllCheckins()
	for _, c := range checkins {
		c.Hits = c.AllHits()
		c.Failures = c.LimitedFailures(64)
	}
	returnJson(checkins, w, r)
}

func apiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin := core.SelectCheckin(vars["api"])
	if checkin == nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	checkin.Hits = checkin.LimitedHits(32)
	checkin.Failures = checkin.LimitedFailures(32)
	returnJson(checkin, w, r)
}

func checkinCreateHandler(w http.ResponseWriter, r *http.Request) {
	var checkin *core.Checkin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&checkin)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	service := core.SelectService(checkin.ServiceId)
	if service == nil {
		sendErrorJson(fmt.Errorf("missing service_id field"), w, r)
		return
	}
	_, err = checkin.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(checkin, "create", w, r)
}

func checkinHitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin := core.SelectCheckin(vars["api"])
	if checkin == nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	hit := &types.CheckinHit{
		Checkin:   checkin.Id,
		From:      ip,
		CreatedAt: utils.Now().UTC(),
	}
	checkinHit := core.ReturnCheckinHit(hit)
	if checkin.Last() == nil {
		checkin.Start()
		go checkin.Routine()
	}
	_, err := checkinHit.Create()
	checkin.Hits = append(checkin.Hits, checkinHit.CheckinHit)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	checkin.Failing = false
	checkin.LastHit = utils.Timezoner(utils.Now().UTC(), core.CoreApp.Timezone)
	sendJsonAction(checkinHit, "update", w, r)
}

func checkinDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin := core.SelectCheckin(vars["api"])
	if checkin == nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	err := checkin.Delete()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(checkin, "delete", w, r)
}
