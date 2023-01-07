package types

import (
	"errors"
)

type IGeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type ILeg struct {
	DistanceInMeters int      `json:"distance_in_meters"`
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
}

type IStep struct {
	DistanceInMeters int      `json:"distance_in_meters"`
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
}

type IDayDrive struct {
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
	DistanceInMeters int      `json:"distance_in_meters"`
}

type IDayDriveWithHotel struct {
	DayDrive IDayDrive `json:"day_drive"`
	Hotels   []IHotel  `json:"hotels"`
}

type ISite struct {
	Location IGeoCode `json:"location"`
}

type IHotel struct {
	Name     string   `json:"name"`
	Location IGeoCode `json:"location"`
}

var ErrorNotEnoughSites = errors.New("Not enough sites to get route")
var ErrorDirectionApiFatal = errors.New("Something went wrong with Directions API")
var ErrorNoRoutesFound = errors.New("Could not find a route given constraints, please loosen constraints")
var ErrorNoHotelFound = errors.New("Could not find enough hotels in route given constraints, please loosen constraints")

var ErrorHotelApiFatal = errors.New("Something went wrong with Hotel API")
var ErrorHotelApiQuotaExceeded = errors.New("We've ran out of free quota for hotel api!")
