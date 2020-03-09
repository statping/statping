package groups

import (
	"github.com/hunterlong/statping/types/null"
)

func Samples() error {
	group1 := &Group{
		Name:   "Main Services",
		Public: null.NewNullBool(true),
		Order:  2,
	}
	if err := group1.Create(); err != nil {
		return err
	}

	group2 := &Group{
		Name:   "Linked Services",
		Public: null.NewNullBool(false),
		Order:  1,
	}
	if err := group2.Create(); err != nil {
		return err
	}

	group3 := &Group{
		Name:   "Empty Group",
		Public: null.NewNullBool(false),
		Order:  3,
	}
	if err := group3.Create(); err != nil {
		return err
	}

	return nil
}
