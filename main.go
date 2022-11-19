package main

import (
	"fmt"
	amadeusApi "github/billcui57/tripplanner/AmadeusApi"
	amadeustypes "github/billcui57/tripplanner/AmadeusApi/AmadeusTypes"
)

func main() {
	res := amadeusApi.GetHotelsByGeocode(amadeustypes.IGetHotelsByGeocodeRequest{Latitude: 64.12616, Longitude: -21.9287223, Radius: 5, RadiusUnit: "KM", HotelSource: "ALL"})

	hotels := res.Data

	for _, hotel := range hotels {
		fmt.Printf("%+v\n", hotel)
	}
}
