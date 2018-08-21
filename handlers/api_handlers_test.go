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
	NEW_HTTP_SERVICE     = `{"name": "Google Website", "domain": "https://google.com", "expected_status": 200, "check_interval": 10, "type": "http", "method": "GET"}`
	UPDATED_HTTP_SERVICE = `{"id": 1, name": "Google Website", "domain": "https://google.com", "expected_status": 200, "check_interval": 10, "type": "http", "method": "GET"}`
	NEW_TCP_SERVICE      = `{"name": "Google DNS", "domain": "8.8.8.8", "expected": "", "check_interval": 5, "type": "tcp"}`
)

func injectDatabase() {
	core.NewCore()
	core.Configs = new(types.Config)
	core.Configs.Connection = "sqlite"
	core.CoreApp.DbConnection = "sqlite"
	core.CoreApp.Version = "DEV"
	core.DbConnection("sqlite", false, utils.Directory)
	core.InitApp()
}

func TestInit(t *testing.T) {
	t.SkipNow()
	injectDatabase()
}

func formatJSON(res string, out interface{}) {
	json.Unmarshal([]byte(res), &out)
}

func TestApiIndexHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "GET", "/api", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	var obj types.Core
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Tester", obj.Name)
	assert.Equal(t, "sqlite", obj.DbConnection)
}

func TestApiAllServicesHandlerHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "GET", "/api/services", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	var obj []types.Service
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google", obj[0].Name)
	assert.Equal(t, "https://google.com", obj[0].Domain)
}

func TestApiServiceHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "GET", "/api/services/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	var obj types.Service
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google", obj.Name)
	assert.Equal(t, "https://google.com", obj.Domain)
}

func TestApiCreateServiceHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "POST", "/api/services", strings.NewReader(NEW_HTTP_SERVICE))
	assert.Nil(t, err)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	t.Log(body)
	var obj types.Service
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google Website", obj.Name)
	assert.Equal(t, "https://google.com", obj.Domain)
}

func TestApiUpdateServiceHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "POST", "/api/services/1", strings.NewReader(UPDATED_HTTP_SERVICE))
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	var obj types.Service
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google Website", obj.Name)
	assert.Equal(t, "https://google.com", obj.Domain)
}

func TestApiDeleteServiceHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "DELETE", "/api/services/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	var obj ApiResponse
	formatJSON(body, &obj)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, "Google Website", obj.Method)
	assert.Equal(t, "https://google.com", obj.Status)
}

func TestApiAllUsersHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "GET", "/api/users", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	assert.Equal(t, 200, rr.Code)
	var obj []types.User
	formatJSON(body, &obj)
	assert.Equal(t, "Google", obj[0].Admin)
	assert.Equal(t, "https://google.com", obj[0].Username)
}

func TestApiCreateUserHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "POST", "/api/users", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "statup_total_services 6")
}

func TestApiViewUserHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "GET", "/api/users/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "statup_total_services 6")
}

func TestApiUpdateUserHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "POST", "/api/users/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	t.Log(body)
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "statup_total_services 6")
}

func TestApiDeleteUserHandler(t *testing.T) {
	t.SkipNow()
	rr, err := httpRequestAPI(t, "DELETE", "/api/users/1", nil)
	assert.Nil(t, err)
	body := rr.Body.String()
	assert.Equal(t, 200, rr.Code)
	assert.Contains(t, body, "statup_total_services 6")
}

func httpRequestAPI(t *testing.T, method, url string, body io.Reader) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", core.CoreApp.ApiSecret)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Nil(t, err)
	return rr, err
}
