package database

import (
	"github.com/hunterlong/statping/types"
	"time"
)

type HitObj struct {
	o *Object
}

func (h *HitObj) All() []*types.Hit {
	var fails []*types.Hit
	h.o.db.Find(&fails)
	return fails
}

func (h *HitObj) Last(amount int) *types.Hit {
	var hits types.Hit
	h.o.db.Limit(amount).Find(&hits)
	return &hits
}

func (h *HitObj) Since(t time.Time) []*types.Hit {
	var hits []*types.Hit
	h.o.db.Since(t).Find(&hits)
	return hits
}

func (h *HitObj) Count() int {
	var amount int
	h.o.db.Count(&amount)
	return amount
}

func (h *HitObj) object() *Object {
	return h.o
}
