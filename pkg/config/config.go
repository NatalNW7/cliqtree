package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	env := os.Getenv("CLIQTREE_ENV")
	if env == "" {
		env = "local"
		os.Setenv("CLIQTREE_ENV", env)
	} else {
		err = godotenv.Load(".env."+env)
		if err != nil {
			logger.Errorf("Error to load env: %v", err)
		}
	}

	db, err = initDatabase(env)
	if err != nil {
		return fmt.Errorf("Error to initializing database: %v", err)
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