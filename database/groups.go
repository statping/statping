package database

import "github.com/hunterlong/statping/types"

type Group struct {
	db    Database
	group *types.Group
}

type Groupser interface {
	Services() Database
}

func (it *Db) GetGroup(id int64) (Groupser, error) {
	var group types.Group
	query := it.Model(&types.Group{}).Where("id = ?", id).Find(&group)
	return &Group{it, &group}, query.Error()
}

func (it *Group) Services() Database {
	return it.db.Model(&types.Service{}).Where("group = ?", it.group.Id)
}
