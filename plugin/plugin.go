package plugin

import (
	"database/sql"
	"html/template"
	"net/http"
)

var (
	DB         *sql.DB
	AllPlugins []Info
)

type PluginInfo struct {
	Info Info
	PluginActions
	Add
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name string
	Form string
}

func (i Info) Template() *template.Template {
	t := template.New("form")
	temp, _ := t.Parse(i.Form)
	return temp
}

type Add func(p PluginInfo)

type PluginActions interface {
	GetInfo() Info
	Routes() []Routing
	SaveForm()
	OnInstall()
	OnUninstall()
	OnFailure()
	OnHit()
	OnSettingsSaved()
	OnNewUser()
	OnNewService()
	OnShutdown()
	OnLoad()
	OnBeforeRequest()
	OnAfterRequest()
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
