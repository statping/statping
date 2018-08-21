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

package core

import (
	"github.com/hunterlong/statup/types"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	newServiceId int64
)

func TestSelectAllServices(t *testing.T) {
	services, err := CoreApp.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(services))
}

func TestSelectHTTPService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Google", service.Name)
	assert.Equal(t, "http", service.Type)
}

func TestSelectTCPService(t *testing.T) {
	service := SelectService(5)
	assert.Equal(t, "Google DNS", service.Name)
	assert.Equal(t, "tcp", service.Type)
}

func TestUpdateService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Google", service.Name)
	srv := service
	srv.Name = "Updated Google"
	err := srv.Update()
	assert.Nil(t, err)
}

func TestUpdateAllServices(t *testing.T) {
	services, err := CoreApp.SelectAllServices()
	assert.Nil(t, err)
	for k, s := range services {
		srv := ReturnService(s)
		srv.Name = "Changed " + srv.Name
		srv.Interval = k + 3
		err := srv.Update()
		assert.Nil(t, err)
	}
}

func TestServiceHTTPCheck(t *testing.T) {
	service := SelectService(1)
	checked := service.Check(true)
	assert.Equal(t, "Changed Updated Google", checked.Name)
	assert.True(t, checked.Online)
}

func TestCheckHTTPService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Changed Updated Google", service.Name)
	assert.True(t, service.Online)
	assert.Equal(t, 200, service.LastStatusCode)
	assert.NotZero(t, service.Latency)
}

func TestServiceTCPCheck(t *testing.T) {
	service := SelectService(5)
	checked := service.Check(true)
	assert.Equal(t, "Changed Google DNS", checked.Name)
	assert.True(t, checked.Online)
}

func TestCheckTCPService(t *testing.T) {
	service := SelectService(5)
	assert.Equal(t, "Changed Google DNS", service.Name)
	assert.True(t, service.Online)
	assert.NotZero(t, service.Latency)
}

func TestServiceOnline24Hours(t *testing.T) {
	service := SelectService(5)
	amount := service.Online24()
	assert.Equal(t, float32(100), amount)
}

func TestServiceSmallText(t *testing.T) {
	service := SelectService(5)
	text := service.SmallText()
	assert.Contains(t, text, "Online since")
}

func TestServiceAvgUptime(t *testing.T) {
	service := SelectService(5)
	uptime := service.AvgUptime()
	assert.Equal(t, "100", uptime)
}

func TestServiceHits(t *testing.T) {
	service := SelectService(5)
	hits, err := service.Hits()
	assert.Nil(t, err)
	assert.Equal(t, int(1), len(hits))
}

func TestServiceLimitedHits(t *testing.T) {
	service := SelectService(5)
	hits, err := service.LimitedHits()
	assert.Nil(t, err)
	assert.Equal(t, int(1), len(hits))
}

func TestServiceTotalHits(t *testing.T) {
	service := SelectService(5)
	hits, err := service.TotalHits()
	assert.Nil(t, err)
	assert.Equal(t, uint64(0x1), hits)
}

func TestServiceSum(t *testing.T) {
	service := SelectService(5)
	sum, err := service.Sum()
	assert.Nil(t, err)
	assert.NotZero(t, sum)
}

func TestCountOnline(t *testing.T) {
	amount := CountOnline()
	assert.Equal(t, 2, amount)
}

func TestCreateService(t *testing.T) {
	s := ReturnService(&types.Service{
		Name:           "That'll do 🐢",
		Domain:         "https://www.youtube.com/watch?v=rjQtzV9IZ0Q",
		ExpectedStatus: 200,
		Interval:       3,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
	})
	var err error
	newServiceId, err = s.Create()
	assert.Nil(t, err)
	assert.NotZero(t, newServiceId)
	newService := SelectService(newServiceId)
	assert.Equal(t, "That'll do 🐢", newService.Name)
}

func TestViewNewService(t *testing.T) {
	newService := SelectService(newServiceId)
	assert.Equal(t, "That'll do 🐢", newService.Name)

}

func TestCreateFailingHTTPService(t *testing.T) {
	s := ReturnService(&types.Service{
		Name:           "Bad URL",
		Domain:         "http://localhost/iamnothere",
		ExpectedStatus: 200,
		Interval:       2,
		Type:           "http",
		Method:         "GET",
		Timeout:        5,
	})
	var err error
	newServiceId, err = s.Create()
	assert.Nil(t, err)
	assert.NotZero(t, newServiceId)
	newService := SelectService(newServiceId)
	assert.Equal(t, "Bad URL", newService.Name)
}

func TestServiceFailedCheck(t *testing.T) {
	service := SelectService(7)
	checked := service.Check(true)
	assert.Equal(t, "Bad URL", checked.Name)
	assert.False(t, checked.Online)
}

func TestCreateFailingTCPService(t *testing.T) {
	s := ReturnService(&types.Service{
		Name:     "Bad TCP",
		Domain:   "localhost",
		Port:     5050,
		Interval: 30,
		Type:     "tcp",
		Timeout:  5,
	})
	var err error
	newServiceId, err = s.Create()
	assert.Nil(t, err)
	assert.NotZero(t, newServiceId)
	newService := SelectService(newServiceId)
	assert.Equal(t, "Bad TCP", newService.Name)
}

func TestServiceFailedTCPCheck(t *testing.T) {
	service := SelectService(8)
	checked := service.Check(true)
	assert.Equal(t, "Bad TCP", checked.Name)
	assert.False(t, checked.Online)
}

func TestCreateServiceFailure(t *testing.T) {
	fail := &types.Failure{
		Issue:  "This is not an issue, but it would container HTTP response errors.",
		Method: "http",
	}
	service := SelectService(8)
	id, err := service.CreateFailure(fail)
	assert.Nil(t, err)
	assert.NotZero(t, id)
}

func TestDeleteService(t *testing.T) {
	service := SelectService(newServiceId)

	count, err := CoreApp.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 8, len(count))

	err = service.Delete()
	assert.Nil(t, err)

	count, err = CoreApp.SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 7, len(count))
}

func TestServiceCloseRoutine(t *testing.T) {
	s := ReturnService(new(types.Service))
	s.Name = "example"
	s.Domain = "https://google.com"
	s.Type = "http"
	s.Method = "GET"
	s.ExpectedStatus = 200
	s.Interval = 1
	s.Start()
	assert.True(t, s.IsRunning())
	t.Log(s.Checkpoint)
	go s.CheckQueue(false)
	t.Log(s.Checkpoint)
	time.Sleep(5 * time.Second)
	t.Log(s.Checkpoint)
	assert.True(t, s.IsRunning())
	s.Close()
	assert.False(t, s.IsRunning())
	s.Close()
	assert.False(t, s.IsRunning())
}

func TestServiceCheckQueue(t *testing.T) {
	s := ReturnService(new(types.Service))
	s.Name = "example"
	s.Domain = "https://google.com"
	s.Type = "http"
	s.Method = "GET"
	s.ExpectedStatus = 200
	s.Interval = 1
	s.Start()
	assert.True(t, s.IsRunning())
	go s.CheckQueue(false)

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
