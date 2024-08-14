package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	LinkId string
	RedirectUrl string
}

type LinkReponse struct {
	ID uint  `json:"id"`
	CreateAt time.Time `json:"createAt"`
	DeleteAt time.Time `json:"deleteAt,omitempty"`
	LinkId string `json:"linkId"`
	RedirectUrl string `json:"redirectUrl"`
}
