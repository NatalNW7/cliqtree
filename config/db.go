package config

import (
	"os"

	"github.com/NatalNW7/link.in/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	logger *Logger
)

func initSQLite() (*gorm.DB, error) {
	logger = GetLogger("sqlite")
	dbPath := "./db/sqlite.db"

	err := createSQLiteFile(dbPath)
	if err != nil {
		logger.Errorf("Error to create sqlite: %v", err)
		return nil, err
	}

	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SQLite error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Link{})
	if  err != nil {
		logger.Errorf("Error to automigrate schema: %v", err)
		return nil, err
	}

	return db, nil
}

func createSQLiteFile(dbPath string) error {
	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("creating db foler")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}

		logger.Info("creating sqlite file")
		file, err := os.Create(dbPath)
		if err != nil {
			return err
		}
		file.Close()
	}

	// os.Stat()

	return nil
}