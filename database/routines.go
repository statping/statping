package database

import (
	"fmt"
	"github.com/statping/statping/utils"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

var (
	log = utils.Log.WithField("type", "database")
)

// Maintenance will automatically delete old records from 'failures' and 'hits'
// this function is currently set to delete records 7+ days old every 60 minutes
// env: REMOVE_AFTER - golang duration parsed time for deleting records older than REMOVE_AFTER duration from now
// env: CLEANUP_INTERVAL - golang duration parsed time for checking old records routine
func Maintenance() {
	dur := utils.Params.GetDuration("REMOVE_AFTER")
	interval := utils.Params.GetDuration("CLEANUP_INTERVAL")

	log.Infof("Database Cleanup runs every %s and will remove records older than %s", interval.String(), dur.String())
	ticker := interval

	for {
		select {
		case <-time.After(ticker):
			deleteAfter := utils.Now().Add(-dur)

			log.Infof("Deleting failures older than %s", deleteAfter.String())
			deleteAllSince("failures", deleteAfter)

			log.Infof("Deleting hits older than %s", deleteAfter.String())
			deleteAllSince("hits", deleteAfter)

			ticker = interval
		}
	}

}

// deleteAllSince will delete a specific table's records based on a time.
func deleteAllSince(table string, date time.Time) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE created_at < '%s'", table, database.FormatTime(date))
	log.Info(sql)
	if err := database.Exec(sql).Error(); err != nil {
		log.WithField("query", sql).Errorln(err)
	}
}
