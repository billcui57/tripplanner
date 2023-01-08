package types

type IStep struct {
	DistanceInMeters  int      `json:"distance_in_meters"`
	DurationInSeconds int64    `json:"duration_in_seconds"`
	StartLocation     IGeoCode `json:"start_location"`
	EndLocation       IGeoCode `json:"end_location"`
}

// IStep builder pattern code
type IStepBuilder struct {
	iStep *IStep
}

func NewIStepBuilder() *IStepBuilder {
	iStep := &IStep{}
	b := &IStepBuilder{iStep: iStep}
	return b
}

func (b *IStepBuilder) DistanceInMeters(distanceInMeters int) *IStepBuilder {
	b.iStep.DistanceInMeters = distanceInMeters
	return b
}

func (b *IStepBuilder) DurationInSeconds(durationInSeconds int64) *IStepBuilder {
	b.iStep.DurationInSeconds = durationInSeconds
	return b
}

func (b *IStepBuilder) StartLocation(startLocation IGeoCode) *IStepBuilder {
	b.iStep.StartLocation = startLocation
	return b
}

func (b *IStepBuilder) EndLocation(endLocation IGeoCode) *IStepBuilder {
	b.iStep.EndLocation = endLocation
	return b
}

func (b *IStepBuilder) Build() *IStep {
	return b.iStep
}
