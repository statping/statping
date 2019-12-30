package handlers

import (
	"fmt"
	"github.com/hunterlong/statping/core"
	_ "github.com/hunterlong/statping/notifiers"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

const (
	serverDomain = "http://localhost:8080"
)

var (
	dir string
)

func init() {
	source.Assets()
	utils.InitLogs()
	dir = utils.Directory
}

func TestResetDatabase(t *testing.T) {
	err := core.TmpRecords("handlers.db")
	require.Nil(t, err)
	require.NotNil(t, core.CoreApp)
}

func TestFailedHTTPServer(t *testing.T) {
	err := RunHTTPServer("missinghost", 0)
	assert.Error(t, err)
}

func TestSetupRoutes(t *testing.T) {

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
	form.Add("email", "info@statping.com")

	tests := []HTTPTest{
		{
			Name:           "Statping Setup Check",
			URL:            "/setup",
			Method:         "GET",
			ExpectedStatus: 303,
		},
		{
			Name:           "Statping Run Setup",
			URL:            "/setup",
			Method:         "POST",
			Body:           form.Encode(),
			ExpectedStatus: 303,
			HttpHeaders:    []string{"Content-Type=application/x-www-form-urlencoded"},
			ExpectedFiles:  []string{dir + "/config.yml", dir + "/tmp/" + types.SqliteFilename},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			if err != nil {
				t.FailNow()
			}
		})
	}
}

func TestMainApiRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping Details",
			URL:              "/api",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"name":"Statping Sample Data","description":"This data is only used to testing"`},
		},
		{
			Name:           "Statping Renew API Keys",
			URL:            "/api/renew",
			Method:         "POST",
			ExpectedStatus: 303,
		},
		{
			Name:           "Statping Clear Cache",
			URL:            "/api/clear_cache",
			Method:         "POST",
			ExpectedStatus: 303,
		},
		{
			Name:           "404 Error Page",
			URL:            "/api/missing_404_page",
			Method:         "GET",
			ExpectedStatus: 404,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

func TestApiServiceRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping All Services",
			URL:              "/api/services",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"id":1,"name":"Google","domain":"https://google.com"`},
		},
		{
			Name:             "Statping Service 1",
			URL:              "/api/services/1",
			Method:           "GET",
			ExpectedContains: []string{`"id":1,"name":"Google","domain":"https://google.com"`},
			ExpectedStatus:   200,
		},
		{
			Name:           "Statping Service 1 Data",
			URL:            "/api/services/1/data",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Ping Data",
			URL:            "/api/services/1/ping",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Heatmap Data",
			URL:            "/api/services/1/heatmap",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Hits",
			URL:            "/api/services/1/hits",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Service 1 Failures",
			URL:            "/api/services/1/failures",
			Method:         "GET",
			Body:           "",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Reorder Services",
			URL:            "/api/services/reorder",
			Method:         "POST",
			Body:           `[{"service":1,"order":1},{"service":5,"order":2},{"service":2,"order":3},{"service":3,"order":4},{"service":4,"order":5}]`,
			ExpectedStatus: 200,
			HttpHeaders:    []string{"Content-Type=application/json"},
		},
		{
			Name:        "Statping Create Service",
			URL:         "/api/services",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
    "name": "New Service",
    "domain": "https://statping.com",
    "expected": "",
    "expected_status": 200,
    "check_interval": 30,
    "type": "http",
    "method": "GET",
    "post_data": "",
    "port": 0,
    "timeout": 30,
    "order_id": 0
}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"service","method":"create"`},
		},
		{
			Name:        "Statping Update Service",
			URL:         "/api/services/1",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
    "name": "Updated New Service",
    "domain": "https://google.com",
    "expected": "",
    "expected_status": 200,
    "check_interval": 60,
    "type": "http",
    "method": "GET",
    "post_data": "",
    "port": 0,
    "timeout": 10,
    "order_id": 0
}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"service","method":"update"`},
		},
		{
			Name:             "Statping Delete Service",
			URL:              "/api/services/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"service","method":"delete"`},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

func TestGroupAPIRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Groups",
			URL:            "/api/groups",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping View Group",
			URL:            "/api/groups/1",
			Method:         "GET",
			ExpectedStatus: 200,
		},
		{
			Name:        "Statping Create Group",
			URL:         "/api/groups",
			HttpHeaders: []string{"Content-Type=application/json"},
			Body: `{
    "name": "New Group",
    "public": true
}`,
			Method:         "POST",
			ExpectedStatus: 200,
		},
		{
			Name:           "Statping Delete Group",
			URL:            "/api/groups/1",
			Method:         "DELETE",
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

func TestApiUsersRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping All Users",
			URL:            "/api/users",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:        "Statping Create User",
			URL:         "/api/users",
			HttpHeaders: []string{"Content-Type=application/json"},
			Method:      "POST",
			Body: `{
    "username": "adminuser2",
    "email": "info@adminemail.com",
    "password": "passsword123",
    "admin": true
}`,
			ExpectedStatus: 200,
		}, {
			Name:           "Statping View User",
			URL:            "/api/users/1",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:   "Statping Update User",
			URL:    "/api/users/1",
			Method: "POST",
			Body: `{
    "username": "adminupdated",
    "email": "info@email.com",
    "password": "password12345",
    "admin": true
}`,
			ExpectedStatus: 200,
		}, {
			Name:           "Statping Delete User",
			URL:            "/api/users/1",
			Method:         "DELETE",
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

func TestApiNotifiersRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Notifiers",
			URL:            "/api/notifiers",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:           "Statping Mobile Notifier",
			URL:            "/api/notifier/mobile",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:   "Statping Update Notifier",
			URL:    "/api/notifier/mobile",
			Method: "POST",
			Body: `{
    "method": "mobile",
    "var1": "ExponentPushToken[ToBadIWillError123456]",
    "enabled": true,
    "limits": 55
}`,
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

func TestMessagesApiRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:             "Statping Messages",
			URL:              "/api/messages",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"id":1,"title":"Routine Downtime"`},
		}, {
			Name:   "Statping Create Message",
			URL:    "/api/messages",
			Method: "POST",
			Body: `{
    "title": "API Message",
    "description": "This is an example a upcoming message for a service!",
    "start_on": "2022-11-17T03:28:16.323797-08:00",
    "end_on": "2022-11-17T05:13:16.323798-08:00",
    "service": 1,
    "notify_users": true,
    "notify_method": "email",
    "notify_before": 6,
    "notify_before_scale": "hour"
}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"message","method":"create"`},
		},
		{
			Name:             "Statping View Message",
			URL:              "/api/messages/1",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"id":1,"title":"Routine Downtime"`},
		}, {
			Name:   "Statping Update Message",
			URL:    "/api/messages/1",
			Method: "POST",
			Body: `{
    "title": "Updated Message",
    "description": "This message was updated",
    "start_on": "2022-11-17T03:28:16.323797-08:00",
    "end_on": "2022-11-17T05:13:16.323798-08:00",
    "service": 1,
    "notify_users": true,
    "notify_method": "email",
    "notify_before": 3,
    "notify_before_scale": "hour"
}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"message","method":"update"`},
		},
		{
			Name:             "Statping Delete Message",
			URL:              "/api/messages/1",
			Method:           "DELETE",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"message","method":"delete"`},
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

func TestApiCheckinRoutes(t *testing.T) {
	tests := []HTTPTest{
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
		}, {
			Name:   "Statping Create Checkin",
			URL:    "/api/checkin",
			Method: "POST",
			Body: `{
    "service_id": 2,
    "name": "Server Checkin",
    "interval": 900,
    "grace": 60
}`,
			ExpectedStatus:   200,
			ExpectedContains: []string{`"status":"success","type":"checkin","method":"create"`},
		},
		{
			Name:           "Statping Checkins",
			URL:            "/api/checkins",
			Method:         "GET",
			ExpectedStatus: 200,
		}}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			_, t, err := RunHTTPTest(v, t)
			require.Nil(t, err)
		})
	}
}

// HTTPTest contains all the parameters for a HTTP Unit Test
type HTTPTest struct {
	Name             string
	URL              string
	Method           string
	Body             string
	ExpectedStatus   int
	ExpectedContains []string
	HttpHeaders      []string
	ExpectedFiles    []string
}

// RunHTTPTest accepts a HTTPTest type to execute the HTTP request
func RunHTTPTest(test HTTPTest, t *testing.T) (string, *testing.T, error) {
	req, err := http.NewRequest(test.Method, serverDomain+test.URL, strings.NewReader(test.Body))
	if err != nil {
		assert.Nil(t, err)
		return "", t, err
	}
	if len(test.HttpHeaders) != 0 {
		for _, v := range test.HttpHeaders {
			splits := strings.Split(v, "=")
			req.Header.Set(splits[0], splits[1])
		}
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	body, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		assert.Nil(t, err)
		return "", t, err
	}
	stringBody := string(body)
	if test.ExpectedStatus != rr.Result().StatusCode {
		assert.Equal(t, test.ExpectedStatus, rr.Result().StatusCode)
		return stringBody, t, fmt.Errorf("status code %v does not match %v", rr.Result().StatusCode, test.ExpectedStatus)
	}
	if len(test.ExpectedContains) != 0 {
		for _, v := range test.ExpectedContains {
			assert.Contains(t, stringBody, v)
		}
	}
	if len(test.ExpectedFiles) != 0 {
		for _, v := range test.ExpectedFiles {
			assert.FileExists(t, v)
		}
	}
	return stringBody, t, err
}
