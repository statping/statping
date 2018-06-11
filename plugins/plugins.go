package plugins

import (
	"database/sql"
	"net/http"
	"fmt"
)

var (
	db           *sql.DB
	PluginRoutes []*Routing
	Plugins []*Plugin
)

type Plugin struct {
	Name string
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Add(name string) {
	plugin := &Plugin{name}
	Plugins = append(Plugins, plugin)
}

func AddRoute(url string, method string, handle func(http.ResponseWriter, *http.Request)) {
	route := &Routing{url, method, handle}
	PluginRoutes = append(PluginRoutes, route)
}

func Authenticated(r *http.Request) bool {


	return true
}

func log(msg... string) {
	fmt.Println(" @plugins: ",msg)
}

func InitDB(database *sql.DB) {
	db = database
}
