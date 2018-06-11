package plugin

import (
	"net/http"
	"database/sql"
)

var (
	DB			 *sql.DB
)

type PluginInfo struct {
	PluginActions
	Name          string
	Creator 	  string
	Version       string
	InstallSQL    string
	Form          string
	Routes		  []*Routing
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type PluginActions interface {
	Plugin() *PluginInfo
	SaveForm()
	OnInstall()
	OnUninstall()
	OnFailure()
	OnHit()
	OnSettingsSaved()
	OnNewUser()
	OnShutdown()
	OnLoad()
}

func SetDatabase(db *sql.DB) {
	DB = db
}

func (p *PluginInfo) InstallPlugin(w http.ResponseWriter, r *http.Request) {

	//sql := "CREATE TABLE " + p.Name + " (enabled BOOLEAN, api_key text, api_secret text, channel text);"
	//db.QueryRow(p.InstallSQL()).Scan()

	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *PluginInfo) UninstallPlugin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *PluginInfo) SavePlugin(w http.ResponseWriter, r *http.Request) {
	//values := r.PostForm
	//p.SaveFunc(values)
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}
