package groups

import (
	"github.com/hunterlong/statping/database"
	"github.com/hunterlong/statping/types/null"
)

func (g *Group) Samples() []database.DbObject {
	group1 := &Group{
		Name:   "Main Services",
		Public: null.NewNullBool(true),
		Order:  2,
	}

	group2 := &Group{
		Name:   "Linked Services",
		Public: null.NewNullBool(false),
		Order:  1,
	}

	group3 := &Group{
		Name:   "Empty Group",
		Public: null.NewNullBool(false),
		Order:  3,
	}

	return []database.DbObject{group1, group2, group3}
}
