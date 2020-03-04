package groups

import (
	"github.com/hunterlong/statping/database"
	"sort"
)

func Find(id int64) (*Group, error) {
	var group *Group
	db := database.DB().Model(&Group{}).Where("id = ?", id).Find(&group)
	return group, db.Error()
}

func All() []*Group {
	var groups []*Group
	database.DB().Model(&Group{}).Find(&groups)
	return groups
}

func (g *Group) Create() error {
	db := database.DB().Create(&g)
	return db.Error()
}

func (g *Group) Update() error {
	db := database.DB().Update(&g)
	return db.Error()
}

func (g *Group) Delete() error {
	db := database.DB().Delete(&g)
	return db.Error()
}

// SelectGroups returns all groups
func SelectGroups(includeAll bool, auth bool) []*Group {
	var validGroups []*Group

	all := All()
	if includeAll {
		sort.Sort(GroupOrder(all))
		return all
	}

	for _, g := range all {
		if !g.Public.Bool {
			if auth {
				validGroups = append(validGroups, g)
			}
		} else {
			validGroups = append(validGroups, g)
		}
	}
	sort.Sort(GroupOrder(validGroups))
	return validGroups
}
