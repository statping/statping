package incidents

func Samples() {
	incident1 := &Incident{
		Title:       "Github Downtime",
		Description: "This is an example of a incident for a service.",
		ServiceId:   2,
	}
	incident1.Create()

	i1 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github's page for Statping seems to be sending a 501 error.",
		Type:       "Investigating",
	}
	i1.Create()

	i2 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Problem is continuing and we are looking at the issues.",
		Type:       "Update",
	}
	i2.Create()

	i3 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}
	i3.Create()
}

func SamplesUpdates() {
	u1 := &IncidentUpdate{
		IncidentId: 1,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}
	u1.Create()
}
