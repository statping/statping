package plugin

import (
	"fmt"
	"net/http"
	"time"
	"upper.io/db.v3/lib/sqlbuilder"
)

//
//     STATUP PLUGIN INTERFACE
//
//            v0.1
//
//       https://statup.io
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

type PluginActions interface {
	GetInfo() Info
	SetInfo(map[string]string) Info
	Routes() []Routing
	OnSave(map[string]string)
	OnFailure(*Service)
	OnSuccess(*Service)
	OnSettingsSaved(map[string]string)
	OnNewUser(*User)
	OnNewService(*Service)
	OnUpdatedService(*Service)
	OnDeletedService(*Service)
	OnInstall()
	OnUninstall()
	OnShutdown()
	OnLoad()
	OnBeforeRequest()
	OnAfterRequest()
}

type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

type Service struct {
	Id             int64
	Name           string
	Domain         string
	Expected       string
	ExpectedStatus int
	Interval       int
	Method         string
	Port           int
	CreatedAt      time.Time
	Data           string
	Online         bool
	Latency        float64
	Online24Hours  float32
	AvgResponse    string
	TotalUptime    string
	Failures       []*Failure
}

type Failure struct {
	Id        int
	Issue     string
	Service   int64
	CreatedAt time.Time
	Ago       string
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name        string
	Description string
	Form        []*FormElement
}

type FormElement struct {
	Name        string
	Description string
	InputName   string
	InputType   string
	Value       string
}
