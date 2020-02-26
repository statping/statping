package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"sort"
)

type Group struct {
	*types.Group
}

// SelectGroups returns all groups
func SelectGroups(includeAll bool, auth bool) []*Group {
	var validGroups []*Group

	groups := database.AllGroups()

	for _, g := range groups {
		if !g.Public.Bool {
			if auth {
				validGroups = append(validGroups, &Group{g.Group})
			}
		} else {
			validGroups = append(validGroups, &Group{g.Group})
		}
	}
	sort.Sort(GroupOrder(validGroups))
	if includeAll {
		validGroups = append(validGroups, &Group{})
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
