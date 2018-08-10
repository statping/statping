package handlers

import (
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func init() {
	utils.InitLogs()
	source.Assets()
}

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestSetupHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/setup", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Setup</title>")
}

func TestProcessSetupHandler(t *testing.T) {
	form := url.Values{}
	form.Add("db_host", "")
	form.Add("db_user", "")
	form.Add("db_password", "")
	form.Add("db_database", "")
	form.Add("db_connection", "sqlite")
	form.Add("db_port", "")
	form.Add("project", "Tester")
	form.Add("username", "admin")
	form.Add("password", "password123")
	form.Add("sample_data", "on")
	form.Add("description", "This is an awesome test")
	form.Add("domain", "http://localhost:8080")
	form.Add("email", "info@statup.io")
	req, err := http.NewRequest("POST", "/setup", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestCheckSetupHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/setup", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestCheckIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 200, rr.Code)
}

func TestServicesViewHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Google Service</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestMissingServiceViewHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/99999999", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

func TestServiceChartHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/charts.js", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "var ctx_1")
	assert.Contains(t, body, "var ctx_2")
	assert.Contains(t, body, "var ctx_3")
	assert.Contains(t, body, "var ctx_4")
}

func TestDashboardHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Dashboard</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestLoginHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "password123")
	req, err := http.NewRequest("POST", "/dashboard", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestBadLoginHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "wrongpassword")
	req, err := http.NewRequest("POST", "/dashboard", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Contains(t, body, "Incorrect login information submitted, try again.")
	assert.Equal(t, 200, rr.Code)
}

func TestServicesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/services", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Services</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestCreateUserHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "newuser")
	form.Add("password", "password123")
	form.Add("email", "info@okokk.com")
	form.Add("admin", "on")
	req, err := http.NewRequest("POST", "/users", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestEditUserHandler(t *testing.T) {
	form := url.Values{}
	form.Add("username", "changedusername")
	form.Add("password", "##########")
	form.Add("email", "info@okokk.com")
	form.Add("admin", "on")
	req, err := http.NewRequest("POST", "/user/2", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Contains(t, body, "<td>admin</td>")
	assert.Contains(t, body, "<td>changedusername</td>")
	assert.Equal(t, 200, rr.Code)
}

func TestDeleteUserHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/2/delete", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestUsersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Users</title>")
	assert.Contains(t, body, "<td>admin</td>")
	assert.NotContains(t, body, "<td>changedusername</td>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestUsersEditHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/user/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Users</title>")
	assert.Contains(t, body, "<h3>User admin</h3>")
	assert.Contains(t, body, "value=\"info@statup.io\"")
	assert.Contains(t, body, "value=\"##########\"")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestSettingsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Settings</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestHelpHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/help", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Help</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestCreateHTTPServiceHandler(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Crystal Castles - Kept")
	form.Add("domain", "https://www.youtube.com/watch?v=CfbCLwNlGwU")
	form.Add("method", "GET")
	form.Add("expected_status", "200")
	form.Add("interval", "30")
	form.Add("port", "")
	form.Add("timeout", "30")
	form.Add("check_type", "http")
	form.Add("post_data", "")

	req, err := http.NewRequest("POST", "/services", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestCreateTCPerviceHandler(t *testing.T) {
	form := url.Values{}
	form.Add("name", "Local Postgres")
	form.Add("domain", "localhost")
	form.Add("method", "GET")
	form.Add("expected_status", "")
	form.Add("interval", "30")
	form.Add("port", "5432")
	form.Add("timeout", "30")
	form.Add("check_type", "tcp")
	form.Add("post_data", "")

	req, err := http.NewRequest("POST", "/services", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestServicesHandler2(t *testing.T) {
	req, err := http.NewRequest("GET", "/services", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Services</title>")
	assert.Contains(t, body, "Crystal Castles - Kept")
	assert.Contains(t, body, "Local Postgres")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestViewHTTPServicesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/6", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Crystal Castles - Kept Service</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestViewTCPServicesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/7", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Local Postgres Service</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestServicesDeleteFailuresHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/7/delete_failures", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestServicesUpdateHandler(t *testing.T) {
	form := url.Values{}
	form.Add("name", "The Bravery - An Honest Mistake")
	form.Add("domain", "https://www.youtube.com/watch?v=O8vzbezVru4")
	form.Add("method", "GET")
	form.Add("expected_status", "")
	form.Add("interval", "30")
	form.Add("port", "")
	form.Add("timeout", "15")
	form.Add("check_type", "http")
	form.Add("post_data", "")
	req, err := http.NewRequest("POST", "/service/6", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | The Bravery - An Honest Mistake Service</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestDeleteServiceHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/service/7/delete", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestLogsHandler(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/logs", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Logs</title>")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestLogsLineHandler(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/logs/line", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	t.Log(body)
	assert.NotEmpty(t, body)
}

func TestSaveSettingsHandler(t *testing.T) {
	form := url.Values{}
	form.Add("project", "Awesome Status")
	form.Add("description", "These tests can probably be better")
	req, err := http.NewRequest("POST", "/settings", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestViewSettingsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Settings</title>")
	assert.Contains(t, body, "Awesome Status")
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestSaveAssetsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings/build", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
	assert.FileExists(t, utils.Directory+"/assets/css/base.css")
	assert.FileExists(t, utils.Directory+"/assets/js/main.js")
	assert.DirExists(t, utils.Directory+"/assets")
	assert.True(t, core.UsingAssets)
}

func TestDeleteAssetsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings/delete_assets", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
	assert.False(t, core.UsingAssets)
}

func TestPrometheusHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/metrics", nil)
	req.Header.Set("Authorization", core.CoreApp.ApiSecret)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "statup_total_services 6")
}

func TestSaveNotificationHandler(t *testing.T) {
	form := url.Values{}
	form.Add("enable", "on")
	form.Add("host", "smtp.emailer.com")
	form.Add("port", "587")
	form.Add("username", "exampleuser")
	form.Add("password", "password123")
	form.Add("var1", "info@betatude.com")
	form.Add("var2", "sendto@gmail.com")
	form.Add("api_key", "")
	form.Add("api_secret", "")
	form.Add("limits", "7")
	req, err := http.NewRequest("POST", "/settings/notifier/1", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}

func TestViewNotificationSettingsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/settings", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "<title>Statup | Settings</title>")
	assert.Contains(t, body, `value="exampleuser" id="smtp_username"`)
	assert.Contains(t, body, `value="##########" id="smtp_password"`)
	assert.Contains(t, body, `value="587" id="smtp_port"`)
	assert.Contains(t, body, `value="info@betatude.com" id="outgoing_email_address"`)
	assert.Contains(t, body, `value="sendto@gmail.com" id="send_alerts_to"`)
	assert.Contains(t, body, `value="7" id="limits_per_hour_email"`)
	assert.Contains(t, body, `id="switch-email" checked`)
	assert.Contains(t, body, "Statup  made with ❤️")
}

func TestError404Handler(t *testing.T) {
	req, err := http.NewRequest("GET", "/404me", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 404, rr.Code)
}

func TestLogoutHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/logout", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Code)
}
