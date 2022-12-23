package main

import (
	tripplancontroller "github/billcui57/tripplanner/Controllers/TripplanController"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// origin := types.IGeoCode{Latitude: 44.1067012, Longitude: -79.4410091}
	// middle := types.IGeoCode{Latitude: 43.597729, Longitude: -80.339313}
	// destination := types.IGeoCode{Latitude: 40.3390486, Longitude: -80.0671657}

	engine := gin.Default()

	engine.POST("/plan-trip", tripplancontroller.Plantrip)

	engine.Run()
}
