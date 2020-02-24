package database

import "github.com/hunterlong/statping/types"

type hits struct {
	DB Database
}

func (h *hits) All() []*types.Hit {
	var fails []*types.Hit
	h.DB = h.DB.Find(&fails)
	return fails
}

func (h *hits) Last(amount int) *types.Hit {
	var hits types.Hit
	h.DB = h.DB.Limit(amount).Find(&hits)
	return &hits
}

func (h *hits) Count() int {
	var amount int
	h.DB = h.DB.Count(&amount)
	return amount
}

func (h *hits) Find(data interface{}) error {
	q := h.Find(&data)
	return q
}
