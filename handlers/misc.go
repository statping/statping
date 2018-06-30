package handlers

import (
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/utils"
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

func FavIconHandler(w http.ResponseWriter, r *http.Request) {
	data, err := core.TmplBox.String("favicon.ico")
	if err != nil {
		utils.Log(2, err)
	}
	w.Header().Set("Content-Type", "image/x-icon")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}
