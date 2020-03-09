package incidents

func Samples() error {
	incident1 := &Incident{
		Title:       "Github Downtime",
		Description: "This is an example of a incident for a service.",
		ServiceId:   2,
	}
	if err := incident1.Create(); err != nil {
		return err
	}

	i1 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github's page for Statping seems to be sending a 501 error.",
		Type:       "Investigating",
	}
	if err := i1.Create(); err != nil {
		return err
	}

	i2 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Problem is continuing and we are looking at the issues.",
		Type:       "Update",
	}
	if err := i2.Create(); err != nil {
		return err
	}

	i3 := &IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}
	if err := i3.Create(); err != nil {
		return err
	}

	return nil
}

func SamplesUpdates() error {
	u1 := &IncidentUpdate{
		IncidentId: 1,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}
	if err := u1.Create(); err != nil {
		return err
	}
	return nil
}
