package messages

import (
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/metrics"
	"gorm.io/gorm"
)

func (m *Message) Validate() error {
	if m.Title == "" {
		return errors.New("missing message title")
	}
	return nil
}

func (m *Message) BeforeUpdate(*gorm.DB) error {
	return m.Validate()
}

func (m *Message) BeforeCreate(*gorm.DB) error {
	return m.Validate()
}

func (m *Message) AfterFind(*gorm.DB) error {
	metrics.Query("message", "find")
	return nil
}

func (m *Message) AfterCreate(*gorm.DB) error {
	metrics.Query("message", "create")
	return nil
}

func (m *Message) AfterUpdate(*gorm.DB) error {
	metrics.Query("message", "update")
	return nil
}

func (m *Message) AfterDelete(*gorm.DB) error {
	metrics.Query("message", "delete")
	return nil
}
