package database

import (
	"fmt"
	"github.com/statping/statping/types"
	"github.com/statping/statping/utils"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	log               = utils.Log
	removeRowsAfter   = types.Day * 90
	maintenceDuration = types.Hour
)

func StartMaintenceRoutine() {
	dur := os.Getenv("REMOVE_AFTER")
	var removeDur time.Duration

	if dur != "" {
		parsedDur, err := time.ParseDuration(dur)
		if err != nil {
			log.Errorf("could not parse duration: %s, using default: %s", dur, removeRowsAfter.String())
			removeDur = removeRowsAfter
		} else {
			removeDur = parsedDur
		}
	} else {
		removeDur = removeRowsAfter
	}

	log.Infof("Service Failure and Hit records will be automatically removed after %s", removeDur.String())
	go databaseMaintence(removeDur)
}

// databaseMaintence will automatically delete old records from 'failures' and 'hits'
// this function is currently set to delete records 7+ days old every 60 minutes
func databaseMaintence(dur time.Duration) {
	//deleteAfter := time.Now().UTC().Add(dur)

	time.Sleep(20 * types.Second)

	for range time.Tick(maintenceDuration) {
		log.Infof("Deleting failures older than %s", dur.String())
		//DeleteAllSince("failures", deleteAfter)

		log.Infof("Deleting hits older than %s", dur.String())
		//DeleteAllSince("hits", deleteAfter)

		maintenceDuration = types.Hour
	}
}

// DeleteAllSince will delete a specific table's records based on a time.
func DeleteAllSince(table string, date time.Time) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE created_at < '%s';", table, database.FormatTime(date))
	q := database.Exec(sql).Debug()
	if q.Error() != nil {
		log.Warnln(q.Error())
	}
}
