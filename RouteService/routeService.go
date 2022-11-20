package routeService

import (
	"context"
	"fmt"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"
	"log"

	"googlemaps.github.io/maps"
)

func getRoute(origin types.IGeoCode, destination types.IGeoCode) {
	client, err := maps.NewClient(maps.WithAPIKey(utils.GetEnvVar("GOOGLE_API_KEY")))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	r := &maps.DirectionsRequest{
		Origin:      utils.TextualizeGeoCode(origin),
		Destination: utils.TextualizeGeoCode(destination),
	}

	routes, _, err := client.Directions(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	if len(routes) == 0 {
		log.Fatal("No routes rip")
	}

	route := routes[0]

	geoCodes, err := route.OverviewPolyline.Decode()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(geoCodes)
}

func GetRoute(pairs []types.OriginDestination) {
	for _, pairs := range pairs {
		getRoute(pairs.Origin, pairs.Destination)
	}
}
