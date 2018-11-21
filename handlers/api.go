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
	"encoding/json"
	"errors"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"net/http"
	"os"
	"strings"
)

type apiResponse struct {
	Status string      `json:"status"`
	Object string      `json:"type,omitempty"`
	Method string      `json:"method,omitempty"`
	Error  string      `json:"error,omitempty"`
	Id     int64       `json:"id,omitempty"`
	Output interface{} `json:"output,omitempty"`
}

func apiIndexHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(core.CoreApp)
}

func apiRenewHandler(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(r) {
		sendUnauthorizedJson(w, r)
		return
	}
	var err error
	core.CoreApp.ApiKey = utils.NewSHA1Hash(40)
	core.CoreApp.ApiSecret = utils.NewSHA1Hash(40)
	core.CoreApp, err = core.UpdateCore(core.CoreApp)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func sendErrorJson(err error, w http.ResponseWriter, r *http.Request) {
	output := apiResponse{
		Status: "error",
		Error:  err.Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func sendJsonAction(obj interface{}, method string, w http.ResponseWriter, r *http.Request) {
	var objName string
	var objId int64
	switch v := obj.(type) {
	case types.ServiceInterface:
		objName = "service"
		objId = v.Select().Id
	case *notifier.Notification:
		objName = "notifier"
		objId = v.Id
	case *core.Core, *types.Core:
		objName = "core"
	case *types.User:
		objName = "user"
		objId = v.Id
	case *core.User:
		objName = "user"
		objId = v.Id
	case *core.Checkin:
		objName = "checkin"
		objId = v.Id
	case *core.CheckinHit:
		objName = "checkin_hit"
		objId = v.Id
	case *types.Message:
		objName = "message"
		objId = v.Id
	case *core.Message:
		objName = "message"
		objId = v.Id
	case *types.Checkin:
		objName = "checkin"
		objId = v.Id
	default:
		objName = "missing"
	}

	output := apiResponse{
		Object: objName,
		Method: method,
		Id:     objId,
		Status: "success",
		Output: obj,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func sendUnauthorizedJson(w http.ResponseWriter, r *http.Request) {
	output := apiResponse{
		Status: "error",
		Error:  errors.New("not authorized").Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(output)
}

func isAuthorized(r *http.Request) bool {
	utils.Http(r)
	if os.Getenv("GO_ENV") == "test" {
		return true
	}
	if IsAuthenticated(r) {
		return true
	}
	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) >= 1 {
		token = tokens[0]
		token = strings.TrimPrefix(token, "Bearer ")
	}
	if token == core.CoreApp.ApiSecret {
		return true
	}
	return false
}
