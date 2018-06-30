package handlers

import (
	"fmt"
	"github.com/hunterlong/statup/core"
	"net/http"
)

type index struct {
	Core     core.Core
	Services []*core.Service
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if core.CoreApp == nil {
		http.Redirect(w, r, "/setup", http.StatusSeeOther)
		return
	}
	out := index{*core.CoreApp, core.CoreApp.Services}
	first, _ := out.Services[0].LimitedHits()
	fmt.Println(out.Services[0].Name, "start:", first[0].Id, "last:", first[len(first)-1].Id)
	ExecuteResponse(w, r, "index.html", out)
}
