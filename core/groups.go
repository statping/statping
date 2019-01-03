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
	err := messagesDb().Delete(g)
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
func SelectGroups() []*Group {
	var groups []*Group
	groupsDb().Find(&groups).Order("id desc")
	return groups
}

// SelectGroup returns a *core.Group
func SelectGroup(id int64) *Group {
	for _, g := range SelectGroups() {
		if g.Id == id {
			return g
		}
	}
	return nil
}
