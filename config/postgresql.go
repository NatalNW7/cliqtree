package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func initPostgreSQL() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("PostgreSQL error: %v", err)
		return nil, err
	}

	return db, nil
}