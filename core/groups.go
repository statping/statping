package core

import (
	"github.com/hunterlong/statping/types"
	"time"
)

type Group struct {
	*types.Group
}

// Delete will remove a group
func (g *Group) Delete() error {
	for _, s := range g.Services() {
		s.GroupId = 0
		s.Update(false)
	}
	err := groupsDb().Delete(g)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

// Create will create a group and insert it into the database
func (g *Group) Create() (int64, error) {
	g.CreatedAt = time.Now()
	db := groupsDb().Create(g)
	return g.Id, db.Error
}

// Services returns all services belonging to a group
func (g *Group) Services() []*Service {
	var services []*Service
	for _, s := range Services() {
		if s.Select().GroupId == int(g.Id) {
			services = append(services, s.(*Service))
		}
	}
	return services
}

// SelectGroups returns all groups
func SelectGroups(includeAll bool, auth bool) []*Group {
	var groups []*Group
	var validGroups []*Group
	groupsDb().Find(&groups).Order("id desc")
	if includeAll {
		emptyGroup := &Group{&types.Group{Id: 0, Public: types.NewNullBool(true)}}
		groups = append(groups, emptyGroup)
	}
	for _, g := range groups {
		if !g.Public.Bool {
			if auth {
				validGroups = append(validGroups, g)
			}
		} else {
			validGroups = append(validGroups, g)
		}
	}
	return validGroups
}

// SelectGroup returns a *core.Group
func SelectGroup(id int64) *Group {
	for _, g := range SelectGroups(false, false) {
		if g.Id == id {
			return g
		}
	}
	return nil
}
