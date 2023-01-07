package types

type IGeoCode struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// IGeoCode builder pattern code
type IGeoCodeBuilder struct {
	iGeoCode *IGeoCode
}

func NewIGeoCodeBuilder() *IGeoCodeBuilder {
	iGeoCode := &IGeoCode{}
	b := &IGeoCodeBuilder{iGeoCode: iGeoCode}
	return b
}

func (b *IGeoCodeBuilder) Latitude(latitude float64) *IGeoCodeBuilder {
	b.iGeoCode.Latitude = latitude
	return b
}

func (b *IGeoCodeBuilder) Longitude(longitude float64) *IGeoCodeBuilder {
	b.iGeoCode.Longitude = longitude
	return b
}

func (b *IGeoCodeBuilder) Build() *IGeoCode {
	return b.iGeoCode
}
