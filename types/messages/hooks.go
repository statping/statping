package messages

import (
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
)

// BeforeCreate for Message will set CreatedAt to UTC
func (m *Message) BeforeCreate() (err error) {
	if m.CreatedAt.IsZero() {
		m.CreatedAt = utils.Now()
		m.UpdatedAt = utils.Now()
	}
	return
}

func (m *Message) AfterFind() {
	metrics.Query("message", "find")
}

func (m *Message) AfterCreate() {
	metrics.Query("message", "create")
}

func (m *Message) AfterUpdate() {
	metrics.Query("message", "update")
}

func (m *Message) AfterDelete() {
	metrics.Query("message", "delete")
}
