package handlers

import (
	"github.com/hunterlong/statping/core"
	"net/http"
)

func apiAllIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	incidents := core.AllIncidents()
	returnJson(incidents, w, r)
}
