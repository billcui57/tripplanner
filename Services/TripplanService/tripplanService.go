package tripplanService

import (
	amadeusService "github/billcui57/tripplanner/Services/HotelService"
	routeService "github/billcui57/tripplanner/Services/RouteService"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"
)

func PlanTrip(sites []types.ISite, maxDrivingHours float64, hotelFindingRadius int) ([]types.IGeoCode, []types.IDayDriveWithHotel, error) {
	route, err := routeService.GetRoute(sites)
	if err != nil {
		return nil, nil, err
	}
	dayDrives, err := routeService.GetSegmentRouteInToDaysDrives(route, maxDrivingHours)
	if err != nil {
		return nil, nil, err
	}
	dayDrivesWithHotels := make([]types.IDayDriveWithHotel, len(dayDrives))
	for i, dayDrive := range dayDrives {
		hotels, _ := amadeusService.FindHotelForDayDrive(dayDrive, hotelFindingRadius)
		if err != nil {
			return nil, nil, err
		}

		hotels = make([]types.IHotel, 0)
		if (hotels == nil) || (len(hotels) == 0) {
			return nil, nil, types.ErrorNoHotelFound
		}
		dayDrivesWithHotels[i] = types.IDayDriveWithHotel{DayDrive: dayDrive, Hotels: hotels}
	}

	routeLatLngs, err := route.OverviewPolyline.Decode()
	if err != nil {
		return nil, nil, err
	}

	return utils.LatLngsToGeoCodes(routeLatLngs), dayDrivesWithHotels, nil
}
