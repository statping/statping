package incidents

import "github.com/hunterlong/statping/database"

func Find(id int64) (*Incident, error) {
	var incident Incident
	db := database.DB().Model(&Incident{}).Where("id = ?", id).Find(&incident)
	return &incident, db.Error()
}

func All() []*Incident {
	var incidents []*Incident
	database.DB().Model(&Incident{}).Find(&incidents)
	return incidents
}

func (i *Incident) Create() error {
	db := database.DB().Create(i)
	return db.Error()
}

func (i *Incident) Update() error {
	db := database.DB().Update(i)
	return db.Error()
}

func (i *Incident) Delete() error {
	db := database.DB().Delete(i)
	return db.Error()
}
