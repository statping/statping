package types

import "time"

type User struct {
	Id            int64     `db:"id,omitempty" json:"id"`
	Username      string    `db:"username" json:"username"`
	Password      string    `db:"password" json:"-"`
	Email         string    `db:"email" json:"-"`
	ApiKey        string    `db:"api_key" json:"api_key"`
	ApiSecret     string    `db:"api_secret" json:"-"`
	Admin         bool      `db:"administrator" json:"admin"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UserInterface `json:"-"`
}

type UserInterface interface {
	// Database functions
	Create() (int64, error)
	Update() error
	Delete() error
}
