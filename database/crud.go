package database

import (
	"reflect"
)

type Object struct {
	Id    int64
	model interface{}
	db    Database
}

type HitsFailures interface {
	Hits() *hits
	Failures() *failures
}

func modelId(model interface{}) int64 {
	iface := reflect.ValueOf(model)
	field := iface.Elem().FieldByName("Id")
	return field.Int()
}

func toModel(model interface{}) Database {
	return database.Model(&model)
}

func Create(data interface{}) (*Object, error) {
	model := toModel(data)
	query := model.Create(data)
	if query.Error() != nil {
		return nil, query.Error()
	}
	obj := &Object{
		Id:    modelId(data),
		model: data,
		db:    model,
	}
	return obj, query.Error()
}

func Update(data interface{}) error {
	model := toModel(data)
	return model.Update(&data).Error()
}

func Delete(data interface{}) error {
	model := toModel(data)
	return model.Delete(data).Error()
}
