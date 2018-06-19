package plugin

import (
	"fmt"
	"net/http"
	"upper.io/db.v3/lib/sqlbuilder"
)

//
//     STATUP PLUGIN INTERFACE
//
//            v0.1
//
//       https://statup.io
//
//
// An expandable plugin framework that will still
// work even if there's an update or addition.
//

var (
	DB sqlbuilder.Database
)

func SetDatabase(database sqlbuilder.Database) {
	DB = database
}

func Throw(err error) {
	fmt.Println(err)
}

type PluginInfo struct {
	Info Info
	PluginActions
}

func (p *PluginInfo) Form() string {
	return "okkokokkok"
}

type PluginActions interface {
	GetInfo() Info
	GetForm() string
	SetInfo(map[string]interface{}) Info
	Routes() []Routing
	OnSave(map[string]interface{})
	OnFailure(map[string]interface{})
	OnSuccess(map[string]interface{})
	OnSettingsSaved(map[string]interface{})
	OnNewUser(map[string]interface{})
	OnNewService(map[string]interface{})
	OnUpdatedService(map[string]interface{})
	OnDeletedService(map[string]interface{})
	OnInstall(map[string]interface{})
	OnUninstall(map[string]interface{})
	OnBeforeRequest(map[string]interface{})
	OnAfterRequest(map[string]interface{})
	OnShutdown()
	OnLoad(sqlbuilder.Database)
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name        string
	Description string
	Form 		string
}
