package incidents

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
)

var (
	db       database.Database
	dbUpdate database.Database
	log      = utils.Log.WithField("type", "service")
)

func SetDB(database database.Database) {
	db = database.Model(&Incident{})
	dbUpdate = database.Model(&IncidentUpdate{})
}

func (i *Incident) AfterFind() {
	metrics.Query("incident", "find")
}

func (i *Incident) AfterCreate() {
	metrics.Query("incident", "create")
}

func (i *Incident) AfterUpdate() {
	metrics.Query("incident", "update")
}

func (i *Incident) AfterDelete() {
	metrics.Query("incident", "delete")
}

func (i *IncidentUpdate) AfterFind() {
	metrics.Query("incident_update", "find")
}

func (i *IncidentUpdate) AfterCreate() {
	metrics.Query("incident_update", "create")
}

func (i *IncidentUpdate) AfterUpdate() {
	metrics.Query("incident_update", "update")
}

func (i *IncidentUpdate) AfterDelete() {
	metrics.Query("incident_update", "delete")
}

func FindUpdate(uid int64) (*IncidentUpdate, error) {
	var update IncidentUpdate
	q := dbUpdate.Where("id = ?", uid).Find(&update)
	return &update, q.Error()
}

func Find(id int64) (*Incident, error) {
	var incident Incident
	q := db.Where("id = ?", id).Find(&incident)
	return &incident, q.Error()
}

func FindByService(id int64) []*Incident {
	var incidents []*Incident
	db.Where("service = ?", id).Find(&incidents)
	for _, i := range incidents {
		var updates []*IncidentUpdate
		dbUpdate.Where("incident = ?", id).Find(&updates)
		i.AllUpdates = updates
	}
	return incidents
}

func All() []*Incident {
	var incidents []*Incident
	db.Find(&incidents)
	return incidents
}

func (i *Incident) Create() error {
	q := db.Create(i)
	return q.Error()
}

func (i *Incident) Update() error {
	q := db.Update(i)
	return q.Error()
}

func (i *Incident) Delete() error {
	q := db.Delete(i)
	return q.Error()
}
