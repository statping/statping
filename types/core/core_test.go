package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// NewCore return a new *core.Core struct
func NewCore() *Core {
	core := &Core{
		Started: time.Now().UTC(),
	}
	return core
}

func TestSelectCore(t *testing.T) {
	core, err := Select()
	assert.Nil(t, err)
	assert.Equal(t, "Statping Sample Data", core.Name)
}
