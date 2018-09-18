// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// +build debug

// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//
//  Debug instance of Statup using pprof and debugcharts
//
//  go get -u github.com/google/pprof
//  go get -v -u github.com/mkevac/debugcharts
//
//  debugcharts web interface is on http://localhost:9090
//
//  - pprof -http=localhost:6060 http://localhost:8080/debug/pprof/profile
//  - pprof -http=localhost:6060 http://localhost:8080/debug/pprof/heap
//  - pprof -http=localhost:6060 http://localhost:8080/debug/pprof/goroutine
//  - pprof -http=localhost:6060 http://localhost:8080/debug/pprof/block
//

package main

import (
	"fmt"
	gorillahandler "github.com/gorilla/handlers"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/handlers"
	_ "github.com/mkevac/debugcharts"
	"net/http"
	"net/http/pprof"
	"os"
	"time"
)

func init() {
	os.Setenv("GO_ENV", "test")
	go func() {
		time.Sleep(5 * time.Second)
		r := handlers.ReturnRouter()
		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)
		r.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		r.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		r.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		r.Handle("/debug/pprof/block", pprof.Handler("block"))
		handlers.UpdateRouter(r)
		time.Sleep(5 * time.Second)
		go ViewPagesLoop()
	}()
	go func() {
		panic(http.ListenAndServe(":9090", gorillahandler.CompressHandler(http.DefaultServeMux)))
	}()
}

func ViewPagesLoop() {
	httpRequest("/")
	httpRequest("/charts.js")
	httpRequest("/css/base.css")
	httpRequest("/css/bootstrap.min.css")
	httpRequest("/js/main.js")
	httpRequest("/js/jquery-3.3.1.min.js")
	httpRequest("/login")
	httpRequest("/dashboard")
	httpRequest("/settings")
	httpRequest("/users")
	httpRequest("/users/1")
	httpRequest("/services")
	httpRequest("/help")
	httpRequest("/logs")
	httpRequest("/404pageishere")
	for i := 1; i <= len(core.CoreApp.Services()); i++ {
		httpRequest(fmt.Sprintf("/service/%v", i))
	}
	defer ViewPagesLoop()
}

func httpRequest(url string) {
	domain := fmt.Sprintf("http://localhost:%v%v", port, url)
	response, err := http.Get(domain)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	defer response.Body.Close()
	time.Sleep(10 * time.Millisecond)
}
