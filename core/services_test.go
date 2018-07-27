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

func TestSelectService(t *testing.T) {
	service := SelectService(1)
	assert.Equal(t, "Google", service.ToService().Name)
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
	checked := ServiceHTTPCheck(service.ToService())
	assert.Equal(t, "Updated Google", checked.Name)
	assert.True(t, checked.Online)
}

func TestCheckService(t *testing.T) {
	service := SelectService(1).ToService()
	assert.Equal(t, "Updated Google", service.Name)
	assert.True(t, service.Online)
	assert.Equal(t, 200, service.LastStatusCode)
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

func TestDeleteService(t *testing.T) {
	service := SelectService(newServiceId).ToService()

	count, err := SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 6, len(count))

	err = DeleteService(service)
	assert.Nil(t, err)

	count, err = SelectAllServices()
	assert.Nil(t, err)
	assert.Equal(t, 5, len(count))
}
