package incidents

import "github.com/hunterlong/statping/database"

func (i *Incident) Updates() []*IncidentUpdate {
	var updates []*IncidentUpdate
	database.DB().Model(&IncidentUpdate{}).Where("incident = ?", i.Id).Find(&updates)
	i.AllUpdates = updates
	return updates
}

func (i *IncidentUpdate) Create() error {
	db := database.DB().Create(i)
	return db.Error()
}

func (i *IncidentUpdate) Update() error {
	db := database.DB().Update(i)
	return db.Error()
}

func (i *IncidentUpdate) Delete() error {
	db := database.DB().Delete(i)
	return db.Error()
}
