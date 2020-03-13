package incidents

import "github.com/statping/statping/database"

var (
	db       database.Database
	dbUpdate database.Database
)

func SetDB(database database.Database) {
	db = database.Model(&Incident{})
	dbUpdate = database.Model(&IncidentUpdate{})
}

func Find(id int64) (*Incident, error) {
	var incident Incident
	q := db.Where("id = ?", id).Find(&incident)
	return &incident, q.Error()
}

func FindUpdate(id int64) (*IncidentUpdate, error) {
	var update IncidentUpdate
	q := dbUpdate.Where("id = ?", id).Find(&update)
	return &update, q.Error()
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
