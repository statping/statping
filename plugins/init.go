package plugins

import (
	"database/sql"
	"net/http"
)

var (
	db           *sql.DB
	PluginRoutes []*Routing
)

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func AddRoute(url string, method string, handle func(http.ResponseWriter, *http.Request)) {
	route := &Routing{url, method, handle}
	PluginRoutes = append(PluginRoutes, route)
}

func InitDB(database *sql.DB) {
	db = database
}
