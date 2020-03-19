// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/statping/statping
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
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/utils"
	"net"
	"net/http"
)

func apiAllCheckinsHandler(w http.ResponseWriter, r *http.Request) {
	chks := checkins.All()
	returnJson(chks, w, r)
}

func apiCheckinHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin, err := checkins.FindByAPI(vars["api"])
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	returnJson(checkin, w, r)
}

func checkinCreateHandler(w http.ResponseWriter, r *http.Request) {
	var checkin *checkins.Checkin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&checkin)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	service, err := services.Find(checkin.ServiceId)
	if err != nil {
		sendErrorJson(fmt.Errorf("missing service_id field"), w, r)
		return
	}
	checkin.ServiceId = service.Id
	err = checkin.Create()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	sendJsonAction(checkin, "create", w, r)
}

func checkinHitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin, err := checkins.FindByAPI(vars["api"])
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	hit := &checkins.CheckinHit{
		Checkin:   checkin.Id,
		From:      ip,
		CreatedAt: utils.Now().UTC(),
	}
	err = hit.Create()
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}
	checkin.Failing = false
	checkin.LastHitTime = utils.Now().UTC()
	sendJsonAction(hit.Id, "update", w, r)
}

func checkinDeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	checkin, err := checkins.FindByAPI(vars["api"])
	if err != nil {
		sendErrorJson(fmt.Errorf("checkin %v was not found", vars["api"]), w, r)
		return
	}

	if err := checkin.Delete(); err != nil {
		sendErrorJson(err, w, r)
		return
	}

	sendJsonAction(checkin, "delete", w, r)
}
