package utils

import (
	"fmt"
	types "github/billcui57/tripplanner/types"
	"log"
	"os"

	"googlemaps.github.io/maps"
)

func GetEnvVar(varName string) string {
	env := os.Getenv(varName)
	if env == "" {
		log.Fatalf("Environment variable %v not set", varName)
	}
	return env
}

func GetEnvVarOrDefault(varName string, def string) string {
	env := os.Getenv(varName)
	if env == "" {
		return def
	}
	return env
}

func IsProduction() bool {
	return GetEnvVar("APP_ENV") == "production"
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

func GoogleSteptoStep(step maps.Step) types.IStep {
	return types.IStep{
		DistanceInMeters: step.Distance.Meters, DurationInSeconds: int64(step.Duration.Seconds()),
		StartLocation: LatLngToGeoCode(step.StartLocation), EndLocation: LatLngToGeoCode(step.EndLocation),
	}
}

func GoogleStepstoSteps(steps []*maps.Step) []types.IStep {

	result := make([]types.IStep, len(steps))
	for i, step := range steps {
		result[i] = GoogleSteptoStep(*step)
	}

	return result
}
