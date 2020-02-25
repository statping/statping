package database

import "github.com/hunterlong/statping/types"

type IncidentObj struct {
	*types.Incident
	o *Object
}

func Incident(id int64) (*IncidentObj, error) {
	var incident types.Incident
	query := database.Incidents().Where("id = ?", id)
	finder := query.Find(&incident)
	return &IncidentObj{Incident: &incident, o: wrapObject(id, &incident, query)}, finder.Error()
}

func AllIncidents() []*types.Incident {
	var incidents []*types.Incident
	database.Incidents().Find(&incidents)
	return incidents
}

func (i *IncidentObj) Updates() []*types.IncidentUpdate {
	var incidents []*types.IncidentUpdate
	database.IncidentUpdates().Where("incident = ?", i.Id).Find(&incidents)
	return incidents
}

func (i *IncidentObj) object() *Object {
	return i.o
}
