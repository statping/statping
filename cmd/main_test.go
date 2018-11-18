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

package main

import (
	"github.com/gorilla/mux"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/core/notifier"
	"github.com/hunterlong/statup/handlers"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	route *mux.Router
	dir   string
)

func init() {
	dir = utils.Directory
}

func Clean() {
	utils.DeleteFile(dir + "/config.yml")
	utils.DeleteFile(dir + "/statup.db")
	utils.DeleteDirectory(dir + "/assets")
	utils.DeleteDirectory(dir + "/logs")
}

func RunInit(db string, t *testing.T) {
	if db == "mssql" {
		os.Setenv("DB_DATABASE", "tempdb")
		os.Setenv("DB_PASS", "PaSsW0rD123")
		os.Setenv("DB_PORT", "1433")
		os.Setenv("DB_USER", "sa")
	}
	source.Assets()
	Clean()
	route = handlers.Router()
	core.CoreApp = core.NewCore()
}

func TestMain(m *testing.M) {
	m.Run()
}

func TestRunAll(t *testing.T) {
	//t.Parallel()

	databases := []string{"postgres", "sqlite", "mysql"}
	if os.Getenv("ONLY_DB") != "" {
		databases = []string{os.Getenv("ONLY_DB")}
	}

	for _, dbt := range databases {
		t.Run(dbt+" init", func(t *testing.T) {
			RunInit(dbt, t)
		})
		t.Run(dbt+" Save Config", func(t *testing.T) {
			RunSaveConfig(t, dbt)
		})
		t.Run(dbt+" Load Configs", func(t *testing.T) {
			RunLoadConfig(t)
		})
		t.Run(dbt+" Connect to Database", func(t *testing.T) {
			err := core.Configs.Connect(false, dir)
			assert.Nil(t, err)
		})
		t.Run(dbt+" Drop Database", func(t *testing.T) {
			assert.NotNil(t, core.Configs)
			RunDropDatabase(t)
		})
		t.Run(dbt+" Connect to Database Again", func(t *testing.T) {
			err := core.Configs.Connect(false, dir)
			assert.Nil(t, err)
		})
		t.Run(dbt+" Inserting Database Structure", func(t *testing.T) {
			RunCreateSchema(t, dbt)
		})
		t.Run(dbt+" Inserting Seed Data", func(t *testing.T) {
			RunInsertSampleData(t)
		})
		t.Run(dbt+" Connect to Database Again", func(t *testing.T) {
			err := core.Configs.Connect(false, dir)
			assert.Nil(t, err)
		})
		t.Run(dbt+" Run Database Migrations", func(t *testing.T) {
			RunDatabaseMigrations(t, dbt)
		})
		t.Run(dbt+" Select Core", func(t *testing.T) {
			RunSelectCoreMYQL(t, dbt)
		})
		t.Run(dbt+" Select Services", func(t *testing.T) {
			RunSelectAllMysqlServices(t)
		})
		t.Run(dbt+" Select Comms", func(t *testing.T) {
			RunSelectAllNotifiers(t)
		})
		t.Run(dbt+" Create Users", func(t *testing.T) {
			RunUserCreate(t)
		})
		t.Run(dbt+" Update user", func(t *testing.T) {
			runUserUpdate(t)
		})
		t.Run(dbt+" Create Non Unique Users", func(t *testing.T) {
			t.SkipNow()
			runUserNonUniqueCreate(t)
		})
		t.Run(dbt+" Select Users", func(t *testing.T) {
			RunUserSelectAll(t)
		})
		t.Run(dbt+" Select Services", func(t *testing.T) {
			RunSelectAllServices(t)
		})
		t.Run(dbt+" Select One Service", func(t *testing.T) {
			RunOneServiceCheck(t)
		})
		t.Run(dbt+" Create Service", func(t *testing.T) {
			RunServiceCreate(t)
		})
		t.Run(dbt+" Create Hits", func(t *testing.T) {
			RunCreateServiceHits(t)
		})
		t.Run(dbt+" Service ToJSON()", func(t *testing.T) {
			RunServiceToJSON(t)
		})
		t.Run(dbt+" Avg Time", func(t *testing.T) {
			runServiceAvgTime(t)
		})
		t.Run(dbt+" Online 24h", func(t *testing.T) {
			RunServiceOnline24(t)
		})
		//t.Run(dbt+" Chart Data", func(t *testing.T) {
		//	RunServiceGraphData(t)
		//})
		t.Run(dbt+" Create Failing Service", func(t *testing.T) {
			RunBadServiceCreate(t)
		})
		t.Run(dbt+" Check Bad Service", func(t *testing.T) {
			RunBadServiceCheck(t)
		})
		t.Run(dbt+" Select Hits", func(t *testing.T) {
			RunServiceHits(t)
		})
		t.Run(dbt+" Select Failures", func(t *testing.T) {
			RunServiceFailures(t)
		})
		t.Run(dbt+" Select Limited Hits", func(t *testing.T) {
			RunServiceLimitedHits(t)
		})
		t.Run(dbt+" Delete Service", func(t *testing.T) {
			RunDeleteService(t)
		})
		t.Run(dbt+" Delete user", func(t *testing.T) {
			RunUserDelete(t)
		})
		t.Run(dbt+" HTTP /", func(t *testing.T) {
			RunIndexHandler(t)
		})
		t.Run(dbt+" HTTP /service/1", func(t *testing.T) {
			RunServiceHandler(t)
		})
		t.Run(dbt+" HTTP /metrics", func(t *testing.T) {
			RunPrometheusHandler(t)
		})
		t.Run(dbt+" HTTP /metrics", func(t *testing.T) {
			RunFailingPrometheusHandler(t)
		})
		t.Run(dbt+" HTTP /login", func(t *testing.T) {
			RunLoginHandler(t)
		})
		t.Run(dbt+" HTTP /dashboard", func(t *testing.T) {
			RunDashboardHandler(t)
		})
		t.Run(dbt+" HTTP /users", func(t *testing.T) {
			RunUsersHandler(t)
		})
		t.Run(dbt+" HTTP /user/1", func(t *testing.T) {
			RunUserViewHandler(t)
		})
		t.Run(dbt+" HTTP /services", func(t *testing.T) {
			RunServicesHandler(t)
		})
		t.Run(dbt+" HTTP /help", func(t *testing.T) {
			RunHelpHandler(t)
		})
		t.Run(dbt+" HTTP /settings", func(t *testing.T) {
			RunSettingsHandler(t)
		})
		t.Run(dbt+" Cleanup", func(t *testing.T) {
			core.Configs.Close()
			core.DbSession = nil
			if dbt == "mssql" {
				os.Setenv("DB_DATABASE", "root")
				os.Setenv("DB_PASS", "password123")
				os.Setenv("DB_PORT", "1433")
			}
			//Clean()
		})

		//<-done

	}

}

func RunSaveConfig(t *testing.T, db string) {
	var err error
	core.Configs = core.EnvToConfig()
	core.Configs.DbConn = db
	core.Configs, err = core.Configs.Save()
	assert.Nil(t, err)
}

func RunCreateSchema(t *testing.T, db string) {
	err := core.Configs.Connect(false, dir)
	assert.Nil(t, err)
	err = core.Configs.CreateDatabase()
	assert.Nil(t, err)
}

func RunDatabaseMigrations(t *testing.T, db string) {
	err := core.Configs.MigrateDatabase()
	assert.Nil(t, err)
}

func RunInsertSampleData(t *testing.T) {
	err := core.InsertLargeSampleData()
	assert.Nil(t, err)
}

func RunLoadConfig(t *testing.T) {
	var err error
	core.Configs, err = core.LoadConfigFile(dir)
	t.Log(core.Configs)
	assert.Nil(t, err)
	assert.NotNil(t, core.Configs)
}

func RunDropDatabase(t *testing.T) {
	err := core.Configs.DropDatabase()
	assert.Nil(t, err)
}

func RunSelectCoreMYQL(t *testing.T, db string) {
	var err error
	core.CoreApp, err = core.SelectCore()
	if err != nil {
		t.FailNow()
	}
	assert.Nil(t, err)
	t.Log("core: ", core.CoreApp.Core)
	assert.Equal(t, "Statup Sample Data", core.CoreApp.Name)
	assert.Equal(t, db, core.CoreApp.DbConnection)
	assert.NotEmpty(t, core.CoreApp.ApiKey)
	assert.NotEmpty(t, core.CoreApp.ApiSecret)
	assert.Equal(t, VERSION, core.CoreApp.Version)
}

func RunSelectAllMysqlServices(t *testing.T) {
	var err error
	services, err := core.CoreApp.SelectAllServices(false)
	assert.Nil(t, err)
	assert.Equal(t, 15, len(services))
}

func RunSelectAllNotifiers(t *testing.T) {
	var err error
	notifier.SetDB(core.DbSession)
	core.CoreApp.Notifications = notifier.Load()
	assert.Nil(t, err)
	assert.Equal(t, 8, len(core.CoreApp.Notifications))
}

func RunUserSelectAll(t *testing.T) {
	users, err := core.SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 4, len(users))
}

func RunUserCreate(t *testing.T) {
	user := core.ReturnUser(&types.User{
		Username: "hunterlong",
		Password: "password123",
		Email:    "info@gmail.com",
		Admin:    types.NewNullBool(true),
	})
	id, err := user.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(3), id)
	user2 := core.ReturnUser(&types.User{
		Username: "superadmin",
		Password: "admin",
		Email:    "info@adminer.com",
		Admin:    types.NewNullBool(true),
	})
	id, err = user2.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(4), id)
}

func runUserUpdate(t *testing.T) {
	user, err := core.SelectUser(1)
	user.Email = "info@updatedemail.com"
	assert.Nil(t, err)
	err = user.Update()
	assert.Nil(t, err)
	updatedUser, err := core.SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "info@updatedemail.com", updatedUser.Email)
}

func runUserNonUniqueCreate(t *testing.T) {
	user := core.ReturnUser(&types.User{
		Username: "admin",
		Password: "admin",
		Email:    "info@testuser.com",
	})
	admin, err := user.Create()
	assert.Error(t, err)
	assert.Nil(t, admin)
}

func RunUserDelete(t *testing.T) {
	user, err := core.SelectUser(2)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	err = user.Delete()
	assert.Nil(t, err)
}

func RunSelectAllServices(t *testing.T) {
	var err error
	services, err := core.CoreApp.SelectAllServices(false)
	assert.Nil(t, err)
	assert.Equal(t, 15, len(services))
	for _, s := range services {
		assert.NotEmpty(t, s.CreatedAt)
	}
}

func RunOneServiceCheck(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	assert.Equal(t, "Google", service.Name)
}

func RunServiceCreate(t *testing.T) {
	service := core.ReturnService(&types.Service{
		Name:           "test service",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       1,
		Port:           0,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
	})
	id, err := service.Create(false)
	assert.Nil(t, err)
	assert.Equal(t, int64(16), id)
}

func RunServiceToJSON(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	jsoned := service.ToJSON()
	assert.NotEmpty(t, jsoned)
}

func runServiceAvgTime(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	avg := service.AvgUptime24()
	assert.Equal(t, "100", avg)
}

func RunServiceOnline24(t *testing.T) {
	var dayAgo = time.Now().Add(-24 * time.Hour).Add(-10 * time.Minute)

	service := core.SelectService(1)
	assert.NotNil(t, service)
	online := service.OnlineSince(dayAgo)
	assert.NotEqual(t, float32(0), online)

	service = core.SelectService(6)
	assert.NotNil(t, service)
	online = service.OnlineSince(dayAgo)
	assert.Equal(t, float32(100), online)

	service = core.SelectService(13)
	assert.NotNil(t, service)
	online = service.OnlineSince(dayAgo)
	assert.True(t, online > 99)

	service = core.SelectService(14)
	assert.NotNil(t, service)
	online = service.OnlineSince(dayAgo)
	assert.True(t, online > float32(49.00))
}

func RunBadServiceCreate(t *testing.T) {
	service := core.ReturnService(&types.Service{
		Name:           "Bad Service",
		Domain:         "https://9839f83h72gey2g29278hd2od2d.com",
		ExpectedStatus: 200,
		Interval:       10,
		Port:           0,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
	})
	id, err := service.Create(false)
	assert.Nil(t, err)
	assert.Equal(t, int64(17), id)
}

func RunBadServiceCheck(t *testing.T) {
	service := core.SelectService(17)
	assert.NotNil(t, service)
	assert.Equal(t, "Bad Service", service.Name)
	for i := 0; i <= 10; i++ {
		service.Check(true)
	}
	assert.True(t, service.IsRunning())
}

func RunDeleteService(t *testing.T) {
	service := core.SelectService(4)
	assert.NotNil(t, service)
	assert.Equal(t, "JSON API Tester", service.Name)
	assert.False(t, service.IsRunning())
	err := service.Delete()
	assert.False(t, service.IsRunning())
	assert.Nil(t, err)
}

func RunCreateServiceHits(t *testing.T) {
	services := core.CoreApp.Services
	assert.NotNil(t, services)
	assert.Equal(t, 16, len(services))
	for _, service := range services {
		service.Check(true)
		assert.NotNil(t, service)
	}
}

func RunServiceHits(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	hits, err := service.Hits()
	assert.Nil(t, err)
	assert.NotZero(t, len(hits))
}

func RunServiceFailures(t *testing.T) {
	service := core.SelectService(17)
	assert.NotNil(t, service)
	assert.Equal(t, "Bad Service", service.Name)
	assert.NotEmpty(t, service.AllFailures())
}

func RunServiceLimitedHits(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	hits, err := service.LimitedHits()
	assert.Nil(t, err)
	assert.NotZero(t, len(hits))
}

func RunIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "Statup"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunServiceHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Google Status</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunPrometheusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	req.Header.Set("Authorization", core.CoreApp.ApiSecret)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	t.Log(rr.Body.String())
	assert.True(t, strings.Contains(rr.Body.String(), "statup_total_services 16"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunFailingPrometheusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Result().StatusCode)
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunLoginHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "password123")
	req, err := http.NewRequest("POST", "/dashboard", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Result().StatusCode)
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunDashboardHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Dashboard</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	t.Log(rr.Body.String())
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Users</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunUserViewHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	t.Log(rr.Body.String())
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | testadmin</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunServicesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/services", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Services</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunHelpHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/help", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Help</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunSettingsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Settings</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "Theme Editor"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
