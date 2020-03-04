package incidents

import (
	"github.com/hunterlong/statping/database"
)

func (s *Incident) Samples() []database.DbObject {
	incident1 := &Incident{
		Title:       "Github Downtime",
		Description: "This is an example of a incident for a service.",
		ServiceId:   2,
	}

	i1 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github's page for Statping seems to be sending a 501 error.",
		Type:       "Investigating",
	}

	i2 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Problem is continuing and we are looking at the issues.",
		Type:       "Update",
	}

	i3 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}

	return []database.DbObject{i1, i2, i3}
}

func (s *IncidentUpdate) Samples() []database.DbObject {

	return nil
}
