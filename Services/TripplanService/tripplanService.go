package tripplanService

import (
	"fmt"
	amadeusService "github/billcui57/tripplanner/Services/HotelService"
	routeService "github/billcui57/tripplanner/Services/RouteService"
	types "github/billcui57/tripplanner/Types"
)

func PlanTrip(sites []types.IGeoCode, maxDrivingHours float64) {
	dayDrives := routeService.GetDaysDrives(sites, maxDrivingHours)
	for _, dayDrive := range dayDrives {
		fmt.Printf("Driving for %v hours from %v to %v", dayDrive.DurationInHours, dayDrive.StartLocation, dayDrive.EndLocation)
		amadeusService.FindHotelForDayDrive(dayDrive)
	}
}
