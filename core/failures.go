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
	"strings"
	"time"
)

func CreateServiceFailure(s *Service, data FailureData) (int64, error) {
	fail := &types.Failure{
		Issue:     data.Issue,
		Service:   s.Id,
		CreatedAt: time.Now(),
	}
	s.Failures = append(s.Failures, fail)
	col := DbSession.Collection("failures")
	uuid, err := col.Insert(fail)
	if err != nil {
		utils.Log(3, err)
	}
	if uuid == nil {
		return 0, err
	}
	return uuid.(int64), err
}

func SelectAllFailures(s *types.Service) []*types.Failure {
	var fails []*types.Failure
	col := DbSession.Collection("failures").Find("service", s.Id).OrderBy("-id")
	err := col.All(&fails)
	if err != nil {
		utils.Log(3, fmt.Sprintf("Issue getting failures for service %v, %v", s.Name, err))
	}
	return fails
}

func DeleteFailures(u *Service) {
	var fails []*Failure
	col := DbSession.Collection("failures")
	col.Find("service", u.Id).All(&fails)
	for _, fail := range fails {
		fail.Delete()
	}
}

func (s *Service) LimitedFailures() []*Failure {
	var fails []*types.Failure
	var failArr []*Failure
	col := DbSession.Collection("failures").Find("service", s.Id).OrderBy("-id").Limit(10)
	col.All(&fails)
	for _, f := range fails {
		failArr = append(failArr, MakeFailure(f))
	}
	return failArr
}

func reverseFailures(input []*types.Failure) []*types.Failure {
	if len(input) == 0 {
		return input
	}
	return append(reverseFailures(input[1:]), input[0])
}

func (fail *Failure) Ago() string {
	f := fail.ToFailure()
	got, _ := timeago.TimeAgoWithTime(time.Now(), f.CreatedAt)
	return got
}

func (fail *Failure) Delete() error {
	f := fail.ToFailure()
	col := DbSession.Collection("failures").Find("id", f.Id)
	return col.Delete()
}

func CountFailures() uint64 {
	col := DbSession.Collection("failures").Find()
	amount, err := col.Count()
	if err != nil {
		utils.Log(2, err)
		return 0
	}
	return amount
}

func (s *Service) TotalFailures() (uint64, error) {
	col := DbSession.Collection("failures").Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (s *Service) TotalFailures24Hours() (uint64, error) {
	col := DbSession.Collection("failures").Find("service", s.Id)
	amount, err := col.Count()
	return amount, err
}

func (f *Failure) ToFailure() *types.Failure {
	return f.F.(*types.Failure)
}

func MakeFailure(f *types.Failure) *Failure {
	fail := &Failure{f}
	return fail
}

func (fail *Failure) ParseError() string {
	f := fail.ToFailure()
	err := strings.Contains(f.Issue, "operation timed out")
	if err {
		return fmt.Sprintf("HTTP Request Timed Out")
	}
	err = strings.Contains(f.Issue, "x509: certificate is valid")
	if err {
		return fmt.Sprintf("SSL Certificate invalid")
	}
	err = strings.Contains(f.Issue, "Client.Timeout exceeded while awaiting headers")
	if err {
		return fmt.Sprintf("Connection Timed Out")
	}
	err = strings.Contains(f.Issue, "no such host")
	if err {
		return fmt.Sprintf("Domain is offline or not found")
	}
	err = strings.Contains(f.Issue, "HTTP Status Code")
	if err {
		return fmt.Sprintf("Incorrect HTTP Status Code")
	}
	err = strings.Contains(f.Issue, "connection refused")
	if err {
		return fmt.Sprintf("Connection Failed")
	}
	err = strings.Contains(f.Issue, "can't assign requested address")
	if err {
		return fmt.Sprintf("Unable to Request Address")
	}
	err = strings.Contains(f.Issue, "no route to host")
	if err {
		return fmt.Sprintf("Domain is offline or not found")
	}
	return f.Issue
}
