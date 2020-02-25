package database

import "github.com/hunterlong/statping/types"

type CheckinHitObj struct {
	hits []*types.CheckinHit
	o    *Object
}
