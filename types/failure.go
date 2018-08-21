package types

import "time"

type Failure struct {
	Id               int       `db:"id,omitempty" json:"id"`
	Issue            string    `db:"issue" json:"issue"`
	Method           string    `db:"method" json:"method,omitempty"`
	Service          int64     `db:"service" json:"service_id"`
	CreatedAt        time.Time `db:"created_at" json:"created_at"`
	FailureInterface `json:"-"`
}

type FailureInterface interface {
	Delete() error
	// method functions
	Ago() string
	ParseError() string
}
