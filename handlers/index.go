package handlers

import (
	"github.com/razorpay/statping/types/core"
	"github.com/razorpay/statping/types/services"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if !core.App.Setup {
		ExecuteResponse(w, r, "base.gohtml", core.App, "setup")
		return
	}
	ExecuteResponse(w, r, "base.gohtml", core.App, nil)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
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
