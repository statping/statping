package handlers

import (
	"net/http"
)

func Error404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	ExecuteResponse(w, r, "error_404.html", nil)
}