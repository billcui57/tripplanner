package main

import (
	tripplanService "github/billcui57/tripplanner/TripplanService"
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

	origin := types.IGeoCode{Latitude: 44.1067012, Longitude: -79.4410091}
	middle := types.IGeoCode{Latitude: 43.597729, Longitude: -80.339313}
	destination := types.IGeoCode{Latitude: 40.3390486, Longitude: -80.0671657}

	tripplanService.PlanTrip([]types.IGeoCode{origin, middle, destination}, 2)

}
