package tripplanService

import (
	amadeusService "github/billcui57/tripplanner/Services/HotelService"
	routeService "github/billcui57/tripplanner/Services/RouteService"
	types "github/billcui57/tripplanner/Types"
)

func PlanTrip(sites []types.ISite, maxDrivingHours float64, hotelFindingRadius int) ([]types.IDayDriveWithHotel, error) {
	dayDrives, err := routeService.GetDaysDrives(sites, maxDrivingHours)
	if err != nil {
		return nil, err
	}
	dayDrivesWithHotels := make([]types.IDayDriveWithHotel, len(dayDrives))
	for i, dayDrive := range dayDrives {
		hotels := amadeusService.FindHotelForDayDrive(dayDrive, hotelFindingRadius)
		if (hotels == nil) || (len(hotels) == 0) {
			return nil, types.ErrorNoHotelFound
		}
		dayDrivesWithHotels[i] = types.IDayDriveWithHotel{DayDrive: dayDrive, Hotels: hotels}
	}
	return dayDrivesWithHotels, nil
}
