package handler

import (
	"github.com/NatalNW7/cliqtree/pkg/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetDBConn()
}