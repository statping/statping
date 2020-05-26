package messages

import (
	"time"
)

func Samples() error {
	log.Infoln("Inserting Sample Messages...")
	m1 := &Message{
		Title:       "Routine Downtime",
		Description: "This is an example a upcoming message for a service!",
		ServiceId:   1,
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
	}

	if err := m1.Create(); err != nil {
		return err
	}

	m2 := &Message{
		Title:       "Server Reboot",
		Description: "This is another example a upcoming message for a service!",
		ServiceId:   3,
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
	}

	if err := m2.Create(); err != nil {
		return err
	}

	return nil
}
