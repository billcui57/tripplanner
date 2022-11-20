package types

type IGeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type OriginDestination struct {
	Origin      IGeoCode
	Destination IGeoCode
}
