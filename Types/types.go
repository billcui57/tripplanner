package types

type IGeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OriginDestination struct {
	Origin      IGeoCode
	Destination IGeoCode
}

type Leg struct {
	Distance        int
	DurationInHours float64
	StartLocation   IGeoCode
	EndLocation     IGeoCode
}

type DaysDrive struct {
	Legs            []Leg
	DurationInHours float64
	StartLocation   IGeoCode
	EndLocation     IGeoCode
	Distance        int
}
