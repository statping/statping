package database

import "github.com/hunterlong/statping/types"

type GroupObj struct {
	*types.Group
	db Database
}

type Grouper interface {
	Services() Database
}

func (o *Object) AsGroup() *types.Group {
	return o.model.(*types.Group)
}

func (it *Db) GetGroup(id int64) (*GroupObj, error) {
	var group types.Group
	query := it.Model(&types.Group{}).Where("id = ?", id).Find(&group)
	return &GroupObj{&group, it}, query.Error()
}

func (it *GroupObj) Services() Database {
	return it.db.Model(&types.Service{}).Where("service = ?", it.Id)
}
