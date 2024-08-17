package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	LinkId string `gorm:"unique_index"`
	RedirectUrl string
}

type LinkReponse struct {
	ID uint  `json:"id"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"updateAt"`
	DeleteAt time.Time `json:"deleteAt,omitempty"`
	LinkId string `json:"linkId"`
	RedirectUrl string `json:"redirectUrl"`
}