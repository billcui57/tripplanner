package tripplanService

import (
	amadeusService "github/billcui57/tripplanner/Services/HotelService"
	routeService "github/billcui57/tripplanner/Services/RouteService"
	types "github/billcui57/tripplanner/Types"
)

func PlanTrip(sites []types.ISite, maxDrivingHours float64, hotelFindingRadius int) ([]types.DayDriveWithHotel, error) {
	dayDrives, err := routeService.GetDaysDrives(sites, maxDrivingHours)
	if err != nil {
		return nil, err
	}
	dayDrivesWithHotels := make([]types.DayDriveWithHotel, len(dayDrives))
	for i, dayDrive := range dayDrives {
		hotelGeoCodes := amadeusService.FindHotelForDayDrive(dayDrive, hotelFindingRadius)
		if hotelGeoCodes == nil {
			return nil, types.ErrorNoHotelFound
		}
		dayDrivesWithHotels[i] = types.DayDriveWithHotel{DayDrive: dayDrive, HotelGeoCodes: hotelGeoCodes}
	}
	return dayDrivesWithHotels, nil
}
