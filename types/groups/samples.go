package groups

import (
	"github.com/hunterlong/statping/types/null"
)

func Samples() {
	group1 := &Group{
		Name:   "Main Services",
		Public: null.NewNullBool(true),
		Order:  2,
	}
	group1.Create()

	group2 := &Group{
		Name:   "Linked Services",
		Public: null.NewNullBool(false),
		Order:  1,
	}
	group2.Create()

	group3 := &Group{
		Name:   "Empty Group",
		Public: null.NewNullBool(false),
		Order:  3,
	}
	group3.Create()

}
