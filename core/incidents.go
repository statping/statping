package core

import (
	"github.com/hunterlong/statping/types"
	"time"
)

type Incident struct {
	*types.Incident
}

type IncidentUpdate struct {
	*types.IncidentUpdate
}

// ReturnIncident returns *core.Incident based off a *types.Incident
func ReturnIncident(u *types.Incident) *Incident {
	return &Incident{u}
}

// SelectIncident returns the Incident based on the Incident's ID.
func SelectIncident(id int64) (*Incident, error) {
	var incident Incident
	err := Database(incident).Where("id = ?", id).First(&incident)
	return &incident, err.Error()
}

// AllIncidents will return all incidents and updates recorded
func AllIncidents() []*Incident {
	var incidents []*Incident
	Database(incidents).Find(&incidents).Order("id desc")
	for _, i := range incidents {
		var updates []*types.IncidentUpdate
		Database(updates).Find(&updates).Order("id desc")
		i.Updates = updates
	}
	return incidents
}

// Incidents will return the all incidents for a service
func (s *Service) Incidents() []*Incident {
	var incidentArr []*Incident
	Database(incidentArr).Where("service = ?", s.Id).Order("id desc").Find(&incidentArr)
	return incidentArr
}

// AllUpdates will return the all updates for an incident
func (i *Incident) AllUpdates() []*IncidentUpdate {
	var updatesArr []*IncidentUpdate
	Database(updatesArr).Where("incident = ?", i.Id).Order("id desc").Find(&updatesArr)
	return updatesArr
}

// Delete will remove a incident
func (i *Incident) Delete() error {
	err := Database(i).Delete(i)
	return err.Error()
}

// Create will create a incident and insert it into the database
func (i *Incident) Create() (int64, error) {
	i.CreatedAt = time.Now().UTC()
	db := Database(i).Create(i)
	return i.Id, db.Error()
}

// Update will update a incident
func (i *Incident) Update() (int64, error) {
	i.UpdatedAt = time.Now().UTC()
	db := Database(i).Update(i)
	return i.Id, db.Error()
}

// Delete will remove a incident update
func (i *IncidentUpdate) Delete() error {
	err := Database(i).Delete(i)
	return err.Error()
}

// Create will create a incident update and insert it into the database
func (i *IncidentUpdate) Create() (int64, error) {
	i.CreatedAt = time.Now().UTC()
	db := Database(i).Create(i)
	return i.Id, db.Error()
}
