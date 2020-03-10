package incidents

func (i *Incident) Updates() []*IncidentUpdate {
	var updates []*IncidentUpdate
	db.Model(&IncidentUpdate{}).Where("incident = ?", i.Id).Find(&updates)
	i.AllUpdates = updates
	return updates
}

func (i *IncidentUpdate) Create() error {
	q := db.Create(i)
	return q.Error()
}

func (i *IncidentUpdate) Update() error {
	q := db.Update(i)
	return q.Error()
}

func (i *IncidentUpdate) Delete() error {
	q := db.Delete(i)
	return q.Error()
}
