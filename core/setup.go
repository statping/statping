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
	"github.com/hunterlong/statup/types"
	"github.com/hunterlong/statup/utils"
	"os"
)

// DeleteConfig will delete the 'config.yml' file
func DeleteConfig() {
	err := os.Remove(utils.Directory + "/config.yml")
	if err != nil {
		utils.Log(3, err)
	}
}

type ErrorResponse struct {
	Error string
}

// LoadSampleData will create the example/dummy services for a brand new Statup installation
func LoadSampleData() error {
	utils.Log(1, "Inserting Sample Data...")
	s1 := ReturnService(&types.Service{
		Name:           "Google",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       10,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
	})
	s2 := ReturnService(&types.Service{
		Name:           "Statup Github",
		Domain:         "https://github.com/hunterlong/statup",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
	})
	s3 := ReturnService(&types.Service{
		Name:           "JSON Users Test",
		Domain:         "https://jsonplaceholder.typicode.com/users",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
	})
	s4 := ReturnService(&types.Service{
		Name:           "JSON API Tester",
		Domain:         "https://jsonplaceholder.typicode.com/posts",
		ExpectedStatus: 201,
		Expected:       `(title)": "((\\"|[statup])*)"`,
		Interval:       30,
		Type:           "http",
		Method:         "POST",
		PostData:       `{ "title": "statup", "body": "bar", "userId": 19999 }`,
		Timeout:        30,
	})
	s5 := ReturnService(&types.Service{
		Name:     "Google DNS",
		Domain:   "8.8.8.8",
		Interval: 20,
		Type:     "tcp",
		Port:     53,
		Timeout:  120,
	})
	id, err := s1.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = s2.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = s3.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = s4.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating Service %v: %v", id, err))
	}
	id, err = s5.Create()
	if err != nil {
		utils.Log(3, fmt.Sprintf("Error creating TCP Service %v: %v", id, err))
	}

	//checkin := &Checkin{
	//	Service:  s2.Id,
	//	Interval: 30,
	//	Api:      utils.NewSHA1Hash(18),
	//}
	//id, err = checkin.Create()
	//if err != nil {
	//	utils.Log(3, fmt.Sprintf("Error creating Checkin %v: %v", id, err))
	//}

	//for i := 0; i < 3; i++ {
	//	s1.Check()
	//	s2.Check()
	//	s3.Check()
	//	s4.Check()
	//}

	utils.Log(1, "Sample data has finished importing")

	return nil
}
