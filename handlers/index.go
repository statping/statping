package handlers

import (
	"github.com/hunterlong/statup/core"
	"net/http"
)

type index struct {
	Core     core.Core
	Services []*core.Service
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if core.CoreApp.Services == nil {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	out := index{*core.CoreApp, core.CoreApp.Services}
	ExecuteResponse(w, r, "index.html", out)
}
