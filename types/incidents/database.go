package incidents

import "github.com/statping/statping/database"

var db database.Database

func SetDB(database database.Database) {
	db = database.Model(&Incident{})
}

func Find(id int64) (*Incident, error) {
	var incident Incident
	q := db.Where("id = ?", id).Find(&incident)
	return &incident, q.Error()
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
