package core

import (
	"github.com/hunterlong/statping/types"
	"sort"
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
	return err.Error
}

// Create will create a group and insert it into the database
func (g *Group) Create() (int64, error) {
	g.CreatedAt = time.Now().UTC()
	db := groupsDb().Create(g)
	return g.Id, db.Error
}

// Update will update a group
func (g *Group) Update() (int64, error) {
	g.UpdatedAt = time.Now().UTC()
	db := groupsDb().Update(g)
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

// VisibleServices returns all services based on authentication
func (g *Group) VisibleServices(auth bool) []*Service {
	var services []*Service
	for _, g := range g.Services() {
		if !g.Public.Bool {
			if auth {
				services = append(services, g)
			}
		} else {
			services = append(services, g)
		}
	}
	return services
}

// SelectGroups returns all groups
func SelectGroups(includeAll bool, auth bool) []*Group {
	var groups []*Group
	var validGroups []*Group
	groupsDb().Find(&groups).Order("order_id desc")
	for _, g := range groups {
		if !g.Public.Bool {
			if auth {
				validGroups = append(validGroups, g)
			}
		} else {
			validGroups = append(validGroups, g)
		}
	}
	sort.Sort(GroupOrder(validGroups))
	if includeAll {
		emptyGroup := &Group{&types.Group{Id: 0, Public: types.NewNullBool(true)}}
		validGroups = append(validGroups, emptyGroup)
	}
	return validGroups
}

// SelectGroup returns a *core.Group
func SelectGroup(id int64) *Group {
	for _, g := range SelectGroups(true, true) {
		if g.Id == id {
			return g
		}
	}
	return nil
}

// GroupOrder will reorder the groups based on 'order_id' (Order)
type GroupOrder []*Group

// Sort interface for resorting the Groups in order
func (c GroupOrder) Len() int           { return len(c) }
func (c GroupOrder) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c GroupOrder) Less(i, j int) bool { return c[i].Order < c[j].Order }
