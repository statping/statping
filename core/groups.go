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
	err := messagesDb().Delete(g)
	if err.Error != nil {
		return err.Error
	}
	return err.Error
}

// Update will update a group in the database
func (g *Group) Update() error {
	err := servicesDB().Update(&g)
	return err.Error
}

// Create will create a group and insert it into the database
func (g *Group) Create() (int64, error) {
	g.CreatedAt = time.Now()
	db := groupsDb().Create(g)
	return g.Id, db.Error
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
