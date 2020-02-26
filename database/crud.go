package database

import (
	"github.com/hunterlong/statping/types"
	"reflect"
)

type CrudObject interface {
	Create()
}

type Object struct {
	Id    int64
	model interface{}
	db    Database
}

type isObject interface {
	object() *Object
}

func wrapObject(id int64, model interface{}, db Database) *Object {
	return &Object{
		Id:    id,
		model: model,
		db:    db,
	}
}

func modelId(model interface{}) int64 {
	switch model.(type) {
	case *types.Core:
		return 0
	default:
		iface := reflect.ValueOf(model)
		field := iface.Elem().FieldByName("Id")
		return field.Int()
	}
}

func Create(data interface{}) (*Object, error) {
	model := database.Model(&data)
	if err := model.Create(data).Error(); err != nil {
		return nil, err
	}
	obj := &Object{
		Id:    modelId(data),
		model: data,
		db:    model,
	}
	return obj, nil
}

func Update(data interface{}) error {
	model := database.Model(&data)
	return model.Update(&data).Error()
}

func Delete(data interface{}) error {
	model := database.Model(&data)
	return model.Delete(data).Error()
}
