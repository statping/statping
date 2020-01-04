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

// AllIncidents will return all incidents and updates recorded
func AllIncidents() []*Incident {
	var incidents []*Incident
	incidentsDB().Find(&incidents).Order("id desc")
	for _, i := range incidents {
		var updates []*types.IncidentUpdate
		incidentsUpdatesDB().Find(&updates).Order("id desc")
		i.Updates = updates
	}
	return incidents
}

// Incidents will return the all incidents for a service
func (s *Service) Incidents() []*Incident {
	var incidentArr []*Incident
	incidentsDB().Where("service = ?", s.Id).Order("id desc").Find(&incidentArr)
	return incidentArr
}

// AllUpdates will return the all updates for an incident
func (i *Incident) AllUpdates() []*IncidentUpdate {
	var updatesArr []*IncidentUpdate
	incidentsUpdatesDB().Where("incident = ?", i.Id).Order("id desc").Find(&updatesArr)
	return updatesArr
}

// Delete will remove a incident
func (i *Incident) Delete() error {
	err := incidentsDB().Delete(i)
	return err.Error
}

// Create will create a incident and insert it into the database
func (i *Incident) Create() (int64, error) {
	i.CreatedAt = time.Now().UTC()
	db := incidentsDB().Create(i)
	return i.Id, db.Error
}

// Update will update a incident
func (i *Incident) Update() (int64, error) {
	i.UpdatedAt = time.Now().UTC()
	db := incidentsDB().Update(i)
	return i.Id, db.Error
}

// Delete will remove a incident update
func (i *IncidentUpdate) Delete() error {
	err := incidentsUpdatesDB().Delete(i)
	return err.Error
}

// Create will create a incident update and insert it into the database
func (i *IncidentUpdate) Create() (int64, error) {
	i.CreatedAt = time.Now().UTC()
	db := incidentsUpdatesDB().Create(i)
	return i.Id, db.Error
}
