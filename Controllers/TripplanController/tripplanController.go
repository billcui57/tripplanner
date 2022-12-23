package tripplancontroller

import (
	tripplanService "github/billcui57/tripplanner/Services/TripplanService"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PlanTripInput struct {
	Sites []string `json:"sites" binding:"required"`
}

func Plantrip(context *gin.Context) {
	var input PlanTripInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	daysDriveWithHotels := tripplanService.PlanTrip(input.Sites, 2)

	context.JSON(http.StatusOK, gin.H{"data": daysDriveWithHotels})

}
