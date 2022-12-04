package tripplanService

import (
	"fmt"
	amadeusApi "github/billcui57/tripplanner/AmadeusApi"
	routeService "github/billcui57/tripplanner/RouteService"
	types "github/billcui57/tripplanner/Types"
)

func PlanTrip(sites []types.IGeoCode, maxDrivingHours float64) {
	dayDrives := routeService.GetDaysDrives(sites, maxDrivingHours)
	for _, dayDrive := range dayDrives {
		fmt.Printf("Driving for %v hours from %v to %v", dayDrive.DurationInHours, dayDrive.StartLocation, dayDrive.EndLocation)
		amadeusApi.FindHotelForDayDrive(dayDrive)
	}
}
