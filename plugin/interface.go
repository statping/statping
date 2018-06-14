package plugin

import (
	"net/http"
	"time"
)

type PluginActions interface {
	GetInfo() Info
	Routes() []Routing
	OnSave(map[string]string)
	OnInstall()
	OnUninstall()
	OnFailure(*Service)
	OnHit(*Service)
	OnSettingsSaved()
	OnNewUser()
	OnNewService()
	OnShutdown()
	OnLoad()
	OnBeforeRequest()
	OnAfterRequest()
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

type PluginInfo struct {
	Info Info
	PluginActions
}

type Routing struct {
	URL     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

type Info struct {
	Name        string
	Description string
	Form        []FormElement
}

type FormElement struct {
	Name        string
	Description string
	SQLValue    string
	SQLType     string
	Value       interface{}
}
