package messages

import (
	"github.com/hunterlong/statping/database"
	"time"
)

func (m *Message) Samples() []database.DbObject {
	m1 := &Message{
		Title:       "Routine Downtime",
		Description: "This is an example a upcoming message for a service!",
		ServiceId:   1,
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
	}

	m2 := &Message{
		Title:       "Server Reboot",
		Description: "This is another example a upcoming message for a service!",
		ServiceId:   3,
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
	}

	return []database.DbObject{m1, m2}
}
