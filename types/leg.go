package types

type ILeg struct {
	DistanceInMeters int      `json:"distance_in_meters"`
	DurationInHours  float64  `json:"duration_in_hours"`
	StartLocation    IGeoCode `json:"start_location"`
	EndLocation      IGeoCode `json:"end_location"`
}

// ILeg builder pattern code
type ILegBuilder struct {
	iLeg *ILeg
}

func NewILegBuilder() *ILegBuilder {
	iLeg := &ILeg{}
	b := &ILegBuilder{iLeg: iLeg}
	return b
}

func (b *ILegBuilder) DistanceInMeters(distanceInMeters int) *ILegBuilder {
	b.iLeg.DistanceInMeters = distanceInMeters
	return b
}

func (b *ILegBuilder) DurationInHours(durationInHours float64) *ILegBuilder {
	b.iLeg.DurationInHours = durationInHours
	return b
}

func (b *ILegBuilder) StartLocation(startLocation IGeoCode) *ILegBuilder {
	b.iLeg.StartLocation = startLocation
	return b
}

func (b *ILegBuilder) EndLocation(endLocation IGeoCode) *ILegBuilder {
	b.iLeg.EndLocation = endLocation
	return b
}

func (b *ILegBuilder) Build() (*ILeg, error) {
	return b.iLeg, nil
}
