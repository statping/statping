package types

import "time"

type Checkin struct {
	Id        int       `db:"id,omitempty"`
	Service   int64     `db:"service"`
	Interval  int64     `db:"check_interval"`
	Api       string    `db:"api"`
	CreatedAt time.Time `db:"created_at"`
	Hits      int64     `json:"hits"`
	Last      time.Time `json:"last"`
	CheckinInterface
}

type CheckinInterface interface {
	// Database functions
	Create() (int64, error)
	//Update() error
	//Delete() error
	Ago() string
	Receivehit()
}
