package handlers

import (
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
	"os"
	"strings"
	"testing"
)

const (
	serverDomain = "http://localhost:18888"
)

var (
	dir string
)

func init() {
	source.Assets()
	utils.InitLogs()
	dir = utils.Directory
	core.New("test")
}

//func TestResetDatabase(t *testing.T) {
//	err := core.TmpRecords("handlers.db")
//	t.Log(err)
//	require.Nil(t, err)
//	require.NotNil(t, core.CoreApp)
//}

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
			Name:           "Statping Run Setup",
			URL:            "/api/setup",
			Method:         "POST",
			Body:           form.Encode(),
			ExpectedStatus: 200,
			HttpHeaders:    []string{"Content-Type=application/x-www-form-urlencoded"},
			ExpectedFiles:  []string{dir + "/config.yml", dir + "/" + "statping.db"},
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
			ExpectedContains: []string{`"description":"This data is only used to testing"`},
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
		//{
		//	Name:           "Prometheus Export Metrics",
		//	URL:            "/metrics",
		//	Method:         "GET",
		//	ExpectedStatus: 200,
		//	ExpectedContains: []string{
		//		`Statping Totals`,
		//		`total_failures`,
		//		`Golang Metrics`,
		//	},
		//}
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
	ResponseLen         int
	SecureRoute         bool
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

	body, err := ioutil.ReadAll(rr.Result().Body)
	if err != nil {
		assert.Nil(t, err)
		return "", t, logTest(t, err)
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
	if test.ResponseLen != 0 {
		var respArray []interface{}
		err := json.Unmarshal(body, &respArray)
		assert.Nil(t, err)
		assert.Equal(t, test.ResponseLen, len(respArray))
	}
	//if test.SecureRoute {
	//	UnsetTestENV()
	//	rec, err := Request(test)
	//	if err != nil {
	//		return "", t, logTest(t, err)
	//	}
	//	defer rec.Result().Body.Close()
	//	assert.Equal(t, http.StatusUnauthorized, rec.Result().StatusCode)
	//}

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
	return os.Setenv("GO_ENV", "test")
}

func UnsetTestENV(t *testing.T) error {
	return os.Setenv("GO_ENV", "production")
}

func StopServices(t *testing.T) error {
	for _, s := range services.All() {
		s.Close()
	}
	return nil
}
