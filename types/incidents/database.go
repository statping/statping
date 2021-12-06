package incidents

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/errors"
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

func (i *Incident) Validate() error {
	if i.Title == "" {
		return errors.New("missing title")
	}
	return nil
}

func (i *Incident) BeforeUpdate() error {
	return i.Validate()
}

func (i *Incident) BeforeCreate() error {
	return i.Validate()
}

func (i *Incident) AfterFind() {
	db.Model(i).Related(&i.Updates).Order("id DESC")
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

func (i *IncidentUpdate) Validate() error {
	if i.Message == "" {
		return errors.New("missing incident update title")
	}
	return nil
}

func (i *IncidentUpdate) BeforeUpdate() error {
	return i.Validate()
}

func (i *IncidentUpdate) BeforeCreate() error {
	return i.Validate()
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
	return incidents
}

func All() []*Incident {
	var incidents []*Incident
	db.Find(&incidents)
	return incidents
}

func (i *Incident) Create() error {
	return db.Create(i).Error()
}

func (i *Incident) Update() error {
	return db.Update(i).Error()
}

func (i *Incident) Delete() error {
	for _, u := range i.Updates {
		if err := u.Delete(); err != nil {
			return err
		}
	}
	return db.Delete(i).Error()
}
