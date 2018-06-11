package plugin

import (
	"net/http"
	"database/sql"
)

var (
	DB			 *sql.DB
)

type Plugin struct {
	PluginActions
	Name          string
	Creator 	  string
	Version       string
	InstallSQL    string
	Routes		  []*Routing
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type PluginActions interface {
	Plugin() *Plugin
	OnLoad()
	Install()
	Uninstall()
	Save()
	Form() string
	OnNewUser()
	OnFailure()
	OnHit()
}

func SetDatabase(db *sql.DB) {
	DB = db
}

func (p *Plugin) InstallPlugin(w http.ResponseWriter, r *http.Request) {

	//sql := "CREATE TABLE " + p.Name + " (enabled BOOLEAN, api_key text, api_secret text, channel text);"
	//db.QueryRow(p.InstallSQL()).Scan()

	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *Plugin) UninstallPlugin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *Plugin) SavePlugin(w http.ResponseWriter, r *http.Request) {
	//values := r.PostForm
	//p.SaveFunc(values)
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}
