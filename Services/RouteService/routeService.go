package routeService

import (
	"context"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"
	"log"

	"googlemaps.github.io/maps"
)

func GetRouteSteps(sites []types.ISite) []types.Step {
	if len(sites) < 2 {
		log.Fatalln("Not enough sites to get route")
	}
	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))

	first := sites[0].Name
	last := sites[len(sites)-1].Name
	rest := utils.ConvertSitesToWaypoints(sites[1 : len(sites)-1])

	request := &maps.DirectionsRequest{
		Origin:        first,
		Destination:   last,
		Mode:          maps.TravelModeDriving,
		DepartureTime: "now",
		Waypoints:     rest,
	}

	routes, _, err := client.Directions(context.Background(), request)
	if err != nil {
		log.Fatalf("Directions API fatal error: %s", err)
	}

	if len(routes) == 0 {
		log.Fatal("No routes rip")
	}

	route := routes[0]

	steps := []*maps.Step{}

	for _, leg := range route.Legs {
		steps = append(steps, leg.Steps...)
	}

	return utils.GoogleStepstoSteps(steps)
}

func GetDaysDrives(sites []types.ISite, maxDrivingHours float64) []types.DayDrive {

	steps := GetRouteSteps(sites)

	daysDrives := []types.DayDrive{}

	var totalDrivingDuration float64
	totalDrivingDuration = 0
	var totalDrivingDistance int
	totalDrivingDistance = 0
	var startLocation types.IGeoCode
	var endLocation types.IGeoCode
	for i, step := range steps {
		if i == 0 {
			startLocation = step.StartLocation
		}

		if (totalDrivingDuration+step.DurationInHours > maxDrivingHours) || (i == len(steps)-1) {
			endLocation = step.EndLocation
			daysDrives = append(daysDrives, types.DayDrive{DurationInHours: totalDrivingDuration, DistanceInMeters: totalDrivingDistance, EndLocation: endLocation, StartLocation: startLocation})
			totalDrivingDuration = 0
			startLocation = step.EndLocation
		}

		totalDrivingDuration += step.DurationInHours
		totalDrivingDistance += step.DistanceInMeters
	}

	return daysDrives
}
