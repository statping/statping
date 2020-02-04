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
	"errors"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"time"
)

type apiResponse struct {
	Status string      `json:"status"`
	Object string      `json:"type,omitempty"`
	Method string      `json:"method,omitempty"`
	Error  string      `json:"error,omitempty"`
	Id     int64       `json:"id,omitempty"`
	Output interface{} `json:"output,omitempty"`
}

func apiIndexHandler(r *http.Request) interface{} {
	coreClone := *core.CoreApp
	return *coreClone.ToCore()
}

func apiRenewHandler(w http.ResponseWriter, r *http.Request) {
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

func apiCoreHandler(w http.ResponseWriter, r *http.Request) {
	var c *core.Core
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	app := core.CoreApp
	if c.Name != "" {
		app.Name = c.Name
	}
	if c.Description != app.Description {
		app.Description = c.Description
	}
	if c.Style != app.Style {
		app.Style = c.Style
	}
	if c.Footer.String != app.Footer.String {
		app.Footer = c.Footer
	}
	if c.Domain != app.Domain {
		app.Domain = c.Domain
	}
	if c.Timezone != app.Timezone {
		app.Timezone = c.Timezone
	}
	app.UpdateNotify = c.UpdateNotify
	app.UseCdn = types.NewNullBool(c.UseCdn.Bool)
	core.CoreApp, err = core.UpdateCore(app)
	returnJson(core.CoreApp, w, r)
}

type cacheJson struct {
	URL        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
	Size       int       `json:"size"`
}

func apiCacheHandler(w http.ResponseWriter, r *http.Request) {
	cache := CacheStorage
	var cacheList []cacheJson
	for k, v := range cache.List() {
		cacheList = append(cacheList, cacheJson{
			URL:        k,
			Expiration: time.Unix(0, v.Expiration).UTC(),
			Size:       len(v.Content),
		})
	}
	returnJson(cacheList, w, r)
}

func apiClearCacheHandler(w http.ResponseWriter, r *http.Request) {
	CacheStorage = NewStorage()
	go CleanRoutine()
	output := apiResponse{
		Status: "success",
	}
	returnJson(output, w, r)
}

func sendErrorJson(err error, w http.ResponseWriter, r *http.Request) {
	log.Warnln(fmt.Errorf("sending error response for %v: %v", r.URL.String(), err.Error()))
	output := apiResponse{
		Status: "error",
		Error:  err.Error(),
	}
	returnJson(output, w, r)
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
	case *types.Group:
		objName = "group"
		objId = v.Id
	case *core.Group:
		objName = "group"
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
	case *core.Incident:
		objName = "incident"
		objId = v.Id
	case *core.IncidentUpdate:
		objName = "incident_update"
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

	returnJson(output, w, r)
}

func sendUnauthorizedJson(w http.ResponseWriter, r *http.Request) {
	output := apiResponse{
		Status: "error",
		Error:  errors.New("not authorized").Error(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	returnJson(output, w, r)
}
