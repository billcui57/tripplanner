package tripplancontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlanTripInput struct {
	Types []string `json:"types" binding:"required"`
}

func Plantrip(context *gin.Context) {
	var input PlanTripInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": input})

}
