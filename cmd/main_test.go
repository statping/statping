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
	"github.com/gorilla/sessions"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/handlers"
	"github.com/hunterlong/statup/notifiers"
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
	route            *mux.Router
	testSession      *sessions.Session
	dir              string
	SERVICE_SINCE, _ = time.Parse(time.RFC3339, "2018-08-30T10:42:08-07:00")
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

func RunInit(t *testing.T) {
	source.Assets()
	Clean()
	route = handlers.Router()
	LoadDotEnvs()
	core.CoreApp = core.NewCore()
}

func TestRunAll(t *testing.T) {
	//t.Parallel()

	databases := []string{"postgres", "sqlite", "mysql"}
	if os.Getenv("ONLY_DB") != "" {
		databases = []string{os.Getenv("ONLY_DB")}
	}

	for _, dbt := range databases {
		t.Run(dbt+" init", func(t *testing.T) {
			RunInit(t)
		})
		t.Run(dbt+" Save Config", func(t *testing.T) {
			RunSaveConfig(t, dbt)
		})
		t.Run(dbt+" Load Configs", func(t *testing.T) {
			RunLoadConfig(t)
			t.Log(core.Configs)
		})
		t.Run(dbt+" Connect to Database", func(t *testing.T) {
			err := core.Configs.Connect(false, dir)
			assert.Nil(t, err)
		})
		t.Run(dbt+" Drop Database", func(t *testing.T) {
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
			t.Log(core.CoreApp)
		})
		t.Run(dbt+" Select Services", func(t *testing.T) {
			RunSelectAllMysqlServices(t)
		})
		t.Run(dbt+" Select Comms", func(t *testing.T) {
			t.SkipNow()
			RunSelectAllMysqlCommunications(t)
		})
		t.Run(dbt+" Create Users", func(t *testing.T) {
			RunUser_Create(t)
		})
		t.Run(dbt+" Update User", func(t *testing.T) {
			RunUser_Update(t)
		})
		t.Run(dbt+" Create Non Unique Users", func(t *testing.T) {
			t.SkipNow()
			RunUser_NonUniqueCreate(t)
		})
		t.Run(dbt+" Select Users", func(t *testing.T) {
			RunUser_SelectAll(t)
		})
		t.Run(dbt+" Select Services", func(t *testing.T) {
			RunSelectAllServices(t)
		})
		t.Run(dbt+" Select One Service", func(t *testing.T) {
			RunOneService_Check(t)
		})
		t.Run(dbt+" Create Service", func(t *testing.T) {
			RunService_Create(t)
		})
		t.Run(dbt+" Create Hits", func(t *testing.T) {
			RunCreateService_Hits(t)
		})
		t.Run(dbt+" Service ToJSON()", func(t *testing.T) {
			RunService_ToJSON(t)
		})
		t.Run(dbt+" Avg Time", func(t *testing.T) {
			RunService_AvgTime(t)
		})
		t.Run(dbt+" Online 24h", func(t *testing.T) {
			RunService_Online24(t)
		})
		t.Run(dbt+" Chart Data", func(t *testing.T) {
			RunService_GraphData(t)
		})
		t.Run(dbt+" Create Failing Service", func(t *testing.T) {
			RunBadService_Create(t)
		})
		t.Run(dbt+" Check Bad Service", func(t *testing.T) {
			RunBadService_Check(t)
		})
		t.Run(dbt+" Select Hits", func(t *testing.T) {
			RunService_Hits(t)
		})
		t.Run(dbt+" Select Failures", func(t *testing.T) {
			RunService_Failures(t)
		})
		t.Run(dbt+" Select Limited Hits", func(t *testing.T) {
			RunService_LimitedHits(t)
		})
		t.Run(dbt+" Delete Service", func(t *testing.T) {
			RunDeleteService(t)
		})
		t.Run(dbt+" Delete User", func(t *testing.T) {
			RunUser_Delete(t)
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
			//Clean()
		})

		//<-done

	}

}

func RunSaveConfig(t *testing.T, db string) {
	var err error
	port := 5432
	if db == "mysql" {
		port = 3306
	}
	core.Configs = &core.DbConfig{DbConfig: &types.DbConfig{
		DbConn:      db,
		DbHost:      os.Getenv("DB_HOST"),
		DbUser:      os.Getenv("DB_USER"),
		DbPass:      os.Getenv("DB_PASS"),
		DbData:      os.Getenv("DB_DATABASE"),
		DbPort:      port,
		Project:     "Testing " + db,
		Description: "This is a test of Statup.io!",
		Domain:      "",
		Username:    "admin",
		Password:    "admin",
		Email:       "",
		Error:       nil,
		Location:    dir,
	}}
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
	core.Configs.SeedDatabase()
	//assert.Nil(t, err)
}

func RunLoadConfig(t *testing.T) {
	var err error
	core.Configs, err = core.LoadConfig(dir)
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
	assert.Equal(t, "Awesome Status", core.CoreApp.Name)
	assert.Equal(t, db, core.CoreApp.DbConnection)
	assert.NotEmpty(t, core.CoreApp.ApiKey)
	assert.NotEmpty(t, core.CoreApp.ApiSecret)
	assert.Equal(t, VERSION, core.CoreApp.Version)
}

func RunSelectAllMysqlServices(t *testing.T) {
	var err error
	services, err := core.CoreApp.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 18, len(services))
}

func RunSelectAllMysqlCommunications(t *testing.T) {
	var err error
	notifiers.Collections = core.DbSession.Table("communication").Model(&notifiers.Notification{})
	comms := notifiers.Load()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(comms))
}

func RunUser_SelectAll(t *testing.T) {
	users, err := core.SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func RunUser_Create(t *testing.T) {
	user := core.ReturnUser(&types.User{
		Username: "hunterlong",
		Password: "password123",
		Email:    "info@gmail.com",
		Admin:    true,
	})
	id, err := user.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(2), id)
	user2 := core.ReturnUser(&types.User{
		Username: "superadmin",
		Password: "admin",
		Email:    "info@adminer.com",
		Admin:    true,
	})
	id, err = user2.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(3), id)
}

func RunUser_Update(t *testing.T) {
	user, err := core.SelectUser(1)
	user.Email = "info@updatedemail.com"
	assert.Nil(t, err)
	err = user.Update()
	assert.Nil(t, err)
	updatedUser, err := core.SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "info@updatedemail.com", updatedUser.Email)
}

func RunUser_NonUniqueCreate(t *testing.T) {
	user := core.ReturnUser(&types.User{
		Username: "admin",
		Password: "admin",
		Email:    "info@testuser.com",
	})
	admin, err := user.Create()
	assert.Error(t, err)
	assert.Nil(t, admin)
}

func RunUser_Delete(t *testing.T) {
	user, err := core.SelectUser(2)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	err = user.Delete()
	assert.Nil(t, err)
}

func RunSelectAllServices(t *testing.T) {
	var err error
	services, err := core.CoreApp.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 18, len(services))
	for _, s := range services {
		assert.NotEmpty(t, s.CreatedAt)
	}
}

func RunOneService_Check(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	t.Log(service)
	assert.Equal(t, "Google", service.Name)
}

func RunService_Create(t *testing.T) {
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
	id, err := service.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(19), id)
}

func RunService_ToJSON(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	jsoned := service.ToJSON()
	assert.NotEmpty(t, jsoned)
}

func RunService_AvgTime(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	avg := service.AvgUptime24()
	assert.Equal(t, "100", avg)
}

func RunService_Online24(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	online := service.OnlineSince(SERVICE_SINCE)
	assert.NotEqual(t, float32(0), online)

	service = core.SelectService(6)
	assert.NotNil(t, service)
	online = service.OnlineSince(SERVICE_SINCE)
	assert.Equal(t, float32(0), online)

	service = core.SelectService(18)
	assert.NotNil(t, service)
	online = service.OnlineSince(SERVICE_SINCE)
	assert.Equal(t, float32(0), online)
}

func RunService_GraphData(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	data := service.GraphData()
	t.Log(data)
	assert.NotEqual(t, "null", data)
	assert.False(t, strings.Contains(data, "0001-01-01T00:00:00Z"))
	assert.NotEmpty(t, data)
}

func RunBadService_Create(t *testing.T) {
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
	id, err := service.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(20), id)
}

func RunBadService_Check(t *testing.T) {
	service := core.SelectService(18)
	assert.NotNil(t, service)
	assert.Equal(t, "Failing URL", service.Name)
	for i := 0; i <= 10; i++ {
		service.Check(true)
	}
	assert.True(t, service.IsRunning())
}

func RunDeleteService(t *testing.T) {
	service := core.SelectService(4)
	assert.NotNil(t, service)
	assert.Equal(t, "JSON API Tester", service.Name)
	assert.True(t, service.IsRunning())
	err := service.Delete()
	assert.False(t, service.IsRunning())
	assert.Nil(t, err)
}

func RunCreateService_Hits(t *testing.T) {
	services := core.CoreApp.Services
	assert.NotNil(t, services)
	assert.Equal(t, 37, len(services))
	for _, service := range services {
		service.Check(true)
		assert.NotNil(t, service)
	}
}

func RunService_Hits(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	hits, err := service.Hits()
	assert.Nil(t, err)
	assert.NotZero(t, len(hits))
}

func RunService_Failures(t *testing.T) {
	service := core.SelectService(18)
	assert.NotNil(t, service)
	assert.Equal(t, "Failing URL", service.Name)
	assert.NotEmpty(t, service.AllFailures())
}

func RunService_LimitedHits(t *testing.T) {
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
	assert.True(t, strings.Contains(rr.Body.String(), "Awesome"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunServiceHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Google Service</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunPrometheusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	req.Header.Set("Authorization", core.CoreApp.ApiSecret)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	t.Log(rr.Body.String())
	assert.True(t, strings.Contains(rr.Body.String(), "statup_total_services 37"))
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
	assert.Equal(t, 303, rr.Result().StatusCode)
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
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Users</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
	assert.True(t, handlers.IsAuthenticated(req))
}

func RunUserViewHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Users</title>"))
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
