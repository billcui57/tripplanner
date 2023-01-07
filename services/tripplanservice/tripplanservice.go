package tripplanService

import (
	amadeusService "github/billcui57/tripplanner/services/hotelservice"
	routeService "github/billcui57/tripplanner/services/routeservice"
	types "github/billcui57/tripplanner/types"
	utils "github/billcui57/tripplanner/utils"
	"sync"
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
	var wg sync.WaitGroup
	errChan := make(chan error, len(dayDrives))
	for i := range dayDrives {
		wg.Add(1)
		go func(index int) {
			defer func() {
				wg.Done()
			}()
			hotels, err := amadeusService.FindHotelForDayDrive(dayDrives[index], hotelFindingRadius)
			if err != nil {
				errChan <- err
				return
			}
			if (hotels == nil) || (len(hotels) == 0) {
				errChan <- types.ErrorNoHotelFound
				return
			}
			dayDrivesWithHotels[index] = types.IDayDriveWithHotel{DayDrive: dayDrives[index], Hotels: hotels}
		}(i)
	}
	wg.Wait()
	if len(errChan) != 0 {
		return nil, nil, <-errChan //returns first error
	}

	routeLatLngs, err := route.OverviewPolyline.Decode()
	if err != nil {
		return nil, nil, err
	}

	return utils.LatLngsToGeoCodes(routeLatLngs), dayDrivesWithHotels, nil
}
