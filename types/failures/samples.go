package failures

import (
	"fmt"
	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"
	"time"
)

var (
	log = utils.Log.WithField("type", "failure")
)

func Example() Failure {
	return Failure{
		Id:        48533,
		Issue:     "Response did not response a 200 status code",
		Method:    "",
		MethodId:  0,
		ErrorCode: 404,
		Service:   1,
		Checkin:   0,
		PingTime:  48309,
		CreatedAt: utils.Now(),
	}
}

func Samples() error {
	log.Infoln("Inserting Sample Service Failures...")
	createdAt := utils.Now().Add(-3 * types.Day)

	for i := int64(1); i <= 4; i++ {
		f1 := &Failure{
			Service:   i,
			Issue:     "Server failure",
			CreatedAt: utils.Now().Add(-time.Duration(3*i) * 86400),
		}
		if err := f1.Create(); err != nil {
			return err
		}

		f2 := &Failure{
			Service:   i,
			Issue:     "Server failure",
			CreatedAt: utils.Now().Add(-time.Duration(5*i) * 12400),
		}
		if err := f2.Create(); err != nil {
			return err
		}

		log.Infoln(fmt.Sprintf("Adding %v Failure records to service", 400))

		var records []interface{}
		for fi := 0.; fi <= float64(400); fi++ {
			failure := &Failure{
				Service:   i,
				Issue:     "testing right here",
				CreatedAt: createdAt.UTC(),
			}
			records = append(records, failure)
			createdAt = createdAt.Add(35 * time.Minute)
		}
		if err := gormbulk.BulkInsert(db.GormDB(), records, db.ChunkSize()); err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
