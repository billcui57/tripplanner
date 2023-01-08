package routeService

import (
	"context"
	"fmt"
	types "github/billcui57/tripplanner/types"
	utils "github/billcui57/tripplanner/utils"

	"googlemaps.github.io/maps"
)

// add api for popularing sites latlng from site name

func GetRoute(sites []types.ISite) (*types.IRoute, error) { //change to use sites.latlng and not sites.name
	if len(sites) < 2 {
		return nil, types.ErrorNotEnoughSites
	}
	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))

	first := utils.TextualizeGeoCode(sites[0].Location, "")
	last := utils.TextualizeGeoCode(sites[len(sites)-1].Location, "")
	var restGeocodes []types.IGeoCode
	for _, site := range sites[1 : len(sites)-1] {
		restGeocodes = append(restGeocodes, site.Location)
	}
	rest := utils.TextualizeGeoCodes(restGeocodes, "via:")

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

	chosenRoute := routes[0]
	steps := getRouteSteps(&chosenRoute)

	path, err := chosenRoute.OverviewPolyline.Decode()
	if err != nil {
		return nil, err
	}
	return &types.IRoute{
		Steps: steps,
		Path:  utils.LatLngsToGeoCodes(path),
	}, nil
}

func getRouteSteps(route *maps.Route) []types.IStep {
	var steps []*maps.Step

	for _, leg := range route.Legs {
		steps = append(steps, leg.Steps...)
	}

	return utils.GoogleStepstoSteps(steps)
}

func SegmentRouteInToDaysDrives(route *types.IRoute, maxDrivingSeconds int64) ([]types.IDayDrive, error) {

	steps := route.Steps
	var daysDrives []types.IDayDrive

	var totalDrivingDuration int64
	totalDrivingDuration = 0
	var totalDrivingDistance int
	totalDrivingDistance = 0
	var startLocation types.IGeoCode
	var endLocation types.IGeoCode
	for i, step := range steps {
		if i == 0 {
			startLocation = step.StartLocation
		}
		if totalDrivingDuration+step.DurationInSeconds > maxDrivingSeconds {
			endLocation = step.StartLocation
			daysDrives = append(
				daysDrives, types.IDayDrive{
					DurationInSeconds: totalDrivingDuration, DistanceInMeters: totalDrivingDistance,
					EndLocation: endLocation, StartLocation: startLocation,
				},
			)
			totalDrivingDuration = 0
			totalDrivingDistance = 0
			startLocation = step.StartLocation
		}

		totalDrivingDuration += step.DurationInSeconds
		totalDrivingDistance += step.DistanceInMeters
	}

	endLocation = steps[len(steps)-1].EndLocation
	daysDrives = append(
		daysDrives, types.IDayDrive{
			DurationInSeconds: totalDrivingDuration, DistanceInMeters: totalDrivingDistance,
			EndLocation: endLocation, StartLocation: startLocation,
		},
	)

	return daysDrives, nil

}
