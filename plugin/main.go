package plugin

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

//
//     STATUP PLUGIN INTERFACE
//
//            v0.1
//
//       https://statup.io
//

var (
	DB *sql.DB
)

func SetDatabase(database *sql.DB) {
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
	Service   int
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
