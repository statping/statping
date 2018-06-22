package main

import (
	"fmt"
	"github.com/ararog/timeago"
	"time"
)

type Checkin struct {
	Id        int       `db:"id,omitempty"`
	Service   int64     `db:"service"`
	Interval  int64     `db:"check_interval"`
	Api       string    `db:"api"`
	CreatedAt time.Time `db:"created_at"`
	Hits      int64     `json:"hits"`
	Last      time.Time `json:"last"`
}

func (s *Service) SelectAllCheckins() []*Checkin {
	var checkins []*Checkin
	col := dbSession.Collection("checkins").Find("service", s.Id).OrderBy("-id")
	col.All(&checkins)
	s.Checkins = checkins
	return checkins
}

func (u *Checkin) Create() (int64, error) {
	u.CreatedAt = time.Now()
	uuid, err := dbSession.Collection("checkins").Insert(u)
	if uuid == nil {
		return 0, err
	}
	fmt.Println(uuid)
	return uuid.(int64), err
}

func SelectCheckinApi(api string) *Checkin {
	var checkin *Checkin
	dbSession.Collection("checkins").Find("api", api).One(&checkin)
	return checkin
}

func (c *Checkin) Receivehit() {
	c.Hits++
	c.Last = time.Now()
}

func (c *Checkin) RecheckCheckinFailure(guard chan struct{}) {
	between := time.Now().Sub(c.Last).Seconds()
	if between > float64(c.Interval) {
		fmt.Println("rechecking every 15 seconds!")
		c.CreateFailure()
		time.Sleep(15 * time.Second)
		guard <- struct{}{}
		c.RecheckCheckinFailure(guard)
	} else {
		fmt.Println("i recovered!!")
	}
	<-guard
}

func (f *Checkin) CreateFailure() {

}

func (f *Checkin) Ago() string {
	got, _ := timeago.TimeAgoWithTime(time.Now(), f.Last)
	return got
}

func FindCheckin(api string) *Checkin {
	for _, s := range services {
		for _, c := range s.Checkins {
			if c.Api == api {
				return c
			}
		}
	}
	return nil
}

func (c *Checkin) Run() {
	if c.Interval == 0 {
		return
	}
	fmt.Println("checking: ", c.Api)
	between := time.Now().Sub(c.Last).Seconds()
	if between > float64(c.Interval) {
		guard := make(chan struct{})
		c.RecheckCheckinFailure(guard)
		<-guard
	}
	time.Sleep(1 * time.Second)
	c.Run()
}

func (s *Service) StartCheckins() {
	for _, c := range s.Checkins {
		checkin := c
		go checkin.Run()
	}
}

func CheckinProcess() {
	for _, s := range services {
		for _, c := range s.Checkins {
			checkin := c
			go checkin.Run()
		}
	}
}
