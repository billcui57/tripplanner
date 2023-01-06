package utils

import (
	"fmt"
	types "github/billcui57/tripplanner/Types"
	"log"
	"math"
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

func SampleNGeoCodes(geoCodes []types.IGeoCode, N int) []types.IGeoCode {
	if N < 2 {
		log.Fatal("Must sample enough points to include start and end")
	}

	chunkSize := len(geoCodes) / N

	var sampledGeoCodes []types.IGeoCode
	for {
		if len(geoCodes) == 0 {
			break
		}

		if len(geoCodes) < chunkSize {
			chunkSize = len(geoCodes)
			sampledGeoCodes = append(sampledGeoCodes, geoCodes[0:chunkSize][len(geoCodes)-1])
		} else {
			sampledGeoCodes = append(sampledGeoCodes, geoCodes[0:chunkSize][0])
		}

		geoCodes = geoCodes[chunkSize:]

	}
	return sampledGeoCodes
}

func GoogleLegToLeg(leg maps.Leg) types.ILeg {
	return types.ILeg{DistanceInMeters: leg.Distance.Meters, DurationInHours: leg.Duration.Hours(), StartLocation: LatLngToGeoCode(leg.StartLocation), EndLocation: LatLngToGeoCode(leg.EndLocation)}
}

func GoogleLegsToLegs(legs []*maps.Leg) []types.ILeg {

	result := make([]types.ILeg, len(legs))
	for i, leg := range legs {
		result[i] = GoogleLegToLeg(*leg)
	}

	return result
}

func GoogleSteptoStep(step maps.Step) types.IStep {
	return types.IStep{DistanceInMeters: step.Distance.Meters, DurationInHours: step.Duration.Hours(), StartLocation: LatLngToGeoCode(step.StartLocation), EndLocation: LatLngToGeoCode(step.EndLocation)}
}

func GoogleStepstoSteps(steps []*maps.Step) []types.IStep {

	result := make([]types.IStep, len(steps))
	for i, step := range steps {
		result[i] = GoogleSteptoStep(*step)
	}

	return result
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
