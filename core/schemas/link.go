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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
	LinkId string `json:"linkId"`
	RedirectUrl string `json:"redirectUrl"`
}