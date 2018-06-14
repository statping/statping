package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	VERSION = "1.1.1"
	RenderBoxes()
}

func TestMakeConfig(t *testing.T) {
	config := &DbConfig{
		"postgres",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_DATABASE"),
		5432,
		"Testing",
		"admin",
		"admin",
	}
	err := config.Save()
	assert.Nil(t, err)
}

func TestSetConfig(t *testing.T) {
	configs = LoadConfig()
}

func TestRun(t *testing.T) {
	configs = LoadConfig()
	go mainProcess()
	time.Sleep(15 * time.Second)
}

func TestServiceUrl(t *testing.T) {
	req, err := http.NewRequest("GET", "/service/1", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	assert.Equal(t, 3305, len(rr.Body.Bytes()), "should be balance")
}

func Test(t *testing.T) {
	req, err := http.NewRequest("GET", "/dashboard", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)

	assert.Equal(t, 2048, len(rr.Body.Bytes()), "should be balance")
}
