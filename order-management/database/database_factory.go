package database

import "fmt"

func CreateDatabase(driver, dsn string) (DatabaseInterface, error) {
	var db DatabaseInterface
	var err error

	switch driver {
	case "mariadb":
		db = &MariaDBDatabase{}
		err = db.SetupDatabase(dsn)
	default:
		return nil, fmt.Errorf("unknown driver: %s", driver)
	}

	if err != nil {
		return nil, err
	}

	return db, nil
}
