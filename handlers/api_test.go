package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/pkg/errors"
	_ "github.com/statping/statping/notifiers"
	"github.com/statping/statping/source"
	"github.com/statping/statping/types"
	"github.com/statping/statping/types/core"
	"github.com/statping/statping/types/groups"
	"github.com/statping/statping/types/services"
	"github.com/statping/statping/types/users"
	"github.com/statping/statping/utils"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

var (
	dir string
)

func init() {
	utils.InitLogs()
	source.Assets()
	dir = utils.Directory
	core.New("test")
}

func TestFailedHTTPServer(t *testing.T) {
	var err error
	go func(err error) {
		err = RunHTTPServer()
	}(err)
	go func() {
		time.Sleep(3 * time.Second)
		StopHTTPServer(nil)
	}()
	assert.Nil(t, err)
}

func TestSetupRoutes(t *testing.T) {
	form := url.Values{}
	form.Add("db_host", utils.Params.GetString("DB_HOST"))
	form.Add("db_user", utils.Params.GetString("DB_USER"))
	form.Add("db_password", utils.Params.GetString("DB_PASS"))
	form.Add("db_database", utils.Params.GetString("DB_DATABASE"))
	form.Add("db_connection", utils.Params.GetString("DB_CONN"))
	form.Add("db_port", utils.Params.GetString("DB_PORT"))
	form.Add("project", "Tester")
	form.Add("username", "admin")
	form.Add("password", "password123")
	form.Add("sample_data", "on")
	form.Add("description", "This is an awesome test")
	form.Add("domain", "http://localhost:8080")
	form.Add("email", "info@statping.com")

	badForm := url.Values{}
	badForm.Add("db_host", "badconnection")
	badForm.Add("db_user", utils.Params.GetString("DB_USER"))
	badForm.Add("db_password", utils.Params.GetString("DB_PASS"))
	badForm.Add("db_database", utils.Params.GetString("DB_DATABASE"))
	badForm.Add("db_connection", "mysql")
	badForm.Add("db_port", utils.Params.GetString("DB_PORT"))
	badForm.Add("project", "Tester")
	badForm.Add("username", "admin")
	badForm.Add("password", "password123")
	badForm.Add("sample_data", "on")
	badForm.Add("description", "This is an awesome test")
	badForm.Add("domain", "http://localhost:8080")
	badForm.Add("email", "info@statping.com")

	tests := []HTTPTest{
		{
			Name:           "Statping Check",
			URL:            "/api",
			Method:         "GET",
			ExpectedStatus: 200,
			FuncTest: func(t *testing.T) error {
				if core.App.Setup {
					return errors.New("core has already been setup")
				}
				return nil
			},
		},
		{
			Name:             "Statping Error Setup",
			URL:              "/api/setup",
			Method:           "POST",
			Body:             badForm.Encode(),
			ExpectedStatus:   500,
			ExpectedContains: []string{BadJSONDatabase},
			HttpHeaders:      []string{"Content-Type=application/x-www-form-urlencoded"},
		},
		{
			Name:   "Statping Run Setup",
			URL:    "/api/setup",
			Method: "POST",
			Body:   form.Encode(),
			//ExpectedStatus: 200,
			HttpHeaders:   []string{"Content-Type=application/x-www-form-urlencoded"},
			ExpectedFiles: []string{utils.Directory + "/config.yml"},
			FuncTest: func(t *testing.T) error {
				if !core.App.Setup {
					return errors.New("core has not been setup")
				}
				if len(services.AllInOrder()) == 0 {
					return errors.New("no services where found")
				}
				if len(users.All()) == 0 {
					return errors.New("no users where found")
				}
				if len(groups.All()) == 0 {
					return errors.New("no groups where found")
				}
				return nil
			},
			AfterTest: StopServices,
		},
	}

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
	date := utils.Now().Format("2006-01")
	tests := []HTTPTest{
		{
			Name:             "Statping Details",
			URL:              "/api",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"description":"This is an awesome test"`},
			FuncTest: func(t *testing.T) error {
				if !core.App.Setup {
					return errors.New("database is not setup")
				}
				return nil
			},
		},
		{
			Name:           "Statping Renew API Keys",
			URL:            "/api/renew",
			Method:         "POST",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
		},
		{
			Name:           "Statping View Cache",
			URL:            "/api/cache",
			Method:         "GET",
			ExpectedStatus: 200,
			BeforeTest:     SetTestENV,
			SecureRoute:    true,
			ResponseLen:    0,
		},
		{
			Name:           "Statping Clear Cache",
			URL:            "/api/clear_cache",
			Method:         "POST",
			ExpectedStatus: 200,
			SecureRoute:    true,
			BeforeTest: func(t *testing.T) error {
				CacheStorage.Set("test", []byte("data here"), types.Day)
				list := CacheStorage.List()
				assert.Len(t, list, 1)
				return nil
			},
			AfterTest: func(t *testing.T) error {
				list := CacheStorage.List()
				assert.Len(t, list, 0)
				return nil
			},
		},
		{
			Name:           "Update Core",
			URL:            "/api/core",
			Method:         "POST",
			ExpectedStatus: 200,
			Body: `{
					"name": "Updated Core"
				}`,
			AfterTest: func(t *testing.T) error {
				assert.Equal(t, "Updated Core", core.App.Name)
				return nil
			},
		},
		{
			Name:           "404 Error Page",
			URL:            "/api/missing_404_page",
			Method:         "GET",
			ExpectedStatus: 404,
		},
		{
			Name:             "Health Check endpoint",
			URL:              "/health",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{`"online":true`, `"setup":true`},
		},
		{
			Name:             "Logs endpoint",
			URL:              "/api/logs",
			Method:           "GET",
			ExpectedStatus:   200,
			GreaterThan:      20,
			ExpectedContains: []string{date},
		},
		{
			Name:             "Logs endpoint",
			URL:              "/api/logs",
			Method:           "GET",
			ExpectedStatus:   200,
			GreaterThan:      20,
			ExpectedContains: []string{date},
		},
		{
			Name:             "Logs Last Line endpoint",
			URL:              "/api/logs/last",
			Method:           "GET",
			ExpectedStatus:   200,
			ExpectedContains: []string{date},
		},
		{
			Name:           "Prometheus Export Metrics",
			URL:            "/metrics",
			Method:         "GET",
			BeforeTest:     SetTestENV,
			AfterTest:      UnsetTestENV,
			ExpectedStatus: 200,
			ExpectedContains: []string{
				`go_goroutines`,
				`go_memstats_alloc_bytes`,
				`go_threads`,
			},
		},
	}

	for _, v := range tests {
		t.Run(v.Name, func(t *testing.T) {
			res, t, err := RunHTTPTest(v, t)
			assert.Nil(t, err)
			t.Log(res)
		})
	}
}

type HttpFuncTest func(*testing.T) error

type ResponseFunc func(*httptest.ResponseRecorder, *testing.T, []byte) error

// HTTPTest contains all the parameters for a HTTP Unit Test
type HTTPTest struct {
	Name                string
	URL                 string
	Method              string
	Body                string
	ExpectedStatus      int
	ExpectedContains    []string
	ExpectedNotContains []string
	HttpHeaders         []string
	ExpectedFiles       []string
	FuncTest            HttpFuncTest
	BeforeTest          HttpFuncTest
	AfterTest           HttpFuncTest
	ResponseFunc        ResponseFunc
	ResponseLen         int
	GreaterThan         int
	SecureRoute         bool
	Skip                bool
}

func logTest(t *testing.T, err error) error {
	e := sentry.NewEvent()
	e.Environment = "testing"
	e.Timestamp = utils.Now().Unix()
	e.Message = fmt.Sprintf("failed test %s", t.Name())
	e.Transaction = t.Name()
	sentry.CaptureEvent(e)
	sentry.CaptureException(err)
	return err
}

// RunHTTPTest accepts a HTTPTest type to execute the HTTP request
func RunHTTPTest(test HTTPTest, t *testing.T) (string, *testing.T, error) {
	if test.Skip {
		t.SkipNow()
	}
	if test.BeforeTest != nil {
		if err := test.BeforeTest(t); err != nil {
			return "", t, logTest(t, err)
		}
	}

	rr, err := Request(test)
	if err != nil {
		return "", t, logTest(t, err)
	}
	defer rr.Result().Body.Close()

	if test.ExpectedStatus != 0 {
		if test.ExpectedStatus != rr.Result().StatusCode {
			assert.Equal(t, test.ExpectedStatus, rr.Result().StatusCode)
			return "", t, fmt.Errorf("status code %v does not match %v", rr.Result().StatusCode, test.ExpectedStatus)
		}
	}

	body, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		assert.Nil(t, err)
		return "", t, logTest(t, err)
	}

	stringBody := string(body)

	if len(test.ExpectedContains) != 0 {
		for _, v := range test.ExpectedContains {
			assert.Contains(t, stringBody, v)
		}
	}
	if len(test.ExpectedNotContains) != 0 {
		for _, v := range test.ExpectedNotContains {
			assert.NotContains(t, stringBody, v)
		}
	}
	if len(test.ExpectedFiles) != 0 {
		for _, v := range test.ExpectedFiles {
			assert.FileExists(t, v)
		}
	}
	if test.FuncTest != nil {
		err := test.FuncTest(t)
		assert.Nil(t, err)
	}
	if test.ResponseFunc != nil {
		err := test.ResponseFunc(rr, t, body)
		assert.Nil(t, err)
	}

	if test.ResponseLen != 0 {
		var respArray []interface{}
		err := json.Unmarshal(body, &respArray)
		assert.Nil(t, err)
		assert.Equal(t, test.ResponseLen, len(respArray))
	}

	if test.GreaterThan != 0 {
		var respArray []interface{}
		err := json.Unmarshal(body, &respArray)
		assert.Nil(t, err)
		assert.GreaterOrEqual(t, len(respArray), test.GreaterThan)
	}

	if test.AfterTest != nil {
		if err := test.AfterTest(t); err != nil {
			return "", t, logTest(t, err)
		}
	}
	return stringBody, t, logTest(t, err)
}

func Request(test HTTPTest) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(test.Method, test.URL, strings.NewReader(test.Body))
	if err != nil {
		return nil, err
	}
	if len(test.HttpHeaders) != 0 {
		for _, v := range test.HttpHeaders {
			splits := strings.Split(v, "=")
			req.Header.Set(splits[0], splits[1])
		}
	}
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	return rr, err
}

func SetTestENV(t *testing.T) error {
	utils.Params.Set("GO_ENV", "test")
	return nil
}

func UnsetTestENV(t *testing.T) error {
	utils.Params.Set("GO_ENV", "production")
	return nil
}

func StopServices(t *testing.T) error {
	for _, s := range services.All() {
		s.Close()
	}
	return nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

var (
	Success = `"status":"success"`

	MethodCreate = `"method":"create"`
	MethodUpdate = `"method":"update"`
	MethodDelete = `"method":"delete"`

	BadJSON         = `{incorrect: JSON %%% formatting, [&]}`
	BadJSONResponse = `{"error":"could not decode incoming JSON"}`
	BadJSONDatabase = `{"error":"error connecting to database`
)
