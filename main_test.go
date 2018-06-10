package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMakeConfig(t *testing.T) {
	config := &DbConfig{
		"postgres",
		"localhost",
		"travis",
		"",
		"postgres",
		5432,
		"Testing",
		"admin",
		"admin",
	}
	err := config.Save()
	assert.Nil(t, err)
}

func Test(t *testing.T) {

	assert.Equal(t, "", "")

}