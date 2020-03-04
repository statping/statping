package groups

import (
	"github.com/hunterlong/statping/database"
	"sort"
)

func DB() database.Database {
	return database.DB().Model(&Group{})
}

func Find(id int64) (*Group, error) {
	var group *Group
	db := DB().Where("id = ?", id).Find(&group)
	return group, db.Error()
}

func All() []*Group {
	var groups []*Group
	DB().Find(&groups)
	return groups
}

func (g *Group) Create() error {
	db := DB().Create(&g)
	return db.Error()
}

func (g *Group) Update() error {
	db := DB().Update(&g)
	return db.Error()
}

func (g *Group) Delete() error {
	db := DB().Delete(&g)
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
