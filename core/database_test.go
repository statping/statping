package core

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hunterlong/statup/core/mock"
	"github.com/stretchr/testify/assert"
)

type databaseMock struct{}

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
ALTER TABLE core ADD COLUMN migration_id integer default 0 NOT NULL;
`
	databaseMock := mock.NewMockDatabase(controller)
	migrator := newMigrator(databaseMock, migrationContent)
	migrations, err := migrator.getMigrations()
	assert.NoError(t, err)
	assert.Equal(t, 4, len(migrations))
	assert.Equal(t, int64(1), migrations[0].version)
	assert.Equal(
		t,
		"ALTER TABLE core ADD COLUMN migration_id integer default 0 NOT NULL;",
		migrations[0].statements,
	)
	assert.Equal(t, int64(1530841150), migrations[1].version)
	assert.Equal(
		t,
		"ALTER TABLE core ADD COLUMN use_cdn BOOL NOT NULL DEFAULT '0';",
		strings.TrimSpace(migrations[1].statements),
	)
	assert.Equal(t, int64(1532068515), migrations[2].version)
	assert.Equal(
		t,
		`ALTER TABLE services ALTER COLUMN order_id integer DEFAULT 0;
ALTER TABLE services ADD COLUMN timeout integer DEFAULT 30;`,
		strings.TrimSpace(migrations[2].statements),
	)
	assert.Equal(t, int64(1532068518), migrations[3].version)
	assert.Equal(
		t,
		`CREATE TABLE core (
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
);`,
		strings.TrimSpace(migrations[3].statements),
	)

}
