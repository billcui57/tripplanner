package types

import "errors"

type IGeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Leg struct {
	DistanceInMeters int      `json:"distance_in_meters"`
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
}

type Step struct {
	DistanceInMeters int      `json:"distance_in_meters"`
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
}

type DayDrive struct {
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
	DistanceInMeters int      `json:"distance_in_meters"`
}

type DayDriveWithHotel struct {
	DayDrive      DayDrive   `json:"day_drive"`
	HotelGeoCodes []IGeoCode `json:"hotel_geocodes"`
}

type ISite struct {
	Name string `json:"name" binding:"required"`
}

var ErrorNotEnoughSites = errors.New("Not enough sites to get route")
var ErrorDirectionApiFatal = errors.New("Something went wrong with Directions API")
var ErrorNoRoutesFound = errors.New("Could not find a route")
var ErrorNoHotelFound = errors.New("Could not find a hotel in route")
