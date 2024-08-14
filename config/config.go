package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = initSQLite()
	if err != nil {
		return fmt.Errorf("Error to initializing sqlite: %v", err)
	}

	return nil
}

func GetDBConn() *gorm.DB{
	return db
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)

	return logger
}