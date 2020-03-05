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

package services

import (
	"github.com/hunterlong/statping/types/failures"
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
	services := All()
	for _, s := range services {
		s.CheckService(false)
		assert.False(t, s.IsRunning())
		t.Logf("ID: %v %v\n", s.Id, s.Name)
	}
	assert.Equal(t, 15, len(services))
}

func TestServiceDowntime(t *testing.T) {
	service, err := Find(15)
	require.Nil(t, err)
	downtime := service.Downtime()
	assert.True(t, downtime.Seconds() > 0)
}

func TestSelectTCPService(t *testing.T) {
	services := All()
	assert.Equal(t, 15, len(services))
	service, err := Find(5)
	require.Nil(t, err)
	assert.NotNil(t, service)
	assert.Equal(t, "Google DNS", service.Name)
	assert.Equal(t, "tcp", service.Type)
}

func TestUpdateService(t *testing.T) {
	service, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Google", service.Name)
	service.Name = "Updated Google"
	service.Interval = 5

	err = service.Update()
	require.Nil(t, err)

	// check if updating pointer array shutdown any other service
	service, err = Find(1)
	require.Nil(t, err)
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

		err := srv.Update()
		require.Nil(t, err)
		i++
	}
}

func TestServiceHTTPCheck(t *testing.T) {
	service, err := Find(1)
	require.Nil(t, err)
	service.CheckService(true)
	assert.Equal(t, "Changed Updated Google", service.Name)
	assert.True(t, service.Online)
}

func TestCheckHTTPService(t *testing.T) {
	service, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "Changed Updated Google", service.Name)
	assert.True(t, service.Online)
	assert.Equal(t, 200, service.LastStatusCode)
	assert.NotZero(t, service.Latency)
	assert.NotZero(t, service.PingTime)
}

func TestServiceTCPCheck(t *testing.T) {
	service, err := Find(5)
	require.Nil(t, err)
	service.CheckService(false)
	assert.Equal(t, "Changed Google DNS", service.Name)
	assert.True(t, service.Online)
}

func TestCheckTCPService(t *testing.T) {
	service, err := Find(5)
	require.Nil(t, err)
	assert.Equal(t, "Changed Google DNS", service.Name)
	assert.True(t, service.Online)
	assert.NotZero(t, service.Latency)
	assert.NotZero(t, service.PingTime)
}

func TestServiceOnline24Hours(t *testing.T) {
	since := utils.Now().Add(-24 * time.Hour).Add(-10 * time.Minute)
	service, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, float32(100), service.OnlineSince(since))
	service2, err := Find(5)
	require.Nil(t, err)
	assert.Equal(t, float32(100), service2.OnlineSince(since))
	service3, err := Find(14)
	require.Nil(t, err)
	assert.True(t, service3.OnlineSince(since) > float32(49))
}

func TestServiceAvgUptime(t *testing.T) {
	since := utils.Now().Add(-24 * time.Hour).Add(-10 * time.Minute)
	service, err := Find(1)
	require.Nil(t, err)
	assert.NotEqual(t, "0.00", service.AvgTime())
	service2, err := Find(5)
	assert.Equal(t, "100", service2.AvgTime())
	service3, err := Find(13)
	assert.NotEqual(t, "0", service3.AvgUptime(since))
	service4, err := Find(15)
	assert.NotEqual(t, "0", service4.AvgUptime(since))
}

func TestCreateService(t *testing.T) {
	s := &Service{
		Name:           "That'll do 🐢",
		Domain:         "https://www.youtube.com/watch?v=rjQtzV9IZ0Q",
		ExpectedStatus: 200,
		Interval:       3,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		GroupId:        1,
	}
	err := s.Create()
	require.Nil(t, err)
	assert.NotZero(t, s.Id)
	newService, err := Find(s.Id)
	assert.Equal(t, "That'll do 🐢", newService.Name)
}

func TestViewNewService(t *testing.T) {
	newService, err := Find(newServiceId)
	require.Nil(t, err)
	assert.Equal(t, "That'll do 🐢", newService.Name)
}

func TestCreateFailingHTTPService(t *testing.T) {
	s := &Service{
		Name:           "Bad URL",
		Domain:         "http://localhost/iamnothere",
		ExpectedStatus: 200,
		Interval:       2,
		Type:           "http",
		Method:         "GET",
		Timeout:        5,
		GroupId:        1,
	}
	err := s.Create()
	require.Nil(t, err)
	assert.NotZero(t, s.Id)
	newService, err := Find(s.Id)
	require.Nil(t, err)
	assert.Equal(t, "Bad URL", newService.Name)
	t.Log("new service ID: ", newServiceId)
}

func TestServiceFailedCheck(t *testing.T) {
	service, err := Find(17)
	require.Nil(t, err)
	assert.Equal(t, "Bad URL", service.Name)
	service.CheckService(false)
	assert.Equal(t, "Bad URL", service.Name)
	assert.False(t, service.Online)
}

func TestCreateFailingTCPService(t *testing.T) {
	s := &Service{
		Name:     "Bad TCP",
		Domain:   "localhost",
		Port:     5050,
		Interval: 30,
		Type:     "tcp",
		Timeout:  5,
		GroupId:  1,
	}
	err := s.Create()
	assert.Nil(t, err)
	assert.NotZero(t, s.Id)
	newService, err := Find(s.Id)
	require.Nil(t, err)
	assert.Equal(t, "Bad TCP", newService.Name)
	t.Log("new failing tcp service ID: ", newServiceId)
}

func TestServiceFailedTCPCheck(t *testing.T) {
	srv, err := Find(newServiceId)
	require.Nil(t, err)
	srv.CheckService(false)
	assert.Equal(t, "Bad TCP", srv.Name)
	assert.False(t, srv.Online)
}

func TestCreateServiceFailure(t *testing.T) {
	service, err := Find(8)
	fail := &failures.Failure{
		Issue:   "This is not an issue, but it would container HTTP response errors.",
		Method:  "http",
		Service: service.Id,
	}
	err = fail.Create()
	assert.Nil(t, err)
	assert.NotZero(t, fail.Id)
}

func TestDeleteService(t *testing.T) {
	service, err := Find(newServiceId)

	count, err := SelectAllServices(false)
	assert.Nil(t, err)
	assert.Equal(t, 18, len(count))

	err = service.Delete()
	assert.Nil(t, err)

	services := All()
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

func TestFindLink(t *testing.T) {
	service, err := Find(1)
	require.Nil(t, err)
	assert.Equal(t, "google", service.Permalink.String)
}
