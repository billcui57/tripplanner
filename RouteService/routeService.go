package routeService

import (
	"context"
	"fmt"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"
	"log"

	"googlemaps.github.io/maps"
)

//
//
// Let A be site, B be site after A. For each latLng between A and B (inclusive), use as both destination and origin for DistanceMatrix Api.
// Store result as a double layer hashmap, with origin as first key, and destination as second key, and travel time as value.
// Then, we walk the route, accumulating travel time by hashmap look up. Maybe call hotel API ever half hour?
//
//
//
//
// ISSUE: Users of the standard API:
// 2,500 free elements per day
// 100 elements per query
// 100 elements per 10 seconds
// Each query sent to the Distance Matrix API generates elements, where the number of origins times the number of destinations equals the number of elements.
// That is not enough for our use case
//
//
//
//
//
//
//
//
//
//
//
//

// func getRoute(origin types.IGeoCode, destination types.IGeoCode) []types.IGeoCode {
// 	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))
// 	if err != nil {
// 		log.Fatalf("fatal error: %s", err)
// 	}

// 	r := &maps.DirectionsRequest{
// 		Origin:        utils.TextualizeGeoCode(origin, ""),
// 		Destination:   utils.TextualizeGeoCode(destination, ""),
// 		Mode:          maps.TravelModeDriving,
// 		DepartureTime: "now",
// 	}

// 	routes, _, err := client.Directions(context.Background(), r)
// 	if err != nil {
// 		log.Fatalf("fatal error: %s", err)
// 	}

// 	if len(routes) == 0 {
// 		log.Fatal("No routes rip")
// 	}

// 	route := routes[0]

// 	routeLatLngs, err := route.OverviewPolyline.Decode()

// 	route.Legs[0].

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return utils.LatLngsToGeoCodes(routeLatLngs)
// }

// func GetRoute(sites []types.IGeoCode) {
// 	if len(sites) < 2 {
// 		log.Fatalln("Not enough sites to get route")
// 	}

// 	for i := 0; i < len(sites)-1; i++ {
// 		getRoute(sites[i], sites[i+1])
// 	}
// }

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

func Test(sites []types.IGeoCode) {
	geoCodes := GetRouteGeoCodes(sites)
	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))

	origins := []types.IGeoCode{}
	destinations := []types.IGeoCode{}

	if len(geoCodes) < 2 {
		log.Fatalln("Not enough geoCodes available")
	}

	for i := 0; i < len(geoCodes)-1; i++ {
		origin := geoCodes[i]
		destination := geoCodes[i+1]
		origins = append(origins, origin)
		destinations = append(destinations, destination)
	}

	request := &maps.DistanceMatrixRequest{
		Origins:      utils.TextualizeGeoCodes(origins, ""),
		Destinations: utils.TextualizeGeoCodes(destinations, ""),
		Mode:         maps.TravelModeDriving,
	}

	distanceMatrix, err := client.DistanceMatrix(context.Background(), request)
	if err != nil {
		log.Fatalf("Distance Matrix API fatal error: %s", err)
	}
	fmt.Println(distanceMatrix)

}
