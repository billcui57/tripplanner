package types

type IHotel struct {
	Name     string   `json:"name"`
	Location IGeoCode `json:"location"`
}

// IHotel builder pattern code
type IHotelBuilder struct {
	iHotel *IHotel
}

func NewIHotelBuilder() *IHotelBuilder {
	iHotel := &IHotel{}
	b := &IHotelBuilder{iHotel: iHotel}
	return b
}

func (b *IHotelBuilder) Name(name string) *IHotelBuilder {
	b.iHotel.Name = name
	return b
}

func (b *IHotelBuilder) Location(location IGeoCode) *IHotelBuilder {
	b.iHotel.Location = location
	return b
}

func (b *IHotelBuilder) Build() *IHotel {
	return b.iHotel
}
