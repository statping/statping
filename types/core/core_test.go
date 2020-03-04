package core

import (
	"github.com/hunterlong/statping/core"
	"github.com/hunterlong/statping/types/services"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// NewCore return a new *core.Core struct
func NewCore() *Core {
	core := &Core{
		Started: time.Now().UTC(),
	}
	core.services = make(map[int64]*services.Service)
	return core
}

func TestSelectCore(t *testing.T) {
	core, err := Select()
	assert.Nil(t, err)
	assert.Equal(t, "Statping Sample Data", core.Name)
}
