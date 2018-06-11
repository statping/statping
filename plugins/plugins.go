package plugins

import (
	"database/sql"
	"fmt"
	"net/http"
)

var (
	db           *sql.DB
	PluginRoutes []*Routing
	Plugins      []*Plugin
)

type Plugin struct {
	Name          string
	InstallSQL    string
	InstallFunc   func()
	UninstallFunc func()
	SaveFunc      func()
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func (p *Plugin) Add() {
	Plugins = append(Plugins, p)
}

func AddRoute(url string, method string, handle func(http.ResponseWriter, *http.Request)) {
	route := &Routing{url, method, handle}
	PluginRoutes = append(PluginRoutes, route)
}

func (p *Plugin) InstallPlugin(w http.ResponseWriter, r *http.Request) {
	p.InstallFunc()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *Plugin) UninstallPlugin(w http.ResponseWriter, r *http.Request) {
	p.UninstallFunc()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func (p *Plugin) SavePlugin(w http.ResponseWriter, r *http.Request) {
	p.SaveFunc()
	http.Redirect(w, r, "/plugins", http.StatusSeeOther)
}

func Authenticated(r *http.Request) bool {

	return true
}

func log(msg ...string) {
	fmt.Println(" @plugins: ", msg)
}

func InitDB(database *sql.DB) {
	db = database
}
