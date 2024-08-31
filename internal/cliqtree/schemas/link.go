package schemas

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	LinkId string `gorm:"unique_index"`
	RedirectUrl string
}
