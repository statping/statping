package database

import (
	"github.com/hunterlong/statping/types"
	"reflect"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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

type CreateCallback func(interface{}, error)

func runCallbacks(data interface{}, err error, fns ...AfterCreate) {
	for _, fn := range fns {
		fn.AfterCreate(data, err)
	}
}

type AfterCreate interface {
	AfterCreate(interface{}, error)
}

func Create(data interface{}, fns ...AfterCreate) (*Object, error) {
	model := database.Model(&data)
	if err := model.Create(data).Error(); err != nil {
		runCallbacks(data, err, fns...)
		return nil, err
	}
	obj := &Object{
		Id:    modelId(data),
		model: data,
		db:    model,
	}
	runCallbacks(data, nil, fns...)
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
