package tripplancontroller

import (
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
	ResultStatus       types.IResultStatus       `json:"result_status"`
	Sites              []types.ISite             `json:"sites"`
}

func Plantrip(context *gin.Context) {
	var input IPlanTripInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dayDriveWithHotels, resultStatus := tripplanService.PlanTrip(input.Sites, input.MaxDrivingHours, input.HotelFindingRadius)

	response := IPlanTripResponse{DayDriveWithHotels: dayDriveWithHotels, ResultStatus: resultStatus, Sites: input.Sites}

	context.JSON(http.StatusOK, response)

}
