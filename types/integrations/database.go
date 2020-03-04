package integrations

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/services"
)

func DB() database.Database {
	return database.DB().Model(&Integration{})
}

func Find(name string) (*Integration, error) {
	var integration *Integration
	db := DB().Where("name = ?", name).Find(&integration)
	return integration, db.Error()
}

func All() []*Integration {
	var integrations []*Integration
	DB().Find(&integrations)
	return integrations
}

func List(i Integrator) ([]*services.Service, error) {
	return i.List()
}

func (i *Integration) Create() error {
	db := DB().Create(&i)
	return db.Error()
}

func (i *Integration) Update() error {
	db := DB().Update(&i)
	return db.Error()
}

func (i *Integration) Delete() error {
	db := DB().Delete(&i)
	return db.Error()
}
