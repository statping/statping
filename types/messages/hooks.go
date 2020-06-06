package messages

import (
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
