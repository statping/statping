package groups

import (
	"github.com/statping/statping/database"
	"github.com/statping/statping/types/errors"
	"github.com/statping/statping/types/metrics"
	"github.com/statping/statping/utils"
	"gorm.io/gorm"
	"sort"
)

var (
	db  *database.Database
	log = utils.Log.WithField("type", "group")
)

func SetDB(dbz *database.Database) {
	db = database.Wrap(dbz.Model(&Group{}))
}

func (g *Group) Validate() error {
	if g.Name == "" {
		return errors.New("group name is empty")
	}
	return nil
}

func (g *Group) AfterFind(*gorm.DB) error {
	metrics.Query("group", "find")
	return nil
}

func (g *Group) AfterUpdate(*gorm.DB) error {
	metrics.Query("group", "update")
	return nil
}

func (g *Group) AfterDelete(*gorm.DB) error {
	metrics.Query("group", "delete")
	return nil
}

func (g *Group) BeforeUpdate(*gorm.DB) error {
	return g.Validate()
}

func (g *Group) BeforeCreate(*gorm.DB) error {
	return g.Validate()
}

func (g *Group) AfterCreate(*gorm.DB) error {
	metrics.Query("group", "create")
	return nil
}

func Find(id int64) (*Group, error) {
	var group Group
	q := db.Where("id = ?", id).Find(&group)
	if q.Error != nil {
		return nil, errors.Missing(group, id)
	}
	return &group, q.Error
}

func All() []*Group {
	var groups []*Group
	db.Find(&groups)
	return groups
}

func (g *Group) Create() error {
	q := db.Create(g)
	return q.Error
}

func (g *Group) Update() error {
	q := db.Save(g)
	return q.Error
}

func (g *Group) Delete() error {
	q := db.Delete(g)
	return q.Error
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
