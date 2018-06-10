package main

import "time"

type Failure struct {
	Id        int
	Issue     string
	Service   int
	CreatedAt time.Time
}

func SelectAllFailures(id int64) []float64 {
	var tks []float64
	rows, err := db.Query("SELECT * FROM failures WHERE service=$1 ORDER BY id ASC", id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var tk Hit
		err = rows.Scan(&tk.Id, &tk.Metric, &tk.Value, &tk.CreatedAt)
		if err != nil {
			panic(err)
		}
		tks = append(tks, tk.Value)
	}
	return tks
}

func (s *Service) TotalFailures() int {
	var amount int
	db.QueryRow("SELECT COUNT(id) FROM failures WHERE service=$1;", s.Id).Scan(&amount)
	return amount
}
