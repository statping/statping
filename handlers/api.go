package handlers

import (
	"fmt"
	"github.com/statping/statping/types/checkins"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/incidents"
	"github.com/statping/statping/types/messages"
	"github.com/statping/statping/types/notifications"
	"github.com/statping/statping/types/null"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"net/http"
	"time"
)

type apiResponse struct {
	Status string      `json:"status"`
	Object string      `json:"type,omitempty"`
	Method string      `json:"method,omitempty"`
	Error  error       `json:"error,omitempty"`
	Id     int64       `json:"id,omitempty"`
	Output interface{} `json:"output,omitempty"`
}

func apiIndexHandler(r *http.Request) interface{} {
	coreClone := *core.App
	_, err := getJwtToken(r)
	if err == nil {
		coreClone.LoggedIn = true
		coreClone.IsAdmin = IsAdmin(r)
	}
	return coreClone
}

func apiRenewHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	core.App.ApiSecret = utils.NewSHA256Hash()
	err = core.App.Update()
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	output := apiResponse{
		Status: "success",
	}
	returnJson(output, w, r)
}

func apiOAuthHandler(r *http.Request) interface{} {
	app := core.App
	return app.OAuth
}

func apiCoreHandler(w http.ResponseWriter, r *http.Request) {
	var c *core.Core
	err := DecodeJSON(r, &c)
	if err != nil {
		sendErrorJson(err, w, r)
		return
	}
	app := core.App
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
	app.OAuth = c.OAuth
	app.UseCdn = null.NewNullBool(c.UseCdn.Bool)
	app.AllowReports = null.NewNullBool(c.AllowReports.Bool)
	utils.SentryInit(nil, app.AllowReports.Bool)
	err = app.Update()
	returnJson(core.App, w, r)
}

type cacheJson struct {
	URL        string    `json:"url"`
	Expiration time.Time `json:"expiration"`
	Size       int       `json:"size"`
}

func apiCacheHandler(w http.ResponseWriter, r *http.Request) {
	var cacheList []cacheJson
	for k, v := range CacheStorage.List() {
		cacheList = append(cacheList, cacheJson{
			URL:        k,
			Expiration: time.Unix(0, v.Expiration).UTC(),
			Size:       len(v.Content),
		})
	}
	returnJson(cacheList, w, r)
}

func apiClearCacheHandler(w http.ResponseWriter, r *http.Request) {
	CacheStorage.StopRoutine()
	CacheStorage = NewStorage()
	output := apiResponse{
		Status: "success",
	}
	returnJson(output, w, r)
}

func sendErrorJson(err error, w http.ResponseWriter, r *http.Request) {
	errCode := 0
	e, ok := err.(errors.Error)
	if ok {
		errCode = e.Status()
	}
	log.WithField("url", r.URL.String()).
		WithField("method", r.Method).
		WithField("code", errCode).
		Errorln(fmt.Errorf("sending error response for %s: %s", r.URL.String(), err.Error()))

	returnJson(err, w, r)
}

func sendJsonAction(obj interface{}, method string, w http.ResponseWriter, r *http.Request) {
	var objName string
	var objId int64
	switch v := obj.(type) {
	case *services.Service:
		objName = "service"
		objId = v.Id
	case *notifications.Notification:
		objName = "notifier"
		objId = v.Id
	case *core.Core:
		objName = "core"
	case *users.User:
		objName = "user"
		objId = v.Id
	case *groups.Group:
		objName = "group"
		objId = v.Id
	case *checkins.Checkin:
		objName = "checkin"
		objId = v.Id
	case *checkins.CheckinHit:
		objName = "checkin_hit"
		objId = v.Id
	case *messages.Message:
		objName = "message"
		objId = v.Id
	case *incidents.Incident:
		objName = "incident"
		objId = v.Id
	case *incidents.IncidentUpdate:
		objName = "incident_update"
		objId = v.Id
	default:
		objName = fmt.Sprintf("%T", v)
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
	w.Header().Set("Content-Type", "application/json")
	returnJson(errors.NotAuthenticated, w, r)
}
