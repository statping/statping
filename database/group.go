package database

import (
	"fmt"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type GroupBy struct {
	db    Database
	query *GroupQuery
}

type GroupByer interface {
	ToTimeValue(interface{}) (*TimeVar, error)
}

type By string

func (b By) String() string {
	return string(b)
}

type GroupQuery struct {
	db        Database
	Start     time.Time
	End       time.Time
	Group     string
	Order     string
	Limit     int
	Offset    int
	FillEmpty bool
}

func (b GroupQuery) Database() Database {
	return b.db
}

var (
	ByCount = By("COUNT(id) as amount")
	BySum   = func(column string) By {
		return By(fmt.Sprintf("SUM(%s) as amount", column))
	}
	ByAverage = func(column string) By {
		return By(fmt.Sprintf("AVG(%s) as amount", column))
	}
)

func (db *Db) GroupQuery(q *GroupQuery, by By) GroupByer {
	dbQuery := db.MultipleSelects(
		db.SelectByTime(q.Group),
		by.String(),
	).Group("timeframe")

	return &GroupBy{dbQuery, q}
}

type TimeVar struct {
	g    *GroupBy
	data []*TimeValue
}

func (t *TimeVar) ToValues() []*TimeValue {
	return t.data
}

func (g *GroupBy) toFloatRows() []*TimeValue {
	rows, err := g.db.Rows()
	if err != nil {
		return nil
	}
	var data []*TimeValue
	for rows.Next() {
		var timeframe time.Time
		amount := float64(0)
		rows.Scan(&timeframe, &amount)
		newTs := types.FixedTime(timeframe, g.duration())
		data = append(data, &TimeValue{
			Timeframe: newTs,
			Amount:    amount,
		})
	}
	return data
}

func (g *GroupBy) ToTimeValue(dbType interface{}) (*TimeVar, error) {
	rows, err := g.db.Rows()
	if err != nil {
		return nil, err
	}
	var data []*TimeValue
	for rows.Next() {
		var timeframe time.Time
		amount := float64(0)
		rows.Scan(&timeframe, &amount)
		newTs := types.FixedTime(timeframe, g.duration())
		data = append(data, &TimeValue{
			Timeframe: newTs,
			Amount:    amount,
		})
	}
	return &TimeVar{g, data}, nil
}

func (t *TimeVar) FillMissing(current, end time.Time) []*TimeValue {
	timeMap := make(map[string]float64)
	var validSet []*TimeValue
	dur := t.g.duration()
	for _, v := range t.data {
		timeMap[v.Timeframe] = v.Amount
	}

	currentStr := types.FixedTime(current, t.g.duration())

	for {
		var amount float64
		if timeMap[currentStr] != 0 {
			amount = timeMap[currentStr]
		}
		validSet = append(validSet, &TimeValue{
			Timeframe: currentStr,
			Amount:    amount,
		})
		if current.After(end) {
			break
		}
		current = current.Add(dur)
		currentStr = types.FixedTime(current, t.g.duration())
	}

	return validSet
}

func (g *GroupBy) duration() time.Duration {
	switch g.query.Group {
	case "second":
		return types.Second
	case "minute":
		return types.Minute
	case "hour":
		return types.Hour
	case "day":
		return types.Day
	case "month":
		return types.Month
	case "year":
		return types.Year
	default:
		return types.Hour
	}
}

func ParseQueries(r *http.Request, db Database) *GroupQuery {
	fields := parseGet(r)
	grouping := fields.Get("group")
	if grouping == "" {
		grouping = "hour"
	}
	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	limit := utils.ToInt(fields.Get("limit"))
	offset := utils.ToInt(fields.Get("offset"))
	fill, _ := strconv.ParseBool(fields.Get("fill"))
	orderBy := fields.Get("order")
	if limit == 0 {
		limit = 10000
	}

	query := &GroupQuery{
		Start:     time.Unix(startField, 0).UTC(),
		End:       time.Unix(endField, 0).UTC(),
		Group:     grouping,
		Order:     orderBy,
		Limit:     int(limit),
		Offset:    int(offset),
		FillEmpty: fill,
	}

	if query.Limit != 0 {
		db = db.Limit(query.Limit)
	}
	if query.Offset > 0 {
		db = db.Offset(query.Offset)
	}
	if !query.Start.IsZero() && !query.End.IsZero() {
		db = db.Where("created_at BETWEEN ? AND ?", db.FormatTime(query.Start), db.FormatTime(query.End))
	} else {
		if !query.Start.IsZero() {
			db = db.Where("created_at > ?", db.FormatTime(query.Start))
		}
		if !query.End.IsZero() {
			db = db.Where("created_at < ?", db.FormatTime(query.End))
		}
	}
	if query.Order != "" {
		db = db.Order(query.Order)
	}
	query.db = db.Debug()

	return query
}

func parseForm(r *http.Request) url.Values {
	r.ParseForm()
	return r.PostForm
}

func parseGet(r *http.Request) url.Values {
	r.ParseForm()
	return r.Form
}
