package database

import (
	"fmt"
	"github.com/hunterlong/statping/types"
	"time"
)

type GroupBy struct {
	db    Database
	query *types.GroupQuery
}

type GroupByer interface {
	ToTimeValue() (*TimeVar, error)
}

type GroupMethod interface {
}

var (
	ByCount = func() GroupMethod {
		return fmt.Sprintf("COUNT(id) as amount")
	}
	BySum = func(column string) GroupMethod {
		return fmt.Sprintf("SUM(%s) as amount", column)
	}
	ByAverage = func(column string) GroupMethod {
		return fmt.Sprintf("SUM(%s) as amount", column)
	}
)

func execute(db Database, query *types.GroupQuery) Database {
	return db.MultipleSelects(
		db.SelectByTime(query.Group),
		CountAmount(),
	).Between(query.Start, query.End).Group("timeframe").Debug()
}

func (db *Db) GroupQuery(query *types.GroupQuery) GroupByer {
	return &GroupBy{execute(db, query), query}
}

type TimeVar struct {
	g    *GroupBy
	data []*TimeValue
}

func (g *GroupBy) ToTimeValue() (*TimeVar, error) {
	rows, err := g.db.Rows()
	if err != nil {
		return nil, err
	}
	var data []*TimeValue
	for rows.Next() {
		var timeframe string
		var amount int64
		if err := rows.Scan(&timeframe, &amount); err != nil {
			return nil, err
		}
		createdTime, _ := g.db.ParseTime(timeframe)
		data = append(data, &TimeValue{
			Timeframe: createdTime,
			Amount:    amount,
		})

	}
	return &TimeVar{g, data}, nil
}

func (t *TimeVar) Values() []*TimeValue {
	var validSet []*TimeValue
	for _, v := range t.data {
		validSet = append(validSet, &TimeValue{
			Timeframe: v.Timeframe,
			Amount:    v.Amount,
		})
	}

	return validSet
}

func (t *TimeVar) FillMissing() []*TimeValue {
	timeMap := make(map[time.Time]*TimeValue)
	var validSet []*TimeValue
	if len(t.data) == 0 {
		return nil
	}
	current := t.data[0].Timeframe
	for _, v := range t.data {
		timeMap[v.Timeframe] = v
	}
	maxTime := t.g.query.End
	for {
		amount := int64(0)
		if timeMap[current] != nil {
			amount = timeMap[current].Amount
		}
		validSet = append(validSet, &TimeValue{
			Timeframe: current,
			Amount:    amount,
		})
		if current.After(maxTime) {
			break
		}
		current = current.Add(t.g.duration())
	}

	return validSet
}

func (g *GroupBy) duration() time.Duration {
	switch g.query.Group {
	case "second":
		return time.Second
	case "minute":
		return time.Minute
	case "hour":
		return time.Hour
	case "day":
		return time.Hour * 24
	case "month":
		return time.Hour * 730
	case "year":
		return time.Hour * 8760
	default:
		return time.Hour
	}
}
