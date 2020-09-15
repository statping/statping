package incidents

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
	"gorm.io/gorm"
)

var (
	db       *database.Database
	dbUpdate *database.Database
	log      = utils.Log.WithField("type", "service")
)

func SetDB(dbz *database.Database) {
	db = database.Wrap(dbz.Model(&Incident{}))
	dbUpdate = database.Wrap(dbz.Model(&IncidentUpdate{}))
}

func (i *Incident) Validate() error {
	if i.Title == "" {
		return errors.New("missing title")
	}
	return nil
}

func (i *Incident) BeforeUpdate(*gorm.DB) error {
	return i.Validate()
}

func (i *Incident) BeforeCreate(*gorm.DB) error {
	return i.Validate()
}

func (i *Incident) AfterFind(*gorm.DB) error {
	db.Model(i).Association("Updates").Find(&i.Updates)
	metrics.Query("incident", "find")
	return nil
}

func (i *Incident) AfterCreate(*gorm.DB) error {
	metrics.Query("incident", "create")
	return nil
}

func (i *Incident) AfterUpdate(*gorm.DB) error {
	metrics.Query("incident", "update")
	return nil
}

func (i *Incident) AfterDelete(*gorm.DB) error {
	metrics.Query("incident", "delete")
	return nil
}

func (i *IncidentUpdate) Validate() error {
	if i.Message == "" {
		return errors.New("missing incident update title")
	}
	return nil
}

func (i *IncidentUpdate) BeforeUpdate(*gorm.DB) error {
	return i.Validate()
}

func (i *IncidentUpdate) BeforeCreate(*gorm.DB) error {
	return i.Validate()
}

func (i *IncidentUpdate) AfterFind(*gorm.DB) error {
	metrics.Query("incident_update", "find")
	return nil
}

func (i *IncidentUpdate) AfterCreate(*gorm.DB) error {
	metrics.Query("incident_update", "create")
	return nil
}

func (i *IncidentUpdate) AfterUpdate(*gorm.DB) error {
	metrics.Query("incident_update", "update")
	return nil
}

func (i *IncidentUpdate) AfterDelete(*gorm.DB) error {
	metrics.Query("incident_update", "delete")
	return nil
}

func FindUpdate(uid int64) (*IncidentUpdate, error) {
	var update IncidentUpdate
	q := dbUpdate.Where("id = ?", uid).Find(&update)
	return &update, q.Error
}

func Find(id int64) (*Incident, error) {
	var incident Incident
	q := db.Where("id = ?", id).Find(&incident)
	return &incident, q.Error
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
	return db.Create(i).Error
}

func (i *Incident) Update() error {
	return db.Save(i).Error
}

func (i *Incident) Delete() error {
	for _, u := range i.Updates {
		if err := u.Delete(); err != nil {
			return err
		}
	}
	return db.Delete(i).Error
}
