package config

import (
	"github.com/NatalNW7/link.in/internal/link.in/schemas"
	"gorm.io/gorm"
)


func initDatabase(env string) (*gorm.DB, error){
	var db *gorm.DB
	var err error
	logger := GetLogger("database")

	if env == "local" {
		db, err = initSQLite()
		if err != nil {
			return nil, err
		}
		logger.Debug("using SQLite")
	} else {
		db, err = initPostgreSQL()
		if err != nil {
			return nil, err
		}
		logger.Debug("using PostgreSQL")
	}

	err = db.AutoMigrate(&schemas.Link{})
	if  err != nil {
		logger.Errorf("Error to automigrate schema: %v", err)
		return nil, err
	}

	return db, nil
}