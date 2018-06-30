package handlers

import (
	"github.com/hunterlong/statup/core"
	"net/http"
)

func RobotsTxtHandler(w http.ResponseWriter, r *http.Request) {
	robots := []byte(`User-agent: *
Disallow: /login
Disallow: /dashboard

Host: ` + core.CoreApp.Domain)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(robots))
}
