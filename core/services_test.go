// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	newServiceId int64
)

func TestSelectHTTPService(t *testing.T) {
	services, err := SelectAllServices(false)
	assert.Nil(t, err)
	assert.Equal(t, 15, len(services))
	assert.Equal(t, "Google", services[0].Name)
	assert.Equal(t, "http", services[0].Type)
}

func TestSelectAllServices(t *testing.T) {
	services := CoreApp.services
	for _, s := range services {
		CheckService(s, false)
		assert.False(t, s.IsRunning())
		t.Logf("ID: %v %v\n", s.Id, s.Name)
	}
	assert.Equal(t, 15, len(services))
}

func TestServiceDowntime(t *testing.T) {
	t.SkipNow()
	service := SelectService(15)
	downtime := service.Downtime()
	assert.True(t, downtime.Seconds() > 0)
}

func TestSelectTCPService(t *testing.T) {
	services := CoreApp.services
	assert.Equal(t, 15, len(services))
	service := SelectService(5)
	assert.NotNil(t, service)
	assert.Equal(t, "Google DNS", service.Name)
	assert.Equal(t, "tcp", service.Type)
}

func TestUpdateService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Google", service.Name)
	service.Name = "Updated Google"
	service.Interval = 5

	err := database.Update(service)
	require.Nil(t, err)

	// check if updating pointer array shutdown any other service
	service = SelectService(1)
	assert.Equal(t, "Updated Google", service.Name)
	assert.Equal(t, 5, service.Interval)
}

func TestUpdateAllServices(t *testing.T) {
	services, err := SelectAllServices(false)
	require.Nil(t, err)
	var i int
	for _, srv := range services {
		srv.Name = "Changed " + srv.Name
		srv.Interval = i + 3

		err := database.Update(srv)
		require.Nil(t, err)
		i++
	}
}

func TestServiceHTTPCheck(t *testing.T) {
	service := SelectService(1)
	CheckService(service, true)
	assert.Equal(t, "Changed Updated Google", service.Name)
	assert.True(t, service.Online)
}

func TestCheckHTTPService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Changed Updated Google", service.Name)
	assert.True(t, service.Online)
	assert.Equal(t, 200, service.LastStatusCode)
	assert.NotZero(t, service.Latency)
	assert.NotZero(t, service.PingTime)
}

func TestServiceTCPCheck(t *testing.T) {
	service := SelectService(5)
	CheckService(service, true)
	assert.Equal(t, "Changed Google DNS", service.Name)
	assert.True(t, service.Online)
}

func TestCheckTCPService(t *testing.T) {
	service := SelectService(5)
	assert.Equal(t, "Changed Google DNS", service.Name)
	assert.True(t, service.Online)
	assert.NotZero(t, service.Latency)
	assert.NotZero(t, service.PingTime)
}

func TestServiceOnline24Hours(t *testing.T) {
	since := utils.Now().Add(-24 * time.Hour).Add(-10 * time.Minute)
	service := SelectService(1)
	assert.Equal(t, float32(100), service.OnlineSince(since))
	service2 := SelectService(5)
	assert.Equal(t, float32(100), service2.OnlineSince(since))
	service3 := SelectService(14)
	assert.True(t, service3.OnlineSince(since) > float32(49))
}

func TestServiceAvgUptime(t *testing.T) {
	since := utils.Now().Add(-24 * time.Hour).Add(-10 * time.Minute)
	service := SelectService(1)
	assert.NotEqual(t, "0.00", service.AvgTime())
	service2 := SelectService(5)
	assert.Equal(t, "100", service2.AvgTime())
	service3 := SelectService(13)
	assert.NotEqual(t, "0", service3.AvgUptime(since))
	service4 := SelectService(15)
	assert.NotEqual(t, "0", service4.AvgUptime(since))
}

func TestCreateService(t *testing.T) {
	s := &types.Service{
		Name:           "That'll do 🐢",
		Domain:         "https://www.youtube.com/watch?v=rjQtzV9IZ0Q",
		ExpectedStatus: 200,
		Interval:       3,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		GroupId:        1,
	}
	obj, err := database.Create(s)
	require.Nil(t, err)
	assert.NotZero(t, obj.Id)
	newService := SelectService(obj.Id)
	assert.Equal(t, "That'll do 🐢", newService.Name)
}

func TestViewNewService(t *testing.T) {
	newService := SelectService(newServiceId)
	assert.Equal(t, "That'll do 🐢", newService.Name)
}

func TestCreateFailingHTTPService(t *testing.T) {
	s := &types.Service{
		Name:           "Bad URL",
		Domain:         "http://localhost/iamnothere",
		ExpectedStatus: 200,
		Interval:       2,
		Type:           "http",
		Method:         "GET",
		Timeout:        5,
		GroupId:        1,
	}
	obj, err := database.Create(s)
	require.Nil(t, err)
	assert.NotZero(t, obj.Id)
	newService := SelectService(obj.Id)
	assert.Equal(t, "Bad URL", newService.Name)
	t.Log("new service ID: ", newServiceId)
}

func TestServiceFailedCheck(t *testing.T) {
	service := SelectService(17)
	assert.Equal(t, "Bad URL", service.Name)
	CheckService(service, false)
	assert.Equal(t, "Bad URL", service.Name)
	assert.False(t, service.Online)
}

func TestCreateFailingTCPService(t *testing.T) {
	s := &types.Service{
		Name:     "Bad TCP",
		Domain:   "localhost",
		Port:     5050,
		Interval: 30,
		Type:     "tcp",
		Timeout:  5,
		GroupId:  1,
	}
	var err error
	obj, err := database.Create(s)
	assert.Nil(t, err)
	assert.NotZero(t, obj.Id)
	newService := SelectService(obj.Id)
	assert.Equal(t, "Bad TCP", newService.Name)
	t.Log("new failing tcp service ID: ", newServiceId)
}

func TestServiceFailedTCPCheck(t *testing.T) {
	service := SelectService(newServiceId)
	CheckService(service, false)
	assert.Equal(t, "Bad TCP", service.Name)
	assert.False(t, service.Online)
}

func TestCreateServiceFailure(t *testing.T) {
	service := SelectService(8)
	fail := &types.Failure{
		Issue:   "This is not an issue, but it would container HTTP response errors.",
		Method:  "http",
		Service: service.Id,
	}
	obj, err := database.Create(fail)
	assert.Nil(t, err)
	assert.NotZero(t, obj.Id)
}

func TestDeleteService(t *testing.T) {
	service := SelectService(newServiceId)

	count, err := SelectAllServices(false)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(count))

	err = service.Delete()
	assert.Nil(t, err)

	services := CoreApp.services
	assert.Equal(t, 17, len(services))
}

func TestServiceCloseRoutine(t *testing.T) {
	s := new(Service)
	s.Name = "example"
	s.Domain = "https://google.com"
	s.Type = "http"
	s.Method = "GET"
	s.ExpectedStatus = 200
	s.Interval = 1
	s.Start()
	assert.True(t, s.IsRunning())
	t.Log(s.Checkpoint)
	t.Log(s.SleepDuration)
	go ServiceCheckQueue(s, false)
	t.Log(s.Checkpoint)
	t.Log(s.SleepDuration)
	time.Sleep(5 * time.Second)
	t.Log(s.Checkpoint)
	t.Log(s.SleepDuration)
	assert.True(t, s.IsRunning())
	s.Close()
	assert.False(t, s.IsRunning())
	s.Close()
	assert.False(t, s.IsRunning())
}

func TestServiceCheckQueue(t *testing.T) {
	s := new(Service)
	s.Name = "example"
	s.Domain = "https://google.com"
	s.Type = "http"
	s.Method = "GET"
	s.ExpectedStatus = 200
	s.Interval = 1
	s.Start()
	assert.True(t, s.IsRunning())
	go ServiceCheckQueue(s, false)

	go func() {
		time.Sleep(5 * time.Second)
		t.Log(s.Checkpoint)
		time.Sleep(6 * time.Second)
	}()

	time.Sleep(5 * time.Second)
	assert.True(t, s.IsRunning())
	s.Close()
	assert.False(t, s.IsRunning())
	s.Close()
	assert.False(t, s.IsRunning())
}

func TestDNScheckService(t *testing.T) {
	s := new(Service)
	s.Name = "example"
	s.Domain = "http://localhost:9000"
	s.Type = "http"
	s.Method = "GET"
	s.ExpectedStatus = 200
	s.Interval = 1
	amount, err := dnsCheck(s)
	assert.Nil(t, err)
	assert.NotZero(t, amount)
}

func TestSelectServiceLink(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "google", service.Permalink.String)
}

func TestGroup_Create(t *testing.T) {
	group := &types.Group{
		Name: "Testing",
	}
	obj, err := database.Create(group)
	assert.Nil(t, err)
	assert.NotZero(t, obj.Id)
}

func TestGroup_Services(t *testing.T) {
	t.SkipNow()
	//group := SelectGroup(1)
	//assert.NotEmpty(t, group.Services())
}

func TestSelectGroups(t *testing.T) {
	groups := SelectGroups(true, false)
	assert.Equal(t, int(3), len(groups))
	groups = SelectGroups(true, true)
	assert.Equal(t, int(5), len(groups))
}
