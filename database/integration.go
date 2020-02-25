package database

import "github.com/hunterlong/statping/types"

type IntegrationObj struct {
	*types.Integration
	o *Object
}

func Integration(id int64) (*IntegrationObj, error) {
	var integration types.Integration
	query := database.Model(&types.Integration{}).Where("id = ?", id)
	finder := query.Find(&integration)
	return &IntegrationObj{Integration: &integration, o: wrapObject(id, &integration, query)}, finder.Error()
}

func (i *IntegrationObj) object() *Object {
	return i.o
}
