package incidents

//func (i Incident) Updates() []*IncidentUpdate {
//	var updates []*IncidentUpdate
//	dbUpdate.Where("incident = ?", i.Id).Find(&updates)
//	return updates
//}

func (i *IncidentUpdate) Create() error {
	q := dbUpdate.Create(i)
	return q.Error()
}

func (i *IncidentUpdate) Update() error {
	q := dbUpdate.Update(i)
	return q.Error()
}

func (i *IncidentUpdate) Delete() error {
	q := dbUpdate.Delete(i)
	return q.Error()
}
