package handlers

import (
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/services"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if !core.App.Setup {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	ExecuteResponse(w, r, "base.gohtml", core.App, nil)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]interface{}{
		"services": len(services.All()),
		"online":   true,
		"setup":    core.App.Setup,
	}
	returnJson(health, w, r)
}
