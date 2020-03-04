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

package database

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	sampleStart = time.Now().Add((-24 * 7) * time.Hour).UTC()
	SampleHits  = 9900.
)

// InsertSampleHits will create a couple new hits for the sample services
func InsertSampleHits() error {
	//tx := Begin(&hits.Hit{})
	//sg := new(sync.WaitGroup)
	//for i := int64(1); i <= 5; i++ {
	//	sg.Add(1)
	//	service := SelectService(i)
	//	seed := time.Now().UnixNano()
	//	log.Infoln(fmt.Sprintf("Adding %v sample hit records to service %v", SampleHits, service.Name))
	//	createdAt := sampleStart
	//	p := utils.NewPerlin(2., 2., 10, seed)
	//	go func(sg *sync.WaitGroup) {
	//		defer sg.Done()
	//		for hi := 0.; hi <= float64(SampleHits); hi++ {
	//			latency := p.Noise1D(hi / 500)
	//			createdAt = createdAt.Add(60 * time.Second)
	//			hit := &hits.Hit{
	//				Service:   service.Id,
	//				CreatedAt: createdAt,
	//				Latency:   latency,
	//			}
	//			tx = tx.Create(&hit)
	//		}
	//	}(sg)
	//}
	//sg.Wait()
	//if err := tx.Commit().Error(); err != nil {
	//	log.Errorln(err)
	//	return types.ErrWrap(err, types.ErrorCreateSampleHits)
	//}
	return nil
}

// TmpRecords is used for testing Statping. It will create a SQLite database file
// with sample data and store it in the /tmp folder to be used by the tests.
//func TmpRecords(dbFile string) error {
//	var sqlFile = utils.Directory + "/" + dbFile
//	if err := utils.CreateDirectory(utils.Directory + "/tmp"); err != nil {
//		log.Error(err)
//	}
//	var tmpSqlFile = utils.Directory + "/tmp/" + types.SqliteFilename
//	SampleHits = 480
//
//	var err error
//	CoreApp = NewCore()
//	CoreApp.Name = "Tester"
//	CoreApp.Setup = true
//	configs := &types.DbConfig{
//		DbConn:   "sqlite",
//		Project:  "Tester",
//		Location: utils.Directory,
//		SqlFile:  sqlFile,
//	}
//	log.Infoln("saving config.yml in: " + utils.Directory)
//	if err := configs.Save(utils.Directory); err != nil {
//		log.Error(err)
//	}
//
//	log.Infoln("loading config.yml from: " + utils.Directory)
//	if configs, err = LoadConfigs(); err != nil {
//		log.Error(err)
//	}
//	log.Infoln("connecting to database")
//
//	exists := utils.FileExists(tmpSqlFile)
//	if exists {
//		log.Infoln(tmpSqlFile + " was found, copying the temp database to " + sqlFile)
//		if err := utils.DeleteFile(sqlFile); err != nil {
//			log.Error(err)
//		}
//		if err := utils.CopyFile(tmpSqlFile, sqlFile); err != nil {
//			log.Error(err)
//		}
//		log.Infoln("loading config.yml from: " + utils.Directory)
//
//		if err := CoreApp.Connect(configs, false, utils.Directory); err != nil {
//			log.Error(err)
//		}
//		log.Infoln("selecting the Core variable")
//		if _, err := SelectCore(); err != nil {
//			log.Error(err)
//		}
//		log.Infoln("inserting notifiers into database")
//		if err := InsertNotifierDB(); err != nil {
//			log.Error(err)
//		}
//		log.Infoln("inserting integrations into database")
//		if err := InsertIntegratorDB(); err != nil {
//			log.Error(err)
//		}
//		log.Infoln("loading all services")
//		if _, err := SelectAllServices(false); err != nil {
//			return err
//		}
//		if err := AttachNotifiers(); err != nil {
//			log.Error(err)
//		}
//		if err := AddIntegrations(); err != nil {
//			log.Error(err)
//		}
//		CoreApp.Notifications = notifier.AllCommunications
//		return nil
//	}
//
//	log.Infoln(tmpSqlFile + " not found, creating a new database...")
//
//	if err := CoreApp.Connect(configs, false, utils.Directory); err != nil {
//		return err
//	}
//	log.Infoln("creating database")
//	if err := CoreApp.CreateDatabase(); err != nil {
//		return err
//	}
//	log.Infoln("migrating database")
//	if err := MigrateDatabase(); err != nil {
//		return err
//	}
//	log.Infoln("insert large sample data into database")
//	if err := InsertLargeSampleData(); err != nil {
//		return err
//	}
//	log.Infoln("selecting the Core variable")
//	if CoreApp, err = SelectCore(); err != nil {
//		return err
//	}
//	log.Infoln("inserting notifiers into database")
//	if err := InsertNotifierDB(); err != nil {
//		return err
//	}
//	log.Infoln("inserting integrations into database")
//	if err := InsertIntegratorDB(); err != nil {
//		return err
//	}
//	log.Infoln("loading all services")
//	if _, err := SelectAllServices(false); err != nil {
//		return err
//	}
//	log.Infoln("copying sql database file to: " + tmpSqlFile)
//	if err := utils.CopyFile(sqlFile, tmpSqlFile); err != nil {
//		return err
//	}
//	return err
//}
