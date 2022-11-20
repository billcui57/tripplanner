package main

import (
	routeService "github/billcui57/tripplanner/RouteService"
	types "github/billcui57/tripplanner/Types"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// res := amadeusApi.GetHotelsByGeocode(amadeustypes.IGetHotelsByGeocodeRequest{Latitude: 64.12616, Longitude: -21.9287223, Radius: 5, RadiusUnit: "KM", HotelSource: "ALL"})

	// hotels := res.Data

	// for _, hotel := range hotels {
	// 	fmt.Printf("%+v\n", hotel)
	// }

	origin := types.IGeoCode{Latitude: 44.0533596, Longitude: -79.4668619}
	destination := types.IGeoCode{Latitude: 44.0526304, Longitude: -79.4484726}

	routeService.GetRoute([]types.OriginDestination{{Origin: origin, Destination: destination}})

}
