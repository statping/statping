package handlers

import (
	"github.com/hunterlong/statup/core"
	"net/http"
	"strings"
)

type PluginSelect struct {
	Plugin string
	Form   string
	Params map[string]interface{}
}

func PluginSavedHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	r.ParseForm()
	//vars := mux.Vars(r)
	//plug := SelectPlugin(vars["name"])
	data := make(map[string]string)
	for k, v := range r.PostForm {
		data[k] = strings.Join(v, "")
	}
	//plug.OnSave(structs.Map(data))
	http.Redirect(w, r, "/settings", http.StatusSeeOther)
}

func PluginsDownloadHandler(w http.ResponseWriter, r *http.Request) {
	auth := IsAuthenticated(r)
	if !auth {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//vars := mux.Vars(r)
	//name := vars["name"]
	//DownloadPlugin(name)
	core.LoadConfig()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}
