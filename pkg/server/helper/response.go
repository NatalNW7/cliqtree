package helper

import (
	"time"

	"github.com/NatalNW7/cliqtree/internal/cliqtree/schemas"
	"github.com/gin-gonic/gin"
)

type LinkResponse struct {
	ID uint  `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt any `json:"deletedAt,omitempty"`
	LinkId string `json:"linkId"`
	RedirectUrl string `json:"redirectUrl"`
}

type SuccessResponse struct {
	Result LinkResponse `json:"result"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Cause string `json:"cause"`
}

type Dependencies struct{
	Database any `json:"database"`
}
type StatusResponse struct {
	UpdatedAt time.Time `json:"updated_at"`
	Env string	`json:"environment"`
	Dependencies `json:"dependencies"`
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
	linkResponse := LinkResponse{
		ID: data.ID,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
		LinkId: data.LinkId,
		RedirectUrl: data.RedirectUrl,
	}
	ctx.JSON(code, SuccessResponse{
		Result: linkResponse,
	})
}