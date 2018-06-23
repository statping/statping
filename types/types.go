package types

import "time"

type Communication struct {
	Id        int64     `db:"id,omitempty" json:"id"`
	Method    string    `db:"method" json:"method"`
	Host      string    `db:"host" json:"host"`
	Port      int       `db:"port" json:"port"`
	Username  string    `db:"username" json:"user"`
	Password  string    `db:"password" json:"-"`
	Var1      string    `db:"var1" json:"var1"`
	Var2      string    `db:"var2" json:"var2"`
	ApiKey    string    `db:"api_key" json:"api_key"`
	ApiSecret string    `db:"api_secret" json:"api_secret"`
	Enabled   bool      `db:"enabled" json:"enabled"`
	Limits    int64     `db:"limits" json:"limits"`
	Removable bool      `db:"removable" json:"removable"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type Email struct {
	To       string
	Subject  string
	Template string
	Data     interface{}
}
