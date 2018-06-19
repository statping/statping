package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestServiceUrl(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/service/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 28, len(rr.Body.Bytes()), "should be balance")
}

func TestApiAllServiceUrl(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/api/services", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var data []Service
	json.Unmarshal(rr.Body.Bytes(), &data)
	assert.Equal(t, "Google", data[0].Name, "should be balance")
}

func TestApiServiceUrl(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/api/services/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var data Service
	json.Unmarshal(rr.Body.Bytes(), &data)
	assert.Equal(t, "Google", data.Name, "should be balance")
}

func TestApiServiceUpdateUrl(t *testing.T) {
	t.SkipNow()
	payload := []byte(`{"name":"test product - updated name","price":11.22}`)
	req, err := http.NewRequest("POST", "/api/services/1", bytes.NewBuffer(payload))
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var data Service
	json.Unmarshal(rr.Body.Bytes(), &data)
	assert.Equal(t, "Google", data.Name, "should be balance")
}

func TestApiUserUrl(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/api/users/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var data User
	json.Unmarshal(rr.Body.Bytes(), &data)
	assert.Equal(t, "testuserhere", data.Username, "should be balance")
}

func TestApiAllUsersUrl(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/api/users", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	var data []User
	json.Unmarshal(rr.Body.Bytes(), &data)
	assert.Equal(t, "testuserhere", data[0].Username, "should be balance")
}

func TestDashboardHandler(t *testing.T) {
	t.SkipNow()
	req, err := http.NewRequest("GET", "/dashboard", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 2095, len(rr.Body.Bytes()), "should be balance")
}

func TestLoginHandler(t *testing.T) {
	t.SkipNow()
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "admin")
	req, err := http.NewRequest("POST", "/dashboard", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 303, rr.Result().StatusCode, "should be balance")
}
