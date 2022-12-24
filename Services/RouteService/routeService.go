package routeService

import (
	"context"
	"fmt"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"

	"googlemaps.github.io/maps"
)

// add api for popularing sites latlng from site name

func GetRouteSteps(sites []types.ISite) ([]types.Step, error) { //change to use sites.latlng and not sites.name
	if len(sites) < 2 {
		return nil, types.ErrorNotEnoughSites
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
		fmt.Printf("Directions API fatal error: %s\n", err)
		return nil, types.ErrorDirectionApiFatal
	}

	if len(routes) == 0 {
		return nil, types.ErrorNoRoutesFound
	}

	route := routes[0]

	steps := []*maps.Step{}

	for _, leg := range route.Legs {
		steps = append(steps, leg.Steps...)
	}

	return utils.GoogleStepstoSteps(steps), nil
}

func GetDaysDrives(sites []types.ISite, maxDrivingHours float64) ([]types.DayDrive, error) {

	steps, err := GetRouteSteps(sites)
	if err != nil {
		return nil, err
	}

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

	return daysDrives, nil
}
