package tripplancontroller

import (
	"errors"
	tripplanService "github/billcui57/tripplanner/services/tripplanservice"
	types "github/billcui57/tripplanner/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPlanTripRequest struct {
	Sites              []types.ISite `json:"sites" binding:"required"`
	MaxDrivingSeconds  int64         `json:"max_driving_seconds" binding:"required"`
	HotelFindingRadius int           `json:"hotel_finding_radius" binding:"required"`
}

type IPlanTripResponse struct {
	DayDriveWithHotels []types.IDayDriveWithHotel `json:"day_drive_with_hotels"`
	Sites              []types.ISite              `json:"sites"`
	LeanRoute          types.ILeanRoute           `json:"lean_route"`
}

func Plantrip(context *gin.Context) {
	var input IPlanTripRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	leanRoute, dayDriveWithHotels, err := tripplanService.PlanTrip(
		input.Sites, input.MaxDrivingSeconds, input.HotelFindingRadius,
	)

	if err != nil {
		switch {
		case errors.Is(err, types.ErrorDirectionApiFatal):
			context.JSON(http.StatusInternalServerError, err.Error())
		case errors.Is(err, types.ErrorNoHotelFound):
			context.JSON(http.StatusBadRequest, err.Error())
		case errors.Is(err, types.ErrorNoRoutesFound):
			context.JSON(http.StatusBadRequest, err.Error())
		case errors.Is(err, types.ErrorNotEnoughSites):
			context.JSON(http.StatusBadRequest, err.Error())
		case errors.Is(err, types.ErrorHotelApiFatal):
			context.JSON(http.StatusInternalServerError, err.Error())
		case errors.Is(err, types.ErrorHotelApiQuotaExceeded):
			context.JSON(http.StatusTooManyRequests, err.Error())
		}
		return
	}

	response := IPlanTripResponse{DayDriveWithHotels: dayDriveWithHotels, Sites: input.Sites, LeanRoute: *leanRoute}
	context.JSON(http.StatusOK, response)

}
