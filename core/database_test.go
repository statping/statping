package core

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hunterlong/statup/core/mock"
)

type databaseMock struct{}

// func (dm databaseMock) Driver() interface{}         { return nil }
// func (dm databaseMock) Open(db.ConnectionURL) error { return nil }
// func (dm databaseMock) Ping() error                 { return nil }
// func (dm databaseMock) Close() error                { return nil }
// func (dm databaseMock) Collection(string) db.Collection { return nil }
// func (dm databaseMock) Collections() ([]string, error) { return nil,nil}
// func (dm databaseMock) Name() string { return "mock"}
// func (dm databaseMock) ConnectionURL() db.ConnectionURL { return nil}
// func (dm databaseMock) ClearCache() {}

func TestMigratorGetMigrations(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	migrationContent := `
=========================================== 1532068518
CREATE TABLE core (
	name VARCHAR(50),
	description text,
	config VARCHAR(50),
	api_key VARCHAR(50),
	api_secret VARCHAR(50),
	style text,
	footer text,
	domain text,
	version VARCHAR(50),
	migration_id INT(6) NOT NULL DEFAULT 0,
	use_cdn BOOL NOT NULL DEFAULT '0'
);
=========================================== 1532068515
ALTER TABLE services ALTER COLUMN order_id integer DEFAULT 0;
ALTER TABLE services ADD COLUMN timeout integer DEFAULT 30;
=========================================== 1530841150
ALTER TABLE core ADD COLUMN use_cdn BOOL NOT NULL DEFAULT '0';
=========================================== 1
`
	migrator := newMigrator(mock.NewMockDatabase(controller), migrationContent, 1)
	migrations := migrator.getMigrations()
}
