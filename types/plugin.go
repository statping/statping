package types

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Plugin struct {
	Name        string
	Description string
}

type PluginObject struct {
	Pluginer
}

type PluginActions interface {
	GetInfo() *Info
	OnLoad() error
}

type PluginRepos struct {
	Plugins []PluginJSON
}

type PluginJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Repo        string `json:"repo"`
	Author      string `json:"author"`
	Namespace   string `json:"namespace"`
}

type Info struct {
	Name        string
	Description string
	Form        string
}

type PluginInfo struct {
	Info   *Info
	Routes []*PluginRoute
}

type PluginRouting struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Pluginer interface {
	Select() *Plugin
}

type Databaser interface {
	StatpingDatabase(*gorm.DB)
}

type Router interface {
	Routes() []*PluginRoute
	AddRoute(string, string, http.HandlerFunc) error
}

type Asseter interface {
	Asset(string) ([]byte, error)
}

type PluginRoute struct {
	Url    string
	Method string
	Func   http.HandlerFunc
}

//
//type Notifier interface {
//	notifier.Notifier
//	notifier.BasicEvents
//}
