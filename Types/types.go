package types

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

type DaysDrive struct {
	Legs             []Leg    `json:"legs"`
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
	DistanceInMeters int      `json:"distance_in_meters"`
}

type DaysDriveWithHotels struct {
	DaysDrive     DaysDrive  `json:"days_drive"`
	HotelGeoCodes []IGeoCode `json:"hotel_geocodes"`
}
