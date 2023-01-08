package tripplanService

import (
	amadeusService "github/billcui57/tripplanner/services/hotelservice"
	routeService "github/billcui57/tripplanner/services/routeservice"
	types "github/billcui57/tripplanner/types"
	"sync"
)

func concurrentAddDayDriveHotels(dayDrives []types.IDayDrive, hotelFindingRadius int) (
	[]types.IDayDriveWithHotel, error,
) {
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
		return nil, <-errChan //returns first error
	}
	return dayDrivesWithHotels, nil
}

func addDayDriveHotels(dayDrives []types.IDayDrive, hotelFindingRadius int) ([]types.IDayDriveWithHotel, error) {
	return concurrentAddDayDriveHotels(dayDrives, hotelFindingRadius)
}

func PlanTrip(sites []types.ISite, maxDrivingSeconds int64, hotelFindingRadius int) (
	*types.ILeanRoute, []types.IDayDriveWithHotel, error,
) {
	route, err := routeService.GetRoute(sites)
	if err != nil {
		return nil, nil, err
	}

	dayDrives, err := routeService.SegmentRouteInToDaysDrives(route, maxDrivingSeconds)
	if err != nil {
		return nil, nil, err
	}

	dayDrivesWithHotels, err := addDayDriveHotels(dayDrives, hotelFindingRadius)

	if err != nil {
		return nil, nil, err
	}

	return types.NewILeanRouteBuilder().Path(route.Path).Build(), dayDrivesWithHotels, nil
}
