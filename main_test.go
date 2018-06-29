package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
)

func init() {
	route = Router()
}

func TestInit(t *testing.T) {
	RenderBoxes()
	os.Remove("./statup.db")
	Router()
	LoadDotEnvs()

}

func TestMySQLMakeConfig(t *testing.T) {
	config := &DbConfig{
		"mysql",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		3306,
		"Testing MYSQL",
		"This is a test of Statup.io!",
		"",
		"admin",
		"admin",
		"",
		nil,
	}
	err := config.Save()
	assert.Nil(t, err)

	_, err = LoadConfig()
	assert.Nil(t, err)
	assert.Equal(t, "mysql", configs.Connection)

	err = DbConnection(configs.Connection)
	assert.Nil(t, err)
	InsertDefaultComms()
}

func TestInsertMysqlSample(t *testing.T) {
	err := LoadSampleData()
	assert.Nil(t, err)
}

func TestSelectCoreMYQL(t *testing.T) {
	var err error
	core, err = SelectCore()
	assert.Nil(t, err)
	assert.Equal(t, "Testing MYSQL", core.Name)
	assert.Equal(t, VERSION, core.Version)
}

func TestSqliteMakeConfig(t *testing.T) {
	config := &DbConfig{
		"sqlite",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		5432,
		"Testing SQLITE",
		"This is a test of Statup.io!",
		"",
		"admin",
		"admin",
		"",
		nil,
	}
	err := config.Save()
	assert.Nil(t, err)

	_, err = LoadConfig()
	assert.Nil(t, err)
	assert.Equal(t, "sqlite", configs.Connection)

	err = DbConnection(configs.Connection)
	assert.Nil(t, err)
	InsertDefaultComms()
}

func TestInsertSqliteSample(t *testing.T) {
	err := LoadSampleData()
	assert.Nil(t, err)
}

func TestPostgresMakeConfig(t *testing.T) {
	config := &DbConfig{
		"postgres",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		5432,
		"Testing POSTGRES",
		"This is a test of Statup.io!",
		"",
		"admin",
		"admin",
		"",
		nil,
	}
	err := config.Save()
	assert.Nil(t, err)

	_, err = LoadConfig()
	assert.Nil(t, err)
	assert.Equal(t, "postgres", configs.Connection)

	err = DbConnection(configs.Connection)
	assert.Nil(t, err)
	InsertDefaultComms()
}

func TestInsertPostgresSample(t *testing.T) {
	err := LoadSampleData()
	assert.Nil(t, err)
}

func TestSelectCorePostgres(t *testing.T) {
	var err error
	core, err = SelectCore()
	assert.Nil(t, err)
	assert.Equal(t, "Testing POSTGRES", core.Name)
	assert.Equal(t, VERSION, core.Version)
}

func TestSelectCore(t *testing.T) {
	var err error
	core, err = SelectCore()
	assert.Nil(t, err)
	assert.Equal(t, "Testing POSTGRES", core.Name)
	assert.Equal(t, VERSION, core.Version)
}

func TestUser_Create(t *testing.T) {
	user := &User{
		Username: "admin",
		Password: "admin",
		Email:    "info@testuser.com",
	}
	id, err := user.Create()
	assert.Nil(t, err)
	assert.NotZero(t, id)
}

func TestOneService_Check(t *testing.T) {
	service := SelectService(1)
	assert.NotNil(t, service)
	assert.Equal(t, "Google", service.Name)
}

func TestService_Create(t *testing.T) {
	service := &Service{
		Name:           "test service",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       1,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	id, err := service.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(5), id)
}

func TestService_Check(t *testing.T) {
	service := SelectService(2)
	assert.NotNil(t, service)
	assert.Equal(t, "Statup.io", service.Name)
	out := service.Check()
	assert.Equal(t, true, out.Online)
}

func TestService_AvgTime(t *testing.T) {
	service := SelectService(1)
	assert.NotNil(t, service)
	avg := service.AvgUptime()
	assert.Equal(t, "100", avg)
}

func TestService_Online24(t *testing.T) {
	service := SelectService(1)
	assert.NotNil(t, service)
	online := service.Online24()
	assert.Equal(t, float32(100), online)
}

func TestService_GraphData(t *testing.T) {
	service := SelectService(1)
	assert.NotNil(t, service)
	data := service.GraphData()
	assert.NotEmpty(t, data)
}

func TestBadService_Create(t *testing.T) {
	service := &Service{
		Name:           "bad service",
		Domain:         "https://9839f83h72gey2g29278hd2od2d.com",
		ExpectedStatus: 200,
		Interval:       10,
		Port:           0,
		Type:           "https",
		Method:         "GET",
	}
	id, err := service.Create()
	assert.Nil(t, err)
	assert.Equal(t, int64(6), id)
}

func TestBadService_Check(t *testing.T) {
	service := SelectService(4)
	assert.NotNil(t, service)
	assert.Equal(t, "Github Failing Check", service.Name)
}

func TestService_Hits(t *testing.T) {
	service := SelectService(1)
	assert.NotNil(t, service)
	hits, err := service.Hits()
	assert.Nil(t, err)
	assert.NotZero(t, len(hits))
}

func TestService_LimitedHits(t *testing.T) {
	service := SelectService(1)
	assert.NotNil(t, service)
	hits, err := service.LimitedHits()
	assert.Nil(t, err)
	assert.NotZero(t, len(hits))
}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "This is a test of Statup.io!"))
}

func TestServiceHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Services</title>"))
}

func TestPrometheusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	t.Log(rr.Body.String())
	assert.True(t, strings.Contains(rr.Body.String(), "statup_total_services 14"))
}

func TestLoginHandler(t *testing.T) {
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

func TestDashboardHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Dashboard</title>"))
}

func TestUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Users</title>"))
}

func TestServicesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/services", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Services</title>"))
}

func TestHelpHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/help", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Help</title>"))
}

func TestSettingsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	route.ServeHTTP(rr, req)
	assert.True(t, strings.Contains(rr.Body.String(), "<title>Statup | Settings</title>"))
	assert.True(t, strings.Contains(rr.Body.String(), "Theme Editor"))
	assert.True(t, strings.Contains(rr.Body.String(), "Email Settings"))
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
	c := testcli.Command("statup", "export")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("Exported Statup index page"))
}

func TestAssetsCommand(t *testing.T) {
	t.SkipNow()
	c := testcli.Command("statup", "assets")
	c.Run()
	t.Log(c.Stdout())
	assert.True(t, c.StdoutContains("Statup v"))
}
