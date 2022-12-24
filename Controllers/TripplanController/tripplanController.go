package tripplancontroller

import (
	"errors"
	tripplanService "github/billcui57/tripplanner/Services/TripplanService"
	types "github/billcui57/tripplanner/Types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPlanTripInput struct {
	Sites              []types.ISite `json:"sites" binding:"required"`
	MaxDrivingHours    float64       `json:"max_driving_hours" binding:"required"`
	HotelFindingRadius int           `json:"hotel_finding_radius" binding:"required"`
}

type IPlanTripResponse struct {
	DayDriveWithHotels []types.DayDriveWithHotel `json:"day_drive_with_hotels"`
	Sites              []types.ISite             `json:"sites"`
}

func Plantrip(context *gin.Context) {
	var input IPlanTripInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dayDriveWithHotels, err := tripplanService.PlanTrip(input.Sites, input.MaxDrivingHours, input.HotelFindingRadius)

	if err != nil {
		switch {
		case errors.Is(err, types.ErrorDirectionApiFatal):
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		case errors.Is(err, types.ErrorNoHotelFound):
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, types.ErrorNoRoutesFound):
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case errors.Is(err, types.ErrorNotEnoughSites):
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	response := IPlanTripResponse{DayDriveWithHotels: dayDriveWithHotels, Sites: input.Sites}
	context.JSON(http.StatusOK, response)

}
