package utils

import (
	"fmt"
	types "github/billcui57/tripplanner/Types"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

func GetEnvVar(varName string) string {
	envVar := os.Getenv(varName)
	if envVar == "" {
		log.Fatal("Env variable not set")
	}
	return envVar
}

func GeoCodeToLatLng(geoCode types.IGeoCode) maps.LatLng {
	return maps.LatLng{Lat: geoCode.Latitude, Lng: geoCode.Longitude}
}

func TextualizeGeoCode(geoCode types.IGeoCode) string {
	return fmt.Sprintf("%v %v", geoCode.Latitude, geoCode.Longitude)

}
