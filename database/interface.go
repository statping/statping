package database

type DbObject interface {
	Create() error
	Update() error
	Delete() error
}

type Sampler interface {
	Sample() DbObject
}

func MigrateTable(table interface{}) error {
	tx := database.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	tx = tx.AutoMigrate(table)

	if err := tx.Commit().Error(); err != nil {
		return err
	}
	return nil
}
