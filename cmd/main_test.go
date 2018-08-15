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
	"github.com/rendon/testcli"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	route       *mux.Router
	testSession *sessions.Session
	dir         string
)

func init() {
	dir = utils.Directory
	os.Remove(dir + "/statup.db")
	//os.Remove(gopath+"/cmd/config.yml")
	os.RemoveAll(dir + "/cmd/assets")
	os.RemoveAll(dir + "/logs")
}

func RunInit(t *testing.T) {
	source.Assets()
	os.Remove(dir + "/statup.db")
	os.Remove(dir + "/cmd/config.yml")
	os.Remove(dir + "/cmd/index.html")
	route = handlers.Router()
	LoadDotEnvs()
	core.CoreApp = core.NewCore()
}

func TestRunAll(t *testing.T) {
	//t.Parallel()

	databases := []string{"sqlite", "postgres", "mysql"}
	if os.Getenv("ONLY_DB") != "" {
		databases = []string{os.Getenv("ONLY_DB")}
	}
	//databases := []string{"sqlite"}

	for _, dbt := range databases {

		t.Run(dbt+" init", func(t *testing.T) {
			RunInit(t)
		})
		t.Run(dbt+" load database config", func(t *testing.T) {
			RunMakeDatabaseConfig(t, dbt)
		})
		t.Run(dbt+" run database migrations", func(t *testing.T) {
			RunDatabaseMigrations(t, dbt)
		})
		t.Run(dbt+" Sample Data", func(t *testing.T) {
			RunInsertMysqlSample(t)
		})
		t.Run(dbt+" Load Configs", func(t *testing.T) {
			RunLoadConfig(t)
		})
		t.Run(dbt+" Select Core", func(t *testing.T) {
			RunSelectCoreMYQL(t, dbt)
		})
		t.Run(dbt+" Select Services", func(t *testing.T) {
			RunSelectAllMysqlServices(t)
		})
		t.Run(dbt+" Select Comms", func(t *testing.T) {
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
		t.Run(dbt+" Check Service", func(t *testing.T) {
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
		//t.Run(dbt+" Cleanup", func(t *testing.T) {
		//	Cleanup(t)
		//})

	}

}

func TestVersionCommand(t *testing.T) {
	c := testcli.Command("statup", "version")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("Statup v"))
}

func TestHelpCommand(t *testing.T) {
	c := testcli.Command("statup", "help")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("statup help               - Shows the user basic information about Statup"))
}

func TestExportCommand(t *testing.T) {
	t.SkipNow()
	c := testcli.Command("statup", "export")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("Exporting Static 'index.html' page"))
	assert.True(t, fileExists(dir+"/cmd/index.html"))
}

func TestAssetsCommand(t *testing.T) {
	c := testcli.Command("statup", "assets")
	c.Run()
	t.Log(c.Stdout())
	t.Log("Directory for Assets: ", dir)
	assert.FileExists(t, dir+"/assets/robots.txt")
	assert.FileExists(t, dir+"/assets/js/main.js")
	assert.FileExists(t, dir+"/assets/scss/base.scss")
}

func RunMakeDatabaseConfig(t *testing.T, db string) {
	port := 5432
	if db == "mysql" {
		port = 3306
	}
	config := &core.DbConfig{
		db,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		port,
		"Testing " + db,
		"This is a test of Statup.io!",
		"",
		"admin",
		"admin",
		"",
		nil,
		dir,
	}
	err := config.Save()
	assert.Nil(t, err)

	_, err = core.LoadConfig()
	assert.Nil(t, err)
	assert.Equal(t, db, core.Configs.Connection)

	err = core.DbConnection(core.Configs.Connection, false, dir)
	assert.Nil(t, err)
}

func RunDatabaseMigrations(t *testing.T, db string) {
	err := core.RunDatabaseUpgrades()
	assert.Nil(t, err)
}

func RunInsertMysqlSample(t *testing.T) {
	err := core.LoadSampleData()
	assert.Nil(t, err)
}

func RunLoadConfig(t *testing.T) {
	var err error
	core.Configs, err = core.LoadConfig()
	assert.Nil(t, err)
	assert.NotNil(t, core.Configs)
}

func RunSelectCoreMYQL(t *testing.T, db string) {
	var err error
	core.CoreApp, err = core.SelectCore()
	assert.Nil(t, err)
	t.Log(core.CoreApp)
	assert.Equal(t, "Testing "+db, core.CoreApp.Name)
	assert.Equal(t, db, core.CoreApp.DbConnection)
	assert.NotEmpty(t, core.CoreApp.ApiKey)
	assert.NotEmpty(t, core.CoreApp.ApiSecret)
	assert.Equal(t, VERSION, core.CoreApp.Version)
}

func RunSelectAllMysqlServices(t *testing.T) {
	var err error
	services, err := core.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(services))
}

func RunSelectAllMysqlCommunications(t *testing.T) {
	var err error
	notifiers.Collections = core.DbSession.Collection("communication")
	comms := notifiers.Load()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(comms))
}

func RunUser_SelectAll(t *testing.T) {
	users, err := core.SelectAllUsers()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func RunUser_Create(t *testing.T) {
	user := &types.User{
		Username: "admin",
		Password: "admin",
		Email:    "info@testuser.com",
		Admin:    true,
	}
	id, err := core.CreateUser(user)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), id)
	user2 := &types.User{
		Username: "superadmin",
		Password: "admin",
		Email:    "info@adminer.com",
		Admin:    true,
	}
	id, err = core.CreateUser(user2)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), id)
}

func RunUser_Update(t *testing.T) {
	user, err := core.SelectUser(1)
	user.Email = "info@updatedemail.com"
	assert.Nil(t, err)
	err = core.UpdateUser(user)
	assert.Nil(t, err)
	updatedUser, err := core.SelectUser(1)
	assert.Nil(t, err)
	assert.Equal(t, "info@updatedemail.com", updatedUser.Email)
}

func RunUser_NonUniqueCreate(t *testing.T) {
	user := &types.User{
		Username: "admin",
		Password: "admin",
		Email:    "info@testuser.com",
	}
	_, err := core.CreateUser(user)
	assert.NotNil(t, err)
}

func RunUser_Delete(t *testing.T) {
	user, err := core.SelectUser(2)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	err = core.DeleteUser(user)
	assert.Nil(t, err)
}

func RunSelectAllServices(t *testing.T) {
	var err error
	services, err := core.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(services))
}

func RunOneService_Check(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	t.Log(service)
	assert.Equal(t, "Google", service.ToService().Name)
}

func RunService_Create(t *testing.T) {
	service := &types.Service{
		Name:           "test service",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       1,
		Port:           0,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
	}
	id, err := core.CreateService(service)
	assert.Nil(t, err)
	assert.Equal(t, int64(6), id)
	t.Log(service)
}

func RunService_AvgTime(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	avg := service.AvgUptime()
	assert.Equal(t, "100", avg)
}

func RunService_Online24(t *testing.T) {
	service := core.SelectService(1)
	assert.NotNil(t, service)
	online := service.Online24()
	assert.Equal(t, float32(100), online)
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
	service := &types.Service{
		Name:           "Bad Service",
		Domain:         "https://9839f83h72gey2g29278hd2od2d.com",
		ExpectedStatus: 200,
		Interval:       10,
		Port:           0,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
	}
	id, err := core.CreateService(service)
	assert.Nil(t, err)
	assert.Equal(t, int64(7), id)
}

func RunBadService_Check(t *testing.T) {
	service := core.SelectService(4)
	assert.NotNil(t, service)
	assert.Equal(t, "JSON API Tester", service.ToService().Name)
}

func RunDeleteService(t *testing.T) {
	service := core.SelectService(4)
	assert.NotNil(t, service)
	assert.Equal(t, "JSON API Tester", service.ToService().Name)
	err := core.DeleteService(service.ToService())
	assert.Nil(t, err)
}

func RunCreateService_Hits(t *testing.T) {
	services, err := core.SelectAllServices()
	assert.Nil(t, err)
	assert.NotNil(t, services)
	for i := 0; i <= 10; i++ {
		for _, s := range services {
			var service *types.Service
			if s.ToService().Type == "http" {
				service = core.ServiceHTTPCheck(s.ToService())
			} else {
				service = core.ServiceTCPCheck(s.ToService())
			}
			assert.NotNil(t, service)
		}
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
	t.SkipNow()
	service := core.SelectService(6)
	assert.NotNil(t, service)
	assert.NotEmpty(t, service.ToService().Failures)
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
	assert.True(t, strings.Contains(rr.Body.String(), "This is a test of Statup.io!"))
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
	assert.True(t, strings.Contains(rr.Body.String(), "statup_total_services 6"))
}

func RunFailingPrometheusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.Equal(t, 401, rr.Result().StatusCode)
}

func RunLoginHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "admin")
	req, err := http.NewRequest("POST", "/dashboard", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Result().StatusCode)
}

func RunDashboardHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Dashboard</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Users</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunUserViewHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Users</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunServicesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/services", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Services</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunHelpHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/help", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Help</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func RunSettingsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Settings</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "Theme Editor"))
	assert.True(t, strings.Contains(rr.Body.String(), "footer"))
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
