package handlers

import (
	"encoding/json"
	"fmt"
	_ "github.com/hunterlong/statping/notifiers"
	"github.com/hunterlong/statping/source"
	"github.com/hunterlong/statping/types/core"
	"github.com/hunterlong/statping/utils"
	"github.com/pkg/errors"
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
			FuncTest: func() error {
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
			FuncTest: func() error {
				if !core.App.Setup {
					return errors.New("core has not been setup")
				}
				return nil
			},
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
			FuncTest: func() error {
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
		},
		{
			Name:           "Statping Clear Cache",
			URL:            "/api/clear_cache",
			Method:         "POST",
			ExpectedStatus: 200,
			FuncTest: func() error {
				if len(CacheStorage.List()) != 0 {
					return errors.New("cache was not reset")
				}
				return nil
			},
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

type HttpFuncTest func() error

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
	FuncTest         HttpFuncTest
	ResponseLen      int
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
	defer rr.Result().Body.Close()

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
	if test.FuncTest != nil {
		err := test.FuncTest()
		assert.Nil(t, err)
	}
	if test.ResponseLen != 0 {
		var respArray []interface{}
		err := json.Unmarshal(body, &respArray)
		assert.Nil(t, err)
		assert.Equal(t, test.ResponseLen, len(respArray))
	}
	return stringBody, t, err
}
