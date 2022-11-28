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

func LatLngToGeoCode(latLng maps.LatLng) types.IGeoCode {
	return types.IGeoCode{Latitude: latLng.Lat, Longitude: latLng.Lng}
}

func LatLngsToGeoCodes(latLngs []maps.LatLng) []types.IGeoCode {

	result := make([]types.IGeoCode, len(latLngs))
	for i, latLng := range latLngs {
		result[i] = LatLngToGeoCode(latLng)
	}

	return result
}

func TextualizeGeoCode(geoCode types.IGeoCode, prefix string) string {
	return fmt.Sprintf("%s%v %v", prefix, geoCode.Latitude, geoCode.Longitude)
}

func TextualizeGeoCodes(geoCodes []types.IGeoCode, prefix string) []string {
	result := make([]string, len(geoCodes))
	for i, geoCode := range geoCodes {
		result[i] = TextualizeGeoCode(geoCode, prefix)
	}
	return result
}
