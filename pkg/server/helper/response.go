package helper

import (
	"time"

	"github.com/NatalNW7/cliqtree/internal/cliqtree/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LinkReponse struct {
	ID uint  `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
	LinkId string `json:"linkId"`
	RedirectUrl string `json:"redirectUrl"`
}

func SendError(ctx *gin.Context, code int, msg string, cause string){
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"cause": cause,
	})
}

func SendSuccess(ctx *gin.Context, code int, data *schemas.Link) {
	ctx.Header("Content-type", "application/json")
	linkResponse := LinkReponse{
		ID: data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		LinkId: data.LinkId,
		RedirectUrl: data.RedirectUrl,
	}
	ctx.JSON(code, gin.H{
		"data": linkResponse,
	})
}