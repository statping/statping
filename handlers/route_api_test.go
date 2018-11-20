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

package handlers

import (
	"encoding/json"
	"github.com/hunterlong/statup/core"
	"github.com/hunterlong/statup/source"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const (
	NEW_HTTP_SERVICE = `{"name": "Google Website", "domain": "https://google.com", "expected_status": 200, "check_interval": 10, "type": "http", "method": "GET"}`
)

var (
	dir string
)

func init() {
	dir = utils.Directory
	utils.InitLogs()
	source.Assets()
}

func loadDatabase() {
	core.NewCore()
	core.LoadConfigFile(dir)
	core.Configs = &core.DbConfig{
		DbConn:   "sqlite",
		Location: dir,
	}
	core.CoreApp.DbConnection = "sqlite"
	core.Configs.Connect(false, utils.Directory)
	core.CoreApp.Version = "DEV"
	core.Configs.Save()
}

func createDatabase() {
	core.Configs.DropDatabase()
	core.Configs.CreateDatabase()
}

func resetDatabase() {
	core.Configs.DropDatabase()
	core.Configs.CreateDatabase()
	core.InsertLargeSampleData()
}

func Clean() {
	utils.DeleteFile(dir + "/config.yml")
	utils.DeleteFile(dir + "/statup.db")
	utils.DeleteDirectory(dir + "/assets")
	utils.DeleteDirectory(dir + "/logs")
}

func TestInit(t *testing.T) {
	//Clean()
	//loadDatabase()
	//resetDatabase()
	//loadDatabase()
	//core.SelectCore()
	//core.InsertNotifierDB()
	//core.CoreApp.SelectAllServices(false)
	//core.CoreApp.Notifications = notifier.Load()
}

func formatJSON(res string, out interface{}) {
	json.Unmarshal([]byte(res), &out)
}

func TestApiIndexHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "GET", "/api", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj types.Core
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Awesome Status", obj.Name)
	assert.Equal(t, "sqlite", obj.DbConnection)
}

func TestApiAllServicesHandlerHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "GET", "/api/services", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj []types.Service
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google", obj[0].Name)
	assert.Equal(t, "https://google.com", obj[0].Domain)
}

func TestApiServiceHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "GET", "/api/services/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj types.Service
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google", obj.Name)
	assert.Equal(t, "https://google.com", obj.Domain)
}

func TestApiCreateServiceHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "POST", "/api/services", strings.NewReader(NEW_HTTP_SERVICE))
	assert.Nil(t, err)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	var obj apiResponse
	formatJSON(body, &obj)
	t.Log(body)
	object := obj.Output.(map[string]interface{})
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google Website", object["name"])
	assert.Equal(t, "https://google.com", object["domain"])
}

func TestApiUpdateServiceHandler(t *testing.T) {
	data := `{
    "name": "Updated Service",
    "domain": "https://google.com",
    "expected": "",
    "expected_status": 200,
    "check_interval": 60,
    "type": "http",
    "method": "GET",
    "post_data": "",
    "port": 0,
    "timeout": 10,
    "order_id": 0}`
	rr, err := httpRequestAPI(t, "POST", "/api/services/1", strings.NewReader(data))
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	var obj apiResponse
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "success", obj.Status)
	assert.Equal(t, "update", obj.Method)
	assert.Equal(t, int64(1), obj.Id)
}

func TestApiDeleteServiceHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "DELETE", "/api/services/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj apiResponse
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "delete", obj.Method)
	assert.Equal(t, "success", obj.Status)
}

func TestApiAllUsersHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "GET", "/api/users", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	var obj []types.User
	formatJSON(body, &obj)
	assert.Equal(t, true, obj[0].Admin.Bool)
	assert.Equal(t, "admin", obj[0].Username)
}

func TestApiCreateUserHandler(t *testing.T) {
	data := `{
    "username": "admin2",
    "email": "info@email.com",
    "password": "password123",
    "admin": true}`
	rr, err := httpRequestAPI(t, "POST", "/api/users", strings.NewReader(data))
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj apiResponse
	formatJSON(body, &obj)
	t.Log(body)
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, "create", obj.Method)
	assert.Contains(t, "success", obj.Status)
}

func TestApiViewUserHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "GET", "/api/users/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	var obj types.User
	formatJSON(body, &obj)
	t.Log(body)
	assert.Equal(t, "admin", obj.Username)
	assert.Equal(t, true, obj.Admin.Bool)
}

func TestApiUpdateUserHandler(t *testing.T) {
	data := `{
    "username": "adminupdated",
    "password": "password123",
    "admin": true}`
	rr, err := httpRequestAPI(t, "POST", "/api/users/1", strings.NewReader(data))
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj apiResponse
	formatJSON(body, &obj)
	t.Log(body)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "update", obj.Method)
	assert.Equal(t, "success", obj.Status)
	assert.Equal(t, int64(1), obj.Id)
}

func TestApiDeleteUserHandler(t *testing.T) {
	rr, err := httpRequestAPI(t, "DELETE", "/api/users/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj apiResponse
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "delete", obj.Method)
	assert.Equal(t, "success", obj.Status)
}

func TestApiServiceDataHandler(t *testing.T) {
	grouping := []string{"minute", "hour", "day"}
	for _, g := range grouping {
		params := "?start=0&end=999999999999&group=" + g
		rr, err := httpRequestAPI(t, "GET", "/api/services/5/data"+params, nil)
		assert.Nil(t, err)
		body := rr.Body.String()
		var obj core.DateScanObj
		formatJSON(body, &obj)
		assert.Equal(t, 200, rr.Code)
		assert.NotZero(t, len(obj.Array))
	}
}

func TestApiRenewHandler(t *testing.T) {
	api := core.CoreApp.ApiKey
	secret := core.CoreApp.ApiSecret
	rr, err := httpRequestAPI(t, "GET", "/api/renew", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj apiResponse
	formatJSON(body, &obj)
	assert.Equal(t, 303, rr.Code)
	assert.NotEqual(t, api, core.CoreApp.ApiKey)
	assert.NotEqual(t, secret, core.CoreApp.ApiSecret)
}

func TestApiCheckinHandler(t *testing.T) {
	t.SkipNow()
	service := core.SelectService(1)
	checkin := service.Checkins()
	rr, err := httpRequestAPI(t, "GET", "/api/checkin/"+checkin[0].ApiKey, nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	var obj apiResponse
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
}

func httpRequestAPI(t *testing.T, method, url string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", core.CoreApp.ApiSecret)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Nil(t, err)
	return rr, err
}
