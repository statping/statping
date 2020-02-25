package database

import "github.com/hunterlong/statping/types"

type GroupObj struct {
	*types.Group
	o *Object

	Grouper
}

type Grouper interface {
	Services() []*types.Service
	Model() *types.Group
}

func AllGroups() []*GroupObj {
	var groups []*types.Group
	query := database.Groups()
	query.Find(&groups)
	return wrapGroups(groups, query)
}

func (g *Db) GetGroup(id int64) (*GroupObj, error) {
	var group types.Group
	query := database.Groups().Where("id = ?", id)
	finder := query.Find(&group)
	return &GroupObj{Group: &group, o: wrapObject(id, &group, query)}, finder.Error()
}

func (g *GroupObj) Services() []*types.Service {
	var services []*types.Service
	database.Services().Where("group = ?", g.Id).Find(&services)
	return services
}

func (g *GroupObj) Model() *types.Group {
	return g.Group
}

func wrapGroups(all []*types.Group, db Database) []*GroupObj {
	var arr []*GroupObj
	for _, v := range all {
		arr = append(arr, &GroupObj{Group: v, o: wrapObject(v.Id, v, db)})
	}
	return arr
}
