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
	"github.com/hunterlong/statping/database"
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
	s1 := &types.Service{
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
	}
	s2 := &types.Service{
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
	}
	s3 := &types.Service{
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
	}
	s4 := &types.Service{
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
	}
	s5 := &types.Service{
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
	}

	database.Create(s1)
	database.Create(s2)
	database.Create(s3)
	database.Create(s4)
	database.Create(s5)

	insertMessages()

	insertSampleIncidents()

	log.Infoln("Sample data has finished importing")

	return nil
}

func insertSampleIncidents() error {
	incident1 := &types.Incident{
		Title:       "Github Downtime",
		Description: "This is an example of a incident for a service.",
		ServiceId:   2,
	}
	if _, err := database.Create(incident1); err != nil {
		return err
	}

	incidentUpdate1 := &types.IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github's page for Statping seems to be sending a 501 error.",
		Type:       "Investigating",
	}
	if _, err := database.Create(incidentUpdate1); err != nil {
		return err
	}

	incidentUpdate2 := &types.IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Problem is continuing and we are looking at the issues.",
		Type:       "Update",
	}
	if _, err := database.Create(incidentUpdate2); err != nil {
		return err
	}

	incidentUpdate3 := &types.IncidentUpdate{
		IncidentId: incident1.Id,
		Message:    "Github is now back online and everything is working.",
		Type:       "Resolved",
	}
	if _, err := database.Create(incidentUpdate3); err != nil {
		return err
	}

	return nil
}

func insertSampleGroups() error {
	group1 := &types.Group{
		Name:   "Main Services",
		Public: types.NewNullBool(true),
		Order:  2,
	}
	if _, err := database.Create(group1); err != nil {
		return err
	}

	group2 := &types.Group{
		Name:   "Linked Services",
		Public: types.NewNullBool(false),
		Order:  1,
	}
	if _, err := database.Create(group2); err != nil {
		return err
	}

	group3 := &types.Group{
		Name:   "Empty Group",
		Public: types.NewNullBool(false),
		Order:  3,
	}
	if _, err := database.Create(group3); err != nil {
		return err
	}
	return nil
}

// insertSampleCheckins will create 2 checkins with 60 successful hits per Checkin
func insertSampleCheckins() error {
	s1 := SelectService(1)
	checkin1 := &types.Checkin{
		ServiceId:   s1.Id,
		Interval:    300,
		GracePeriod: 300,
	}

	if _, err := database.Create(checkin1); err != nil {
		return err
	}

	s2 := SelectService(1)
	checkin2 := &types.Checkin{
		ServiceId:   s2.Id,
		Interval:    900,
		GracePeriod: 300,
	}

	if _, err := database.Create(checkin2); err != nil {
		return err
	}

	checkTime := time.Now().UTC().Add(-24 * time.Hour)
	for i := 0; i <= 60; i++ {
		checkHit := &types.CheckinHit{
			Checkin:   checkin1.Id,
			From:      "192.168.0.1",
			CreatedAt: checkTime.UTC(),
		}

		if _, err := database.Create(checkHit); err != nil {
			return err
		}

		checkTime = checkTime.Add(10 * time.Minute)
	}
	return nil
}

// InsertSampleHits will create a couple new hits for the sample services
func InsertSampleHits() error {
	tx := Database(&Hit{}).Begin()
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
	err := tx.Commit().Error()
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

	_, err := database.Create(core)

	return err
}

// insertSampleUsers will create 2 admin users for a seed database
func insertSampleUsers() error {
	u2 := &types.User{
		Username: "testadmin",
		Password: "password123",
		Email:    "info@betatude.com",
		Admin:    types.NewNullBool(true),
	}

	if _, err := database.Create(u2); err != nil {
		return err
	}

	u3 := &types.User{
		Username: "testadmin2",
		Password: "password123",
		Email:    "info@adminhere.com",
		Admin:    types.NewNullBool(true),
	}

	if _, err := database.Create(u3); err != nil {
		return err
	}

	return nil
}

func insertMessages() error {
	m1 := &types.Message{
		Title:       "Routine Downtime",
		Description: "This is an example a upcoming message for a service!",
		ServiceId:   1,
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
	}

	if _, err := database.Create(m1); err != nil {
		return err
	}

	m2 := &types.Message{
		Title:       "Server Reboot",
		Description: "This is another example a upcoming message for a service!",
		ServiceId:   3,
		StartOn:     time.Now().UTC().Add(15 * time.Minute),
		EndOn:       time.Now().UTC().Add(2 * time.Hour),
	}

	if _, err := database.Create(m2); err != nil {
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
	s6 := &types.Service{
		Name:           "JSON Lint",
		Domain:         "https://jsonlint.com",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          6,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s6); err != nil {
		return err
	}

	s7 := &types.Service{
		Name:           "Demo Page",
		Domain:         "https://demo.statping.com",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        15,
		Order:          7,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s7); err != nil {
		return err
	}

	s8 := &types.Service{
		Name:           "Golang",
		Domain:         "https://golang.org",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          8,
	}

	if _, err := database.Create(s8); err != nil {
		return err
	}

	s9 := &types.Service{
		Name:           "Santa Monica",
		Domain:         "https://www.santamonica.com",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          9,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s9); err != nil {
		return err
	}

	s10 := &types.Service{
		Name:           "Oeschs Die Dritten",
		Domain:         "https://www.oeschs-die-dritten.ch/en/",
		ExpectedStatus: 200,
		Interval:       15,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          10,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s10); err != nil {
		return err
	}

	s11 := &types.Service{
		Name:           "XS Project - Bochka, Bass, Kolbaser",
		Domain:         "https://www.youtube.com/watch?v=VLW1ieY4Izw",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		Order:          11,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s11); err != nil {
		return err
	}

	s12 := &types.Service{
		Name:           "Github",
		Domain:         "https://github.com/hunterlong",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        20,
		Order:          12,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s12); err != nil {
		return err
	}

	s13 := &types.Service{
		Name:           "Failing URL",
		Domain:         "http://thisdomainisfakeanditsgoingtofail.com",
		ExpectedStatus: 200,
		Interval:       45,
		Type:           "http",
		Method:         "GET",
		Timeout:        10,
		Order:          13,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s13); err != nil {
		return err
	}

	s14 := &types.Service{
		Name:           "Oesch's die Dritten - Die Jodelsprache",
		Domain:         "https://www.youtube.com/watch?v=k3GTxRt4iao",
		ExpectedStatus: 200,
		Interval:       60,
		Type:           "http",
		Method:         "GET",
		Timeout:        12,
		Order:          14,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s14); err != nil {
		return err
	}

	s15 := &types.Service{
		Name:           "Gorm",
		Domain:         "http://gorm.io/",
		ExpectedStatus: 200,
		Interval:       30,
		Type:           "http",
		Method:         "GET",
		Timeout:        12,
		Order:          15,
		CreatedAt:      createdOn,
	}

	if _, err := database.Create(s15); err != nil {
		return err
	}

	var dayAgo = time.Now().UTC().Add((-24 * 90) * time.Hour)

	insertHitRecords(dayAgo, 5450)

	insertFailureRecords(dayAgo, 730)

	return nil
}

// insertFailureRecords will create failures for 15 services from seed
func insertFailureRecords(since time.Time, amount int) {
	for i := int64(14); i <= 15; i++ {
		service := SelectService(i)
		log.Infoln(fmt.Sprintf("Adding %v Failure records to service %v", amount, service.Name))
		createdAt := since

		for fi := 1; fi <= amount; fi++ {
			createdAt = createdAt.Add(2 * time.Minute)

			failure := &types.Failure{
				Service:   service.Id,
				Issue:     "testing right here",
				CreatedAt: createdAt,
			}

			database.Create(failure)
		}
	}
}

// insertHitRecords will create successful Hit records for 15 services
func insertHitRecords(since time.Time, amount int) {
	for i := int64(1); i <= 15; i++ {
		service := SelectService(i)
		log.Infoln(fmt.Sprintf("Adding %v hit records to service %v", amount, service.Name))
		createdAt := since
		p := utils.NewPerlin(2, 2, 5, time.Now().UnixNano())
		for hi := 1; hi <= amount; hi++ {
			latency := p.Noise1D(float64(hi / 10))
			createdAt = createdAt.Add(1 * time.Minute)
			hit := &types.Hit{
				Service:   service.Id,
				CreatedAt: createdAt.UTC(),
				Latency:   latency,
			}
			database.Create(hit)
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
	CoreApp.Setup = true
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
		log.Infoln("inserting integrations into database")
		if err := InsertIntegratorDB(); err != nil {
			return err
		}
		log.Infoln("loading all services")
		if _, err := SelectAllServices(false); err != nil {
			return err
		}
		if err := AttachNotifiers(); err != nil {
			return err
		}
		if err := AddIntegrations(); err != nil {
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
	log.Infoln("inserting integrations into database")
	if err := InsertIntegratorDB(); err != nil {
		return err
	}
	log.Infoln("loading all services")
	if _, err := SelectAllServices(false); err != nil {
		return err
	}
	log.Infoln("copying sql database file to: " + tmpSqlFile)
	if err := utils.CopyFile(sqlFile, tmpSqlFile); err != nil {
		return err
	}
	return err
}
