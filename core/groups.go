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
func SelectGroups(includeAll bool, auth bool) map[int64]*Group {
	validGroups := make(map[int64]*Group)

	groups := database.AllGroups()

	for _, g := range groups {
		if !g.Public.Bool {
			if auth {
				validGroups[g.Id] = &Group{g.Group}
			}
		} else {
			validGroups[g.Id] = &Group{g.Group}
		}
	}
	sort.Sort(GroupOrder(validGroups))
	//if includeAll {
	//	validGroups = append(validGroups, &Group{})
	//}
	return validGroups
}

// SelectGroup returns a *core.Group
func SelectGroup(id int64) *Group {
	groups := SelectGroups(true, true)
	if groups[id] != nil {
		return groups[id]
	}
	return nil
}

// GroupOrder will reorder the groups based on 'order_id' (Order)
type GroupOrder map[int64]*Group

// Sort interface for resorting the Groups in order
func (c GroupOrder) Len() int      { return len(c) }
func (c GroupOrder) Swap(i, j int) { c[int64(i)], c[int64(j)] = c[int64(j)], c[int64(i)] }
func (c GroupOrder) Less(i, j int) bool {
	if c[int64(i)] == nil {
		return false
	}
	if c[int64(j)] == nil {
		return false
	}
	return c[int64(i)].Order < c[int64(j)].Order
}
