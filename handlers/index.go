package handlers

import (
	"github.com/hunterlong/statup/core"
	"net/http"
)

type index struct {
	Core *core.Core
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if core.CoreApp.DbConnection == "" {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "index.html", core.CoreApp)
}
