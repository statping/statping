package core

import (
	"fmt"
	"github.com/ararog/timeago"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"time"
)

type Checkin types.Checkin

func (c *Checkin) String() string {
	return c.Api
}

func FindCheckin(api string) *types.Checkin {
	for _, ser := range CoreApp.Services {
		s := ser.ToService()
		for _, c := range s.Checkins {
			if c.Api == api {
				return c
			}
		}
	}
	return nil
}

func SelectAllCheckins(s *types.Service) []*types.Checkin {
	var checkins []*types.Checkin
	col := DbSession.Collection("checkins").Find("service", s.Id).OrderBy("-id")
	col.All(&checkins)
	s.Checkins = checkins
	return checkins
}

func (u *Checkin) Create() (int64, error) {
	u.CreatedAt = time.Now()
	uuid, err := DbSession.Collection("checkins").Insert(u)
	if uuid == nil {
		utils.Log(2, err)
		return 0, err
	}
	fmt.Println("new checkin: ", uuid)
	return uuid.(int64), err
}

func SelectCheckinApi(api string) *Checkin {
	var checkin *Checkin
	DbSession.Collection("checkins").Find("api", api).One(&checkin)
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

//func (c *Checkin) Run() {
//	if c.Interval == 0 {
//		return
//	}
//	fmt.Println("checking: ", c.Api)
//	between := time.Now().Sub(c.Last).Seconds()
//	if between > float64(c.Interval) {
//		guard := make(chan struct{})
//		c.RecheckCheckinFailure(guard)
//		<-guard
//	}
//	time.Sleep(1 * time.Second)
//	c.Run()
//}
//
//func (s *Service) StartCheckins() {
//	for _, c := range s.Checkins {
//		checkin := c.(*Checkin)
//		go checkin.Run()
//	}
//}
//
//func CheckinProcess() {
//	for _, s := range CoreApp.Services {
//		for _, c := range s.Checkins {
//			checkin := c
//			go checkin.Run()
//		}
//	}
//}
