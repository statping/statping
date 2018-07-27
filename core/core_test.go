package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testCore *Core
)

func TestNewCore(t *testing.T) {
	testCore = NewCore()
	assert.NotNil(t, testCore)
}
