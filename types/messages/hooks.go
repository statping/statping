package messages

import (
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/metrics"
)

func (m *Message) Validate() error {
	if m.Title == "" {
		return errors.New("missing message title")
	}
	return nil
}

func (m *Message) BeforeUpdate() error {
	return m.Validate()
}

func (m *Message) BeforeCreate() error {
	return m.Validate()
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
