package config

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	env := os.Getenv("CLIQTREE_ENV")
	switch env {
	case "development":
		err = godotenv.Load(".env."+env)
		if err != nil {
			logger.Errorf("Error to load env: %v", err)
		}
	case "production":
		gin.SetMode(gin.ReleaseMode)
		logger.Info("Runnig on production enviroment")
	default:
		env = "local"
		os.Setenv("CLIQTREE_ENV", "local")
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