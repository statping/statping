package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"html/template"
	"net/http"
	"reflect"
	"time"
)

var (
	basePath = "/"
)

var handlerFuncs = func(w http.ResponseWriter, r *http.Request) template.FuncMap {
	return template.FuncMap{
		"js": func(html interface{}) template.JS {
			return template.JS(utils.ToString(html))
		},
		"safe": func(html string) template.HTML {
			return template.HTML(html)
		},
		"safeURL": func(u string) template.URL {
			return template.URL(u)
		},
		"Auth": func() bool {
			return IsFullAuthenticated(r)
		},
		"IsUser": func() bool {
			return IsUser(r)
		},
		"VERSION": func() string {
			return core.VERSION
		},
		"CoreApp": func() *core.Core {
			return core.CoreApp
		},
		"Services": func() []types.ServiceInterface {
			return core.CoreApp.Services
		},
		"VisibleServices": func() []*core.Service {
			auth := IsUser(r)
			return core.SelectServices(auth)
		},
		"VisibleGroupServices": func(group *core.Group) []*core.Service {
			auth := IsUser(r)
			return group.VisibleServices(auth)
		},
		"Groups": func(includeAll bool) []*core.Group {
			auth := IsUser(r)
			return core.SelectGroups(includeAll, auth)
		},
		"Group": func(id int) *core.Group {
			return core.SelectGroup(int64(id))
		},
		"len": func(g interface{}) int {
			val := reflect.ValueOf(g)
			return val.Len()
		},
		"IsNil": func(g interface{}) bool {
			return g == nil
		},
		"USE_CDN": func() bool {
			return core.CoreApp.UseCdn.Bool
		},
		"UPDATENOTIFY": func() bool {
			return core.CoreApp.UpdateNotify.Bool
		},
		"QrAuth": func() string {
			return fmt.Sprintf("statping://setup?domain=%v&api=%v", core.CoreApp.Domain, core.CoreApp.ApiSecret)
		},
		"Type": func(g interface{}) []string {
			fooType := reflect.TypeOf(g)
			var methods []string
			methods = append(methods, fooType.String())
			for i := 0; i < fooType.NumMethod(); i++ {
				method := fooType.Method(i)
				fmt.Println(method.Name)
				methods = append(methods, method.Name)
			}
			return methods
		},
		"ToJSON": func(g interface{}) template.HTML {
			data, _ := json.Marshal(g)
			return template.HTML(string(data))
		},
		"underscore": func(html string) string {
			return utils.UnderScoreString(html)
		},
		"URL": func() string {
			return basePath + r.URL.String()
		},
		"CHART_DATA": func() string {
			return ""
		},
		"Error": func() string {
			return ""
		},
		"Cache": func() Cacher {
			return CacheStorage
		},
		"ToString": func(v interface{}) string {
			return utils.ToString(v)
		},
		"Ago": func(t time.Time) string {
			return utils.Timestamp(t).Ago()
		},
		"Duration": func(t time.Duration) string {
			duration, _ := time.ParseDuration(fmt.Sprintf("%vs", t.Seconds()))
			return utils.FormatDuration(duration)
		},
		"ToUnix": func(t time.Time) int64 {
			return t.UTC().Unix()
		},
		"ParseTime": func(t time.Time, format string) string {
			return t.Format(format)
		},
		"FromUnix": func(t int64) string {
			return utils.Timezoner(time.Unix(t, 0), core.CoreApp.Timezone).Format("Monday, January 02")
		},
		"UnixTime": func(t int64, nano bool) string {
			if nano {
				t = t / 1e9
			}
			return utils.Timezoner(time.Unix(t, 0), core.CoreApp.Timezone).String()
		},
		"ServiceLink": func(s *core.Service) string {
			if s.Permalink.Valid {
				return s.Permalink.String
			}
			return utils.ToString(s.Id)
		},
		"NewService": func() *types.Service {
			return new(types.Service)
		},
		"NewUser": func() *types.User {
			return new(types.User)
		},
		"NewCheckin": func() *types.Checkin {
			return new(types.Checkin)
		},
		"NewMessage": func() *types.Message {
			return new(types.Message)
		},
		"NewGroup": func() *types.Group {
			return new(types.Group)
		},
		"BasePath": func() string {
			return basePath
		},
	}
}
