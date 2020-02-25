package core

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types"
	"sort"
)

type Group struct {
	database.Grouper
}

// SelectGroups returns all groups
func SelectGroups(includeAll bool, auth bool) []database.Grouper {
	var validGroups []database.Grouper

	groups := database.AllGroups()

	for _, g := range groups {
		if !g.Model().Public.Bool {
			if auth {
				validGroups = append(validGroups, g)
			}
		} else {
			validGroups = append(validGroups, g)
		}
	}
	sort.Sort(GroupOrder(validGroups))
	if includeAll {
		emptyGroup := &Group{}
		validGroups = append(validGroups, emptyGroup)
	}
	return validGroups
}

// SelectGroup returns a *core.Group
func SelectGroup(id int64) *types.Group {
	for _, g := range SelectGroups(true, true) {
		if g.Model().Id == id {
			return g.Model()
		}
	}
	return nil
}

// GroupOrder will reorder the groups based on 'order_id' (Order)
type GroupOrder []database.Grouper

// Sort interface for resorting the Groups in order
func (c GroupOrder) Len() int           { return len(c) }
func (c GroupOrder) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c GroupOrder) Less(i, j int) bool { return c[i].Model().Order < c[j].Model().Order }
