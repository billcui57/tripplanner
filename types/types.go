package types

import (
	"errors"
)

var ErrorNotEnoughSites = errors.New("Not enough sites to get route")
var ErrorDirectionApiFatal = errors.New("Something went wrong with Directions API")
var ErrorNoRoutesFound = errors.New("Could not find a route given constraints, please loosen constraints")
var ErrorNoHotelFound = errors.New("Could not find enough hotels in route given constraints, please loosen constraints")

var ErrorHotelApiFatal = errors.New("Something went wrong with Hotel API")
var ErrorHotelApiQuotaExceeded = errors.New("We've ran out of free quota for hotel api!")
