package routeService

import (
	"context"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"
	"log"

	"googlemaps.github.io/maps"
)

func GetRouteGeoCodes(sites []types.IGeoCode) []types.IGeoCode {
	if len(sites) < 2 {
		log.Fatalln("Not enough sites to get route")
	}
	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))

	first := utils.TextualizeGeoCode(sites[0], "")
	last := utils.TextualizeGeoCode(sites[len(sites)-1], "")
	rest := utils.TextualizeGeoCodes(sites[1:len(sites)-1], "via:")

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
	latLngs, err := route.OverviewPolyline.Decode()
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return utils.LatLngsToGeoCodes(latLngs)
}

func SplitRouteIntoLegs(sites []types.IGeoCode) []types.Leg {
	geoCodes := GetRouteGeoCodes(sites)
	shortenedGeoCodes := utils.SampleNGeoCodes(geoCodes, 10)

	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))

	first := utils.TextualizeGeoCode(shortenedGeoCodes[0], "")
	last := utils.TextualizeGeoCode(shortenedGeoCodes[len(shortenedGeoCodes)-1], "")
	rest := utils.TextualizeGeoCodes(shortenedGeoCodes[1:len(shortenedGeoCodes)-1], "")

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

	return utils.GoogleLegsToLegs(routes[0].Legs)
}

func GetDaysDrives(sites []types.IGeoCode, maxDrivingHours float64) []types.DaysDrive {

	legs := SplitRouteIntoLegs(sites)

	daysDrives := []types.DaysDrive{}

	dayDriveLegs := []types.Leg{}
	var totalDrivingDuration float64
	totalDrivingDuration = 0
	var totalDrivingDistance int
	totalDrivingDistance = 0
	var startLocation types.IGeoCode
	var endLocation types.IGeoCode
	for i, leg := range legs {
		if i == 0 {
			startLocation = leg.StartLocation
		}

		if totalDrivingDuration+leg.DurationInHours > maxDrivingHours {
			endLocation = leg.EndLocation
			daysDrives = append(daysDrives, types.DaysDrive{Legs: dayDriveLegs, DurationInHours: totalDrivingDuration, Distance: totalDrivingDistance, EndLocation: endLocation, StartLocation: startLocation})
			totalDrivingDuration = 0
			dayDriveLegs = []types.Leg{}
			startLocation = leg.EndLocation
		}

		dayDriveLegs = append(dayDriveLegs, leg)
		totalDrivingDuration += leg.DurationInHours
		totalDrivingDistance += leg.Distance
	}

	return daysDrives
}
