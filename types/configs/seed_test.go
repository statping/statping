package configs

import (
	"github.com/romanyx/polluter"
	"github.com/statping/statping/database"
	"github.com/statping/statping/utils"
	"os"
	"testing"
)

//func preparePostgresDB(t *testing.T) (database.Database, error) {
//	dbName := fmt.Sprintf("db_%d", time.Now().UnixNano())
//	db, err := database.Openw("sqlite3", dbName)
//	if err != nil {
//		t.Fatalf("open connection: %s", err)
//	}
//
//	return db, db.Error()
//}

func TestSeedDatabase(t *testing.T) {
	t.SkipNow()
	dir := utils.Directory
	f, err := os.Open(dir + "/testdata.yml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	p := polluter.New(polluter.PostgresEngine(database.DB().DB()))

	if err := p.Pollute(f); err != nil {
		t.Fatalf("failed to pollute: %s", err)
	}

}
