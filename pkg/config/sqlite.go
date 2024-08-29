package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	logger = GetLogger("sqlite")
)

func initSQLite() (*gorm.DB, error) {
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
		defer file.Close()
	}

	return nil
}

type SQLiteInfo struct {
	DatabaseInfo
	PageSize string `json:"page_size"`
	CacheSize string `json:"cache_size"`
}

func sqliteInfo() SQLiteInfo {
	var version string
	var maxConn string
	var pgSize string
	var chSize string

	db.Raw("SELECT sqlite_version()").Scan(&version)
	db.Raw("PRAGMA max_page_count").Scan(&maxConn)
	db.Raw("PRAGMA page_size").Scan(&pgSize)
	db.Raw("PRAGMA cache_size").Scan(&chSize)

	return SQLiteInfo{
		DatabaseInfo: DatabaseInfo{
			Name: "SQLite",
			Version: version,
			MaxConnections: maxConn,
		},
		PageSize: pgSize,
		CacheSize: chSize,
	}
}