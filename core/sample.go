// Statping
// Copyright (C) 2018.  Hunter Long and the project contributors
// Written by Hunter Long <info@socialeck.com> and the project contributors
//
// https://github.com/hunterlong/statping
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
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"math/rand"
	"time"
)

// InsertSampleData will create the example/dummy services for a brand new Statping installation
func InsertSampleData() error {
	utils.Log(1, "Inserting Sample Data...")

	insertSampleGroups()

	s1 := ReturnService(&types.Service{
		Name:           "Google",
		Domain:         "https://google.com",
		ExpectedStatus: 200,
		Interval:       10,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          1,
		GroupId:        1,
	})
	s2 := ReturnService(&types.Service{
		Name:           "Statping Github",
		Domain:         "https://github.com/hunterlong/statping",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		Order:          2,
	})
	s3 := ReturnService(&types.Service{
		Name:           "JSON Users Test",
		Domain:         "https://jsonplaceholder.typicode.com/users",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        30,
		Order:          3,
		Public:         types.NewNullBool(true),
		GroupId:        2,
	})
	s4 := ReturnService(&types.Service{
		Name:           "JSON API Tester",
		Domain:         "https://jsonplaceholder.typicode.com/posts",
		ExpectedStatus: 201,
		Expected:       types.NewNullString(`(title)": "((\\"|[statup])*)"`),
		Interval:       30,
		Type:           "http",
		Method:         "POST",
		PostData:       types.NewNullString(`{ "title": "statup", "body": "bar", "userId": 19999 }`),
		Timeout:        30,
		Order:          4,
		Public:         types.NewNullBool(true),
		GroupId:        2,
	})
	s5 := ReturnService(&types.Service{
		Name:     "Google DNS",
		Domain:   "8.8.8.8",
		Interval: 20,
		Type:     "tcp",
		Port:     53,
		Timeout:  120,
		Order:    5,
		Public:   types.NewNullBool(true),
		GroupId:  1,
	})

	s1.Create(false)
	s2.Create(false)
	s3.Create(false)
	s4.Create(false)
	s5.Create(false)

	insertMessages()

	utils.Log(1, "Sample data has finished importing")

	return nil
}

func insertSampleGroups() error {
	group1 := &Group{&types.Group{
		Name: "Main Services",
	}}
	_, err := group1.Create()
	group2 := &Group{&types.Group{
		Name: "Linked Services",
	}}
	_, err = group2.Create()
	return err
}

// insertSampleCheckins will create 2 checkins with 60 successful hits per Checkin
func insertSampleCheckins() error {
	s1 := SelectService(1)
	checkin1 := ReturnCheckin(&types.Checkin{
		ServiceId:   s1.Id,
		Interval:    300,
		GracePeriod: 300,
	})
	checkin1.Update()

	s2 := SelectService(1)
	checkin2 := ReturnCheckin(&types.Checkin{
		ServiceId:   s2.Id,
		Interval:    900,
		GracePeriod: 300,
	})
	checkin2.Update()

	checkTime := time.Now().Add(-24 * time.Hour)
	for i := 0; i <= 60; i++ {
		checkHit := ReturnCheckinHit(&types.CheckinHit{
			Checkin:   checkin1.Id,
			From:      "192.168.0.1",
			CreatedAt: checkTime.UTC(),
		})
		checkHit.Create()
		checkTime = checkTime.Add(10 * time.Minute)
	}
	return nil
}

// InsertSampleHits will create a couple new hits for the sample services
func InsertSampleHits() error {
	since := time.Now().Add((-24 * 7) * time.Hour).UTC()
	for i := int64(1); i <= 5; i++ {
		service := SelectService(i)
		utils.Log(1, fmt.Sprintf("Adding %v sample hit records to service %v", 360, service.Name))
		createdAt := since
		alpha := float64(1.05)

		for hi := int64(1); hi <= 168; hi++ {
			alpha += 0.01
			rand.Seed(time.Now().UnixNano())
			latency := rand.Float64() * alpha
			createdAt = createdAt.Add(1 * time.Hour)
			hit := &types.Hit{
				Service:   service.Id,
				CreatedAt: createdAt,
				Latency:   latency,
			}
			service.CreateHit(hit)
		}
	}
	return nil
}

// insertSampleCore will create a new Core for the seed
func insertSampleCore() error {
	core := &types.Core{
		Name:        "Statping Sample Data",
		Description: "This data is only used to testing",
		ApiKey:      "sample",
		ApiSecret:   "samplesecret",
		Domain:      "http://localhost:8080",
		Version:     "test",
		CreatedAt:   time.Now(),
		UseCdn:      types.NewNullBool(false),
	}
	query := coreDB().Create(core)
	return query.Error
}

// insertSampleUsers will create 2 admin users for a seed database
func insertSampleUsers() error {
	u2 := ReturnUser(&types.User{
		Username: "testadmin",
		Password: "password123",
		Email:    "info@betatude.com",
		Admin:    types.NewNullBool(true),
	})

	u3 := ReturnUser(&types.User{
		Username: "testadmin2",
		Password: "password123",
		Email:    "info@adminhere.com",
		Admin:    types.NewNullBool(true),
	})

	_, err := u2.Create()
	_, err = u3.Create()
	return err
}

func insertMessages() error {
	m1 := ReturnMessage(&types.Message{
		Title:       "Routine Downtime",
		Description: "This is an example a upcoming message for a service!",
		ServiceId:   1,
		StartOn:     time.Now().Add(15 * time.Minute),
		EndOn:       time.Now().Add(2 * time.Hour),
	})
	if _, err := m1.Create(); err != nil {
		return err
	}
	m2 := ReturnMessage(&types.Message{
		Title:       "Server Reboot",
		Description: "This is another example a upcoming message for a service!",
		ServiceId:   3,
		StartOn:     time.Now().Add(15 * time.Minute),
		EndOn:       time.Now().Add(2 * time.Hour),
	})
	if _, err := m2.Create(); err != nil {
		return err
	}
	return nil
}

// InsertLargeSampleData will create the example/dummy services for testing the Statping server
func InsertLargeSampleData() error {
	if err := insertSampleCore(); err != nil {
		return err
	}
	if err := InsertSampleData(); err != nil {
		return err
	}
	if err := insertSampleUsers(); err != nil {
		return err
	}
	if err := insertSampleCheckins(); err != nil {
		return err
	}
	if err := insertMessages(); err != nil {
		return err
	}
	s6 := ReturnService(&types.Service{
		Name:           "JSON Lint",
		Domain:         "https://jsonlint.com",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          6,
	})

	s7 := ReturnService(&types.Service{
		Name:           "Demo Page",
		Domain:         "https://demo.statping.com",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        15,
		Order:          7,
	})

	s8 := ReturnService(&types.Service{
		Name:           "Golang",
		Domain:         "https://golang.org",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          8,
	})

	s9 := ReturnService(&types.Service{
		Name:           "Santa Monica",
		Domain:         "https://www.santamonica.com",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          9,
	})

	s10 := ReturnService(&types.Service{
		Name:           "Oeschs Die Dritten",
		Domain:         "https://www.oeschs-die-dritten.ch/en/",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          10,
	})

	s11 := ReturnService(&types.Service{
		Name:           "XS Project - Bochka, Bass, Kolbaser",
		Domain:         "https://www.youtube.com/watch?v=VLW1ieY4Izw",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		Order:          11,
	})

	s12 := ReturnService(&types.Service{
		Name:           "Github",
		Domain:         "https://github.com/hunterlong",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		Order:          12,
	})

	s13 := ReturnService(&types.Service{
		Name:           "Failing URL",
		Domain:         "http://thisdomainisfakeanditsgoingtofail.com",
		ExpectedStatus: 200,
		Interval:       45,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          13,
	})

	s14 := ReturnService(&types.Service{
		Name:           "Oesch's die Dritten - Die Jodelsprache",
		Domain:         "https://www.youtube.com/watch?v=k3GTxRt4iao",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        12,
		Order:          14,
	})

	s15 := ReturnService(&types.Service{
		Name:           "Gorm",
		Domain:         "http://gorm.io/",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        12,
		Order:          15,
	})

	s6.Create(false)
	s7.Create(false)
	s8.Create(false)
	s9.Create(false)
	s10.Create(false)
	s11.Create(false)
	s12.Create(false)
	s13.Create(false)
	s14.Create(false)
	s15.Create(false)

	var dayAgo = time.Now().Add(-24 * time.Hour).Add(-10 * time.Minute)

	insertHitRecords(dayAgo, 1450)

	insertFailureRecords(dayAgo, 730)

	return nil
}

// insertFailureRecords will create failures for 15 services from seed
func insertFailureRecords(since time.Time, amount int64) {
	for i := int64(14); i <= 15; i++ {
		service := SelectService(i)
		utils.Log(1, fmt.Sprintf("Adding %v Failure records to service %v", amount, service.Name))
		createdAt := since

		for fi := int64(1); fi <= amount; fi++ {
			createdAt = createdAt.Add(2 * time.Minute)

			failure := &Failure{&types.Failure{
				Service:   service.Id,
				Issue:     "testing right here",
				CreatedAt: createdAt,
			}}

			service.CreateFailure(failure)
		}
	}
}

// insertHitRecords will create successful Hit records for 15 services
func insertHitRecords(since time.Time, amount int64) {
	for i := int64(1); i <= 15; i++ {
		service := SelectService(i)
		utils.Log(1, fmt.Sprintf("Adding %v hit records to service %v", amount, service.Name))
		createdAt := since

		for hi := int64(1); hi <= amount; hi++ {
			rand.Seed(time.Now().UnixNano())
			latency := rand.Float64()
			createdAt = createdAt.Add(1 * time.Minute)
			hit := &types.Hit{
				Service:   service.Id,
				CreatedAt: createdAt,
				Latency:   latency,
			}
			service.CreateHit(hit)
		}

	}

}
