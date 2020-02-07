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
	"github.com/hunterlong/statping/core/notifier"
	"github.com/hunterlong/statping/types"
	"github.com/hunterlong/statping/utils"
	"sync"
	"time"
)

var (
	sampleStart = time.Now().Add((-24 * 7) * time.Hour).UTC()
	SampleHits  = 9900.
)

// InsertSampleData will create the example/dummy services for a brand new Statping installation
func InsertSampleData() error {
	log.Infoln("Inserting Sample Data...")

	insertSampleGroups()
	createdOn := time.Now().Add(((-24 * 30) * 3) * time.Hour).UTC()
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
		Permalink:      types.NewNullString("google"),
		VerifySSL:      types.NewNullBool(true),
		CreatedAt:      createdOn,
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
		Permalink:      types.NewNullString("statping_github"),
		VerifySSL:      types.NewNullBool(true),
		CreatedAt:      createdOn,
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
		VerifySSL:      types.NewNullBool(true),
		GroupId:        2,
		CreatedAt:      createdOn,
	})
	s4 := ReturnService(&types.Service{
		Name:           "JSON API Tester",
		Domain:         "https://jsonplaceholder.typicode.com/posts",
		ExpectedStatus: 201,
		Expected:       types.NewNullString(`(title)": "((\\"|[statping])*)"`),
		Interval:       30,
		Type:           "http",
		Method:         "POST",
		PostData:       types.NewNullString(`{ "title": "statping", "body": "bar", "userId": 19999 }`),
		Timeout:        30,
		Order:          4,
		Public:         types.NewNullBool(true),
		VerifySSL:      types.NewNullBool(true),
		GroupId:        2,
		CreatedAt:      createdOn,
	})
	s5 := ReturnService(&types.Service{
		Name:      "Google DNS",
		Domain:    "8.8.8.8",
		Interval:  20,
		Type:      "tcp",
		Port:      53,
		Timeout:   120,
		Order:     5,
		Public:    types.NewNullBool(true),
		GroupId:   1,
		CreatedAt: createdOn,
	})

	s1.Create(false)
	s2.Create(false)
	s3.Create(false)
	s4.Create(false)
	s5.Create(false)

	insertMessages()

	insertSampleIncidents()

	log.Infoln("Sample data has finished importing")

	return nil
}

func insertSampleIncidents() error {
	incident1 := &Incident{&types.Incident{
		Title:       "Github Downtime",
		Description: "This is an example of a incident for a service.",
		ServiceId:   2,
	}}
	_, err := incident1.Create()

	incidentUpdate1 := &IncidentUpdate{&types.IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github's page for Statping seems to be sending a 501 error.",
		Type:       "Investigating",
	}}
	_, err = incidentUpdate1.Create()

	incidentUpdate2 := &IncidentUpdate{&types.IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Problem is continuing and we are looking at the issues.",
		Type:       "Update",
	}}
	_, err = incidentUpdate2.Create()

	incidentUpdate3 := &IncidentUpdate{&types.IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}}
	_, err = incidentUpdate3.Create()

	return err
}

func insertSampleGroups() error {
	group1 := &Group{&types.Group{
		Name:   "Main Services",
		Public: types.NewNullBool(true),
		Order:  2,
	}}
	_, err := group1.Create()
	group2 := &Group{&types.Group{
		Name:   "Linked Services",
		Public: types.NewNullBool(false),
		Order:  1,
	}}
	_, err = group2.Create()
	group3 := &Group{&types.Group{
		Name:   "Empty Group",
		Public: types.NewNullBool(false),
		Order:  3,
	}}
	_, err = group3.Create()
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

	checkTime := time.Now().UTC().Add(-24 * time.Hour)
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
	tx := hitsDB().Begin()
	sg := new(sync.WaitGroup)
	for i := int64(1); i <= 5; i++ {
		sg.Add(1)
		service := SelectService(i)
		seed := time.Now().UnixNano()
		log.Infoln(fmt.Sprintf("Adding %v sample hit records to service %v", SampleHits, service.Name))
		createdAt := sampleStart
		p := utils.NewPerlin(2., 2., 10, seed)
		go func() {
			defer sg.Done()
			for hi := 0.; hi <= float64(SampleHits); hi++ {
				latency := p.Noise1D(hi / 500)
				createdAt = createdAt.Add(60 * time.Second)
				hit := &types.Hit{
					Service:   service.Id,
					CreatedAt: createdAt,
					Latency:   latency,
				}
				tx = tx.Create(&hit)
			}
		}()
	}
	sg.Wait()
	err := tx.Commit().Error
	if err != nil {
		log.Errorln(err)
	}
	return err
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
		CreatedAt:   time.Now().UTC(),
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
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
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
	createdOn := time.Now().UTC().Add((-24 * 90) * time.Hour)
	s6 := ReturnService(&types.Service{
		Name:           "JSON Lint",
		Domain:         "https://jsonlint.com",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          6,
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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
		CreatedAt:      createdOn,
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

	var dayAgo = time.Now().UTC().Add((-24 * 90) * time.Hour)

	insertHitRecords(dayAgo, 5450)

	insertFailureRecords(dayAgo, 730)

	return nil
}

// insertFailureRecords will create failures for 15 services from seed
func insertFailureRecords(since time.Time, amount int64) {
	for i := int64(14); i <= 15; i++ {
		service := SelectService(i)
		log.Infoln(fmt.Sprintf("Adding %v Failure records to service %v", amount, service.Name))
		createdAt := since

		for fi := int64(1); fi <= amount; fi++ {
			createdAt = createdAt.Add(2 * time.Minute)

			failure := &types.Failure{
				Service:   service.Id,
				Issue:     "testing right here",
				CreatedAt: createdAt,
			}

			service.CreateFailure(failure)
		}
	}
}

// insertHitRecords will create successful Hit records for 15 services
func insertHitRecords(since time.Time, amount int64) {
	for i := int64(1); i <= 15; i++ {
		service := SelectService(i)
		log.Infoln(fmt.Sprintf("Adding %v hit records to service %v", amount, service.Name))
		createdAt := since
		p := utils.NewPerlin(2, 2, 5, time.Now().UnixNano())
		for hi := int64(1); hi <= amount; hi++ {
			latency := p.Noise1D(float64(hi / 10))
			createdAt = createdAt.Add(1 * time.Minute)
			hit := &types.Hit{
				Service:   service.Id,
				CreatedAt: createdAt.UTC(),
				Latency:   latency,
			}
			service.CreateHit(hit)
		}

	}

}

// TmpRecords is used for testing Statping. It will create a SQLite database file
// with sample data and store it in the /tmp folder to be used by the tests.
func TmpRecords(dbFile string) error {
	var sqlFile = utils.Directory + "/" + dbFile
	utils.CreateDirectory(utils.Directory + "/tmp")
	var tmpSqlFile = utils.Directory + "/tmp/" + types.SqliteFilename
	SampleHits = 480

	var err error
	CoreApp = NewCore()
	CoreApp.Name = "Tester"
	configs := &types.DbConfig{
		DbConn:   "sqlite",
		Project:  "Tester",
		Location: utils.Directory,
		SqlFile:  sqlFile,
	}
	log.Infoln("saving config.yml in: " + utils.Directory)
	if configs, err = CoreApp.SaveConfig(configs); err != nil {
		return err
	}
	log.Infoln("loading config.yml from: " + utils.Directory)
	if configs, err = LoadConfigFile(utils.Directory); err != nil {
		return err
	}
	log.Infoln("connecting to database")

	exists := utils.FileExists(tmpSqlFile)
	if exists {
		log.Infoln(tmpSqlFile + " was found, copying the temp database to " + sqlFile)
		if err := utils.DeleteFile(sqlFile); err != nil {
			log.Infoln(sqlFile + " was not found")
		}
		if err := utils.CopyFile(tmpSqlFile, sqlFile); err != nil {
			return err
		}
		log.Infoln("loading config.yml from: " + utils.Directory)

		if err := CoreApp.Connect(false, utils.Directory); err != nil {
			return err
		}
		log.Infoln("selecting the Core variable")
		if _, err := SelectCore(); err != nil {
			return err
		}
		log.Infoln("inserting notifiers into database")
		if err := InsertNotifierDB(); err != nil {
			return err
		}
		log.Infoln("loading all services")
		if _, err := CoreApp.SelectAllServices(false); err != nil {
			return err
		}
		if err := AttachNotifiers(); err != nil {
			return err
		}
		CoreApp.Notifications = notifier.AllCommunications
		return nil
	}

	log.Infoln(tmpSqlFile + " not found, creating a new database...")

	if err := CoreApp.Connect(false, utils.Directory); err != nil {
		return err
	}
	log.Infoln("creating database")
	if err := CoreApp.CreateDatabase(); err != nil {
		return err
	}
	log.Infoln("migrating database")
	if err := CoreApp.MigrateDatabase(); err != nil {
		return err
	}
	log.Infoln("insert large sample data into database")
	if err := InsertLargeSampleData(); err != nil {
		return err
	}
	log.Infoln("selecting the Core variable")
	if CoreApp, err = SelectCore(); err != nil {
		return err
	}
	log.Infoln("inserting notifiers into database")
	if err := InsertNotifierDB(); err != nil {
		return err
	}
	log.Infoln("loading all services")
	if _, err := CoreApp.SelectAllServices(false); err != nil {
		return err
	}
	log.Infoln("copying sql database file to: " + tmpSqlFile)
	if err := utils.CopyFile(sqlFile, tmpSqlFile); err != nil {
		return err
	}
	return err
}
