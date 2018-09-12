// Statup
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statup
//
// The licenses for most software and other practical works are designed
// to take away your freedom to share and change the works.  By contrast,
// the GNU General Public License is intended to guarantee your freedom to
// share and change all versions of a program--to make sure it remains free
// software for all its users.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"fmt"
	"github.com/ararog/timeago"
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"time"
)

type Checkin struct {
	*types.Checkin
}

func (c *Checkin) String() string {
	return c.Api
}

func ReturnCheckin(s *types.Checkin) *Checkin {
	return &Checkin{Checkin: s}
}

func FindCheckin(api string) *types.Checkin {
	for _, ser := range CoreApp.Services {
		service := ser.Select()
		for _, c := range service.Checkins {
			if c.Api == api {
				return c
			}
		}
	}
	return nil
}

func (s *Service) AllCheckins() []*types.Checkin {
	var checkins []*types.Checkin
	col := checkinDB().Where("service = ?", s.Id).Order("id desc")
	col.Find(&checkins)
	s.Checkins = checkins
	return checkins
}

func (u *Checkin) Create() (int64, error) {
	u.CreatedAt = time.Now()
	row := checkinDB().Create(u)
	if row.Error == nil {
		utils.Log(2, row.Error)
		return 0, row.Error
	}
	return u.Id, row.Error
}

func SelectCheckinApi(api string) *Checkin {
	var checkin *Checkin
	checkinDB().Where("api = ?", api).Find(&checkin)
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
//	for _, s := range CoreApp.DbServices {
//		for _, c := range s.Checkins {
//			checkin := c
//			go checkin.Run()
//		}
//	}
//}
