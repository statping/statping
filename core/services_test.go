package core

import (
	"github.com/hunterlong/statup/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	newServiceId int64
)

func TestSelectAllServices(t *testing.T) {
	services, err := SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(services))
}

func TestSelectHTTPService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Google", service.ToService().Name)
	assert.Equal(t, "http", service.ToService().Type)
}

func TestSelectTCPService(t *testing.T) {
	service := SelectService(5)
	assert.Equal(t, "Google DNS", service.ToService().Name)
	assert.Equal(t, "tcp", service.ToService().Type)
}

func TestUpdateService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Google", service.ToService().Name)
	srv := service.ToService()
	srv.Name = "Updated Google"
	newService := UpdateService(srv)
	assert.Equal(t, "Updated Google", newService.Name)
}

func TestServiceHTTPCheck(t *testing.T) {
	service := SelectService(1)
	checked := ServiceCheck(service.ToService())
	assert.Equal(t, "Updated Google", checked.Name)
	assert.True(t, checked.Online)
}

func TestCheckHTTPService(t *testing.T) {
	service := SelectService(1).ToService()
	assert.Equal(t, "Updated Google", service.Name)
	assert.True(t, service.Online)
	assert.Equal(t, 200, service.LastStatusCode)
	assert.NotZero(t, service.Latency)
}

func TestServiceTCPCheck(t *testing.T) {
	service := SelectService(5)
	checked := ServiceCheck(service.ToService())
	assert.Equal(t, "Google DNS", checked.Name)
	assert.True(t, checked.Online)
}

func TestCheckTCPService(t *testing.T) {
	service := SelectService(5).ToService()
	assert.Equal(t, "Google DNS", service.Name)
	assert.True(t, service.Online)
	assert.NotZero(t, service.Latency)
}

func TestCreateService(t *testing.T) {
	s := &types.Service{
		Name:           "Interpol - All The Rage Back Home",
		Domain:         "https://www.youtube.com/watch?v=-u6DvRyyKGU",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
	}
	var err error
	newServiceId, err = CreateService(s)
	assert.Nil(t, err)
	assert.NotZero(t, newServiceId)
	newService := SelectService(newServiceId).ToService()
	assert.Equal(t, "Interpol - All The Rage Back Home", newService.Name)
}

func TestCreateFailingHTTPService(t *testing.T) {
	s := &types.Service{
		Name:           "Bad URL",
		Domain:         "http://localhost/iamnothere",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        5,
	}
	var err error
	newServiceId, err = CreateService(s)
	assert.Nil(t, err)
	assert.NotZero(t, newServiceId)
	newService := SelectService(newServiceId).ToService()
	assert.Equal(t, "Bad URL", newService.Name)
}

func TestServiceFailedCheck(t *testing.T) {
	service := SelectService(7)
	checked := ServiceCheck(service.ToService())
	assert.Equal(t, "Bad URL", checked.Name)
	assert.False(t, checked.Online)
}

func TestCreateFailingTCPService(t *testing.T) {
	s := &types.Service{
		Name:     "Bad TCP",
		Domain:   "localhost",
		Port:     5050,
		Interval: 30,
		Type:     "tcp",
		Timeout:  5,
	}
	var err error
	newServiceId, err = CreateService(s)
	assert.Nil(t, err)
	assert.NotZero(t, newServiceId)
	newService := SelectService(newServiceId).ToService()
	assert.Equal(t, "Bad TCP", newService.Name)
}

func TestServiceFailedTCPCheck(t *testing.T) {
	service := SelectService(8)
	checked := ServiceCheck(service.ToService())
	assert.Equal(t, "Bad TCP", checked.Name)
	assert.False(t, checked.Online)
}

func TestCreateServiceFailure(t *testing.T) {

	fail := FailureData{
		Issue: "This is not an issue, but it would container HTTP response errors.",
	}
	service := SelectService(8)

	id, err := CreateServiceFailure(service.ToService(), fail)
	assert.Nil(t, err)
	assert.NotZero(t, id)
}

func TestDeleteService(t *testing.T) {
	service := SelectService(newServiceId).ToService()

	count, err := SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 8, len(count))

	err = DeleteService(service)
	assert.Nil(t, err)

	count, err = SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 7, len(count))
}
