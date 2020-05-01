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

func (h Hitters) Sum() int64 {
	var r IntResult

	h.db.Select("CAST(SUM(latency) as INT) as amount").Scan(&r)
	return r.Amount
}

type IntResult struct {
	Amount int64
}

func (h Hitters) Avg() int64 {
	var r IntResult
	var q database.Database
	switch h.db.DbType() {
	case "mysql":
		q = h.db.Select("CAST(AVG(latency) as UNSIGNED INTEGER) as amount")
	case "postgres":
		q = h.db.Select("CAST(AVG(latency) as bigint) as amount")
	default:
		q = h.db.Select("CAST(AVG(latency) as INT) as amount")
	}
	if err := q.Scan(&r).Error(); err != nil {
		log.Errorln(err)
	}
	return r.Amount
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
