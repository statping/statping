package main

import (
	"github.com/ararog/timeago"
	"time"
)

type Failure struct {
	Id        int
	Issue     string
	Service   int
	CreatedAt time.Time
	Ago       string
}

func (s *Service) SelectAllFailures() []*Failure {
	var tks []*Failure
	rows, err := db.Query("SELECT * FROM failures WHERE service=$1 ORDER BY id DESC LIMIT 10", s.Id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tk Failure
		err = rows.Scan(&tk.Id, &tk.Issue, &tk.Service, &tk.CreatedAt)
		if err != nil {
			panic(err)
		}

		tk.Ago, _ = timeago.TimeAgoWithTime(time.Now(), tk.CreatedAt)

		tks = append(tks, &tk)
	}
	return tks
}

func CountFailures() int {
	var amount int
	db.QueryRow("SELECT COUNT(id) FROM failures;").Scan(&amount)
	return amount
}

func (s *Service) TotalFailures() int {
	var amount int
	db.QueryRow("SELECT COUNT(id) FROM failures WHERE service=$1;", s.Id).Scan(&amount)
	return amount
}

func (s *Service) TotalFailures24Hours() int {
	var amount int
	t := time.Now()
	x := t.AddDate(0, 0, -1)
	db.QueryRow("SELECT COUNT(id) FROM failures WHERE service=$1 AND created_at>=$2 AND created_at<$3;", s.Id, x, t).Scan(&amount)
	return amount
}
