package plugin

import (
	"github.com/hunterlong/statup/core/notifier"
	"github.com/jinzhu/gorm"
	"net/http"
)

type PluginObject struct {
	PluginInfo
}

type Pluginer interface {
	Select() *PluginObject
}

type Databaser interface {
	StatupDatabase(*gorm.DB)
}

type Router interface {
	Routes() []interface{}
	AddRoute(string, string, http.HandlerFunc) error
}

type Asseter interface {
	Asset(string) ([]byte, error)
}

type Notifier interface {
	notifier.Notifier
	notifier.BasicEvents
}

type AdvancedNotifier interface {
	notifier.Notifier
	notifier.BasicEvents
	notifier.UserEvents
	notifier.CoreEvents
	notifier.NotifierEvents
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name        string
	Description string
	Form        string
}

type Database *gorm.DB

type Plugin struct {
	Name        string
	Description string
}

type PluginDatabase interface {
	Database(gorm.DB)
	Update() error
}

type PluginInfo struct {
	i *Info
}
