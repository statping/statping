package incidents

import (
	"github.com/statping/statping/utils"
	"time"
)

func Samples() error {
	incident1 := &Incident{
		Title:       "Github Downtime",
		Description: "This is an example of a incident for a service.",
		ServiceId:   2,
	}
	if err := incident1.Create(); err != nil {
		return err
	}

	incident2 := &Incident{
		Title:       "Recent Downtime",
		Description: "We've noticed an issue with authentications and we're looking into it now.",
		ServiceId:   4,
	}
	if err := incident2.Create(); err != nil {
		return err
	}

	return nil
}

func SamplesUpdates() error {
	t := utils.Now()

	i1 := &IncidentUpdate{
		IncidentId: 1,
		Message:    "Github's page for Statping seems to be sending a 501 error.",
		Type:       "Investigating",
		CreatedAt:  t.Add(-60 * time.Minute),
	}
	if err := i1.Create(); err != nil {
		return err
	}

	i2 := &IncidentUpdate{
		IncidentId: 1,
		Message:    "Problem is continuing and we are looking at the issues.",
		Type:       "Update",
		CreatedAt:  t.Add(-30 * time.Minute),
	}
	if err := i2.Create(); err != nil {
		return err
	}

	i3 := &IncidentUpdate{
		IncidentId: 1,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
		CreatedAt:  t.Add(-5 * time.Minute),
	}
	if err := i3.Create(); err != nil {
		return err
	}

	u1 := &IncidentUpdate{
		IncidentId: 2,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
		CreatedAt:  t.Add(-120 * time.Minute),
	}
	if err := u1.Create(); err != nil {
		return err
	}

	u2 := &IncidentUpdate{
		IncidentId: 2,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
		CreatedAt:  t.Add(-60 * time.Minute),
	}
	if err := u2.Create(); err != nil {
		return err
	}
	return nil
}
