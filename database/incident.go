package database

import "github.com/hunterlong/statping/types"

type IncidentObj struct {
	*types.Incident
	db Database
}

func (o *IncidentObj) AsIncident() *types.Incident {
	return o.Incident
}

func Incident(id int64) (*IncidentObj, error) {
	var incident types.Incident
	query := database.Model(&types.Incident{}).Where("id = ?", id).Find(&incident)
	return &IncidentObj{Incident: &incident, db: query}, query.Error()
}
