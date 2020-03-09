package database

import (
	"fmt"
	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
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
	ToTimeValue() (*TimeVar, error)
}

type By string

func (b By) String() string {
	return string(b)
}

type GroupQuery struct {
	Start     time.Time
	End       time.Time
	Group     time.Duration
	Order     string
	Limit     int
	Offset    int
	FillEmpty bool

	db Database
}

func (b GroupQuery) Find(data interface{}) error {
	return b.db.Find(data).Error()
}

func (b GroupQuery) Database() Database {
	return b.db
}

var (
	ByCount   = By("COUNT(id) as amount")
	ByAverage = func(column string, multiplier int) By {
		switch database.DbType() {
		case "mysql":
			return By(fmt.Sprintf("CAST(AVG(%s)*%d as UNSIGNED) as amount", column, multiplier))
		case "postgres":
			return By(fmt.Sprintf("cast(AVG(%s)*%d as int) as amount", column, multiplier))
		default:
			return By(fmt.Sprintf("cast(AVG(%s)*%d as int) as amount", column, multiplier))
		}
	}
)

type TimeVar struct {
	g    *GroupQuery
	data []*TimeValue
}

func (t *TimeVar) ToValues() ([]*TimeValue, error) {
	return t.data, nil
}

// GraphData will return all hits or failures
func (g *GroupQuery) GraphData(by By) ([]*TimeValue, error) {

	dbQuery := g.db.MultipleSelects(
		g.db.SelectByTime(g.Group),
		by.String(),
	).Group("timeframe")

	g.db = dbQuery

	caller, err := g.ToTimeValue()
	if err != nil {
		return nil, err
	}

	if g.FillEmpty {
		return caller.FillMissing(g.Start, g.End)
	}
	return caller.ToValues()
}

func (g *GroupQuery) ToTimeValue() (*TimeVar, error) {
	rows, err := g.db.Rows()
	if err != nil {
		return nil, err
	}
	var data []*TimeValue
	for rows.Next() {
		var timeframe string
		amount := int64(0)
		if err := rows.Scan(&timeframe, &amount); err != nil {
			log.Error(err, timeframe)
		}
		trueTime, _ := g.db.ParseTime(timeframe)
		newTs := types.FixedTime(trueTime, g.Group)
		data = append(data, &TimeValue{
			Timeframe: newTs,
			Amount:    amount,
		})
	}
	return &TimeVar{g, data}, nil
}

func (t *TimeVar) FillMissing(current, end time.Time) ([]*TimeValue, error) {
	timeMap := make(map[string]int64)
	var validSet []*TimeValue
	dur := t.g.Group
	for _, v := range t.data {
		timeMap[v.Timeframe] = v.Amount
	}

	currentStr := types.FixedTime(current, t.g.Group)

	for {
		var amount int64
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
		currentStr = types.FixedTime(current, t.g.Group)
	}

	return validSet, nil
}

type isObject interface {
	Db() Database
}

func ParseQueries(r *http.Request, o isObject) *GroupQuery {
	fields := parseGet(r)
	grouping := fields.Get("group")
	startField := utils.ToInt(fields.Get("start"))
	endField := utils.ToInt(fields.Get("end"))
	limit := utils.ToInt(fields.Get("limit"))
	offset := utils.ToInt(fields.Get("offset"))
	fill, _ := strconv.ParseBool(fields.Get("fill"))
	orderBy := fields.Get("order")
	if limit == 0 {
		limit = 10000
	}

	db := o.Db()

	if grouping == "" {
		grouping = "1h"
	}
	groupDur, err := time.ParseDuration(grouping)
	if err != nil {
		log.Errorln(err)
		groupDur = 1 * time.Hour
	}

	query := &GroupQuery{
		Start:     time.Unix(startField, 0).UTC(),
		End:       time.Unix(endField, 0).UTC(),
		Group:     groupDur,
		Order:     orderBy,
		Limit:     int(limit),
		Offset:    int(offset),
		FillEmpty: fill,
		db:        db,
	}

	if startField == 0 {
		query.Start = time.Now().Add(-7 * types.Day).UTC()
	}
	if endField == 0 {
		query.End = time.Now().UTC()
	}
	if query.End.After(utils.Now()) {
		query.End = utils.Now()
	}

	if query.Limit != 0 {
		db = db.Limit(query.Limit)
	}
	if query.Offset > 0 {
		db = db.Offset(query.Offset)
	}

	db = db.Where("created_at BETWEEN ? AND ?", db.FormatTime(query.Start), db.FormatTime(query.End))

	if query.Order != "" {
		db = db.Order(query.Order)
	}
	query.db = db

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
