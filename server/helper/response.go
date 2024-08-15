package helper

import (
	"github.com/NatalNW7/link.in/core/schemas"
	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, msg string){
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
		"status": code,
	})
}

func SendSuccess(ctx *gin.Context, code int, data schemas.Link) {
	ctx.Header("Content-type", "application/json")
	linkResponse := schemas.LinkReponse{
		ID: data.ID,
		CreateAt: data.CreatedAt,
		// DeleteAt: data.DeletedAt,
		LinkId: data.LinkId,
		RedirectUrl: data.RedirectUrl,
	}
	ctx.JSON(code, gin.H{
		"data": linkResponse,
	})
}