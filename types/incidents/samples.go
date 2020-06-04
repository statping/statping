package incidents

import (
	"github.com/statping/statping/utils"
	"time"
)

func Samples() error {
	log.Infoln("Inserting Sample Incidents...")
	incident1 := &Incident{
		Title:       "Github Issues",
		Description: "There are new features for Statping, if you have any issues please visit the Github Repo.",
		ServiceId:   4,
	}
	if err := incident1.Create(); err != nil {
		return err
	}

	incident2 := &Incident{
		Title:       "Recent Downtime",
		Description: "We've noticed an issue with authentications and we're looking into it now.",
		ServiceId:   5,
	}
	if err := incident2.Create(); err != nil {
		return err
	}

	t := utils.Now()

	i1 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github seems be be having an issue right now.",
		Type:       "investigating",
		CreatedAt:  t.Add(-60 * time.Minute),
	}
	if err := i1.Create(); err != nil {
		return err
	}

	i2 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Problem is continuing and we are looking at the issue.",
		Type:       "update",
		CreatedAt:  t.Add(-30 * time.Minute),
	}
	if err := i2.Create(); err != nil {
		return err
	}

	i3 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
		CreatedAt:  t.Add(-5 * time.Minute),
	}
	if err := i3.Create(); err != nil {
		return err
	}

	u1 := &IncidentUpdate{
		IncidentId: incident2.Id,
		Message:    "Github is acting odd, probably getting DDOS-ed by China.",
		Type:       "investigating",
		CreatedAt:  t.Add(-120 * time.Minute),
	}
	if err := u1.Create(); err != nil {
		return err
	}

	u2 := &IncidentUpdate{
		IncidentId: incident2.Id,
		Message:    "Still seems to be an issue",
		Type:       "update",
		CreatedAt:  t.Add(-60 * time.Minute),
	}

	if err := u2.Create(); err != nil {
		return err
	}

	u3 := &IncidentUpdate{
		IncidentId: incident2.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "resolved",
		CreatedAt:  t.Add(-5 * time.Minute),
	}
	if err := u3.Create(); err != nil {
		return err
	}

	return nil
}
