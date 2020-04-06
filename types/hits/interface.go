package hits

import (
	// "github.com/statping/statping/types/services"
	"github.com/statping/statping/database"
	"github.com/montanaflynn/stats"
	"time"
	"database/sql"
	"fmt"
)

type Percentile struct {
	Rank int `json:"rank"`
}

var CurrentPercentile = Percentile{95} // default value

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

func (h Hitters) Percentile() int64 {
	var r int64
	var latencies []float64
	var rows *sql.Rows 
	rows, _ = h.db.Select("latency").Rows()
	
	for rows.Next() {
		rows.Scan(&r)
		latencies = append(latencies, float64(r))
	}

	percentileValue, _ := stats.Percentile(latencies, float64(CurrentPercentile.Rank))
	return int64(percentileValue)
}

func (h Hitters) Avg() int64 {
	var r IntResult
	switch h.db.DbType() {
	case "mysql":
		h.db.Select("CAST(AVG(latency) as UNSIGNED INTEGER) as amount").Scan(&r)
	case "postgres":
		h.db.Select("CAST(AVG(latency) as bigint) as amount").Scan(&r)
	default:
		h.db.Select("CAST(AVG(latency) as INT) as amount").Scan(&r)
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
