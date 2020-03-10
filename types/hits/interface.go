package hits

import (
	"fmt"
	"github.com/statping/statping/database"
	"time"
)

type ColumnIDInterfacer interface {
	HitsColumnID() (string, int64)
}

type Hitters struct {
	db database.Database
}

func (h Hitters) Db() database.Database {
	return h.db
}

func (h Hitters) First() *Hit {
	var hit Hit
	h.db.Order("id ASC").Limit(1).Find(&hit)
	return &hit
}

func (h Hitters) Last() *Hit {
	var hit Hit
	h.db.Order("id DESC").Limit(1).Find(&hit)
	return &hit
}

func (h Hitters) Since(t time.Time) []*Hit {
	var hits []*Hit
	h.db.Since(t).Find(&hits)
	return hits
}

func (h Hitters) List() []*Hit {
	var hits []*Hit
	h.db.Find(&hits)
	return hits
}

func (h Hitters) LastAmount(amount int) []*Hit {
	var hits []*Hit
	h.db.Order("id asc").Limit(amount).Find(&hits)
	return hits
}

func (h Hitters) Count() int {
	var count int
	h.db.Count(&count)
	return count
}

func (h Hitters) DeleteAll() error {
	q := h.db.Delete(&Hit{})
	return q.Error()
}

func (h Hitters) Sum() float64 {
	result := struct {
		amount float64
	}{0}

	h.db.Select("AVG(latency) as amount").Scan(&result)
	return result.amount
}

func (h Hitters) Avg() float64 {
	result := struct {
		amount float64
	}{0}

	h.db.Select("AVG(latency) as amount").Scan(&result)
	return result.amount
}

func AllHits(obj ColumnIDInterfacer) Hitters {
	column, id := obj.HitsColumnID()
	return Hitters{db.Where(fmt.Sprintf("%s = ?", column), id)}
}

func Since(t time.Time, obj ColumnIDInterfacer) Hitters {
	column, id := obj.HitsColumnID()
	timestamp := db.FormatTime(t)
	return Hitters{db.Where(fmt.Sprintf("%s = ? AND created_at > ?", column), id, timestamp)}
}
