package main

import "time"

var (
	Communications []*Communication
)

type Communication struct {
	Id        int64     `db:"id,omitempty" json:"id"`
	Method    string    `db:"method" json:"method"`
	Host      string    `db:"host" json:"host"`
	Port      int64     `db:"port" json:"port"`
	User      string    `db:"user" json:"user"`
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

func OnCommunicate() {
	for _, c := range Communications {
		if c.Enabled {
			c.Run()
		}
	}
}

func (c *Communication) Run() {

}

func SelectAllCommunications() ([]*Communication, error) {
	var c []*Communication
	col := dbSession.Collection("communication").Find()
	err := col.All(&c)
	Communications = c
	return c, err
}

func (c *Communication) Create() (int64, error) {
	c.CreatedAt = time.Now()
	uuid, err := dbSession.Collection("communication").Insert(c)
	if uuid == nil {
		return 0, err
	}
	c.Id = uuid.(int64)
	Communications = append(Communications, c)
	return uuid.(int64), err
}

func (c *Communication) Disable() {
	c.Enabled = false
	c.Update()
}

func (c *Communication) Enable() {
	c.Enabled = true
	c.Update()
}

func (c *Communication) Update() *Communication {
	col := dbSession.Collection("communication").Find("id", c.Id)
	col.Update(c)
	return c
}

func SelectCommunication(id int64) *Communication {
	for _, c := range Communications {
		if c.Id == id {
			return c
		}
	}
	return nil
}
