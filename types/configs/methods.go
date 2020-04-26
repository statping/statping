package configs

import (
	"fmt"
	"github.com/statping/statping/utils"
	"gopkg.in/yaml.v2"
)

// Save will initially create the config.yml file
func (d *DbConfig) Save(directory string) error {
	c, err := yaml.Marshal(d)
	if err != nil {
		return err
	}
	if err := utils.SaveFile(directory+"/config.yml", c); err != nil {
		return nil
	}
	return nil
}

func (d *DbConfig) ConnectionString() string {
	var conn string
	postgresSSL := utils.Params.GetString("POSTGRES_SSLMODE")

	switch d.DbConn {
	case "memory", ":memory:":
		conn = "sqlite3"
		d.DbConn = ":memory:"
		return d.DbConn
	case "sqlite", "sqlite3":
		conn, err := findDbFile(d)
		if err != nil {
			log.Errorln(err)
		}
		d.SqlFile = conn
		log.Infof("SQL database file at: %s", d.SqlFile)
		d.DbConn = "sqlite3"
		return d.SqlFile
	case "mysql":
		host := fmt.Sprintf("%v:%v", d.DbHost, d.DbPort)
		conn = fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=UTC&time_zone=%%27UTC%%27", d.DbUser, d.DbPass, host, d.DbData)
		return conn
	case "postgres":
		conn = fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v timezone=UTC sslmode=%v", d.DbHost, d.DbPort, d.DbUser, d.DbData, d.DbPass, postgresSSL)
		return conn
	}
	return conn
}
