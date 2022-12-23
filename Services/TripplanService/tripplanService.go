package tripplanService

import (
	amadeusService "github/billcui57/tripplanner/Services/HotelService"
	routeService "github/billcui57/tripplanner/Services/RouteService"
	types "github/billcui57/tripplanner/Types"
)

func PlanTrip(sites []string, maxDrivingHours float64) []types.DaysDriveWithHotels {
	dayDrives := routeService.GetDaysDrives(sites, maxDrivingHours)
	dayDrivesWithHotels := make([]types.DaysDriveWithHotels, len(dayDrives))
	for i, dayDrive := range dayDrives {
		hotelGeoCodes := amadeusService.FindHotelForDayDrive(dayDrive)
		dayDrivesWithHotels[i] = types.DaysDriveWithHotels{DaysDrive: dayDrive, HotelGeoCodes: hotelGeoCodes}
	}
	return dayDrivesWithHotels
}
