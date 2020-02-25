package database

import (
	"fmt"
	"github.com/hunterlong/statping/types"
	"reflect"
)

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
	fmt.Printf("%T\n", model)
	iface := reflect.ValueOf(model)
	field := iface.Elem().FieldByName("Id")
	return field.Int()
}

func toModel(model interface{}) Database {
	switch model.(type) {
	case *types.Core:
		return database.Model(&types.Core{}).Table("core")
	default:
		return database.Model(&model)
	}
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
