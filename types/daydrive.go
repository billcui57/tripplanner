package types

type IDayDrive struct {
	DurationInSeconds int64    `json:"duration_in_seconds"`
	StartLocation     IGeoCode `json:"start_location"`
	EndLocation       IGeoCode `json:"end_location"`
	DistanceInMeters  int      `json:"distance_in_meters"`
}

// IDayDrive builder pattern code
type IDayDriveBuilder struct {
	iDayDrive *IDayDrive
}

func NewIDayDriveBuilder() *IDayDriveBuilder {
	iDayDrive := &IDayDrive{}
	b := &IDayDriveBuilder{iDayDrive: iDayDrive}
	return b
}

func (b *IDayDriveBuilder) DurationInSeconds(durationInSeconds int64) *IDayDriveBuilder {
	b.iDayDrive.DurationInSeconds = durationInSeconds
	return b
}

func (b *IDayDriveBuilder) StartLocation(startLocation IGeoCode) *IDayDriveBuilder {
	b.iDayDrive.StartLocation = startLocation
	return b
}

func (b *IDayDriveBuilder) EndLocation(endLocation IGeoCode) *IDayDriveBuilder {
	b.iDayDrive.EndLocation = endLocation
	return b
}

func (b *IDayDriveBuilder) DistanceInMeters(distanceInMeters int) *IDayDriveBuilder {
	b.iDayDrive.DistanceInMeters = distanceInMeters
	return b
}

func (b *IDayDriveBuilder) Build() *IDayDrive {
	return b.iDayDrive
}
