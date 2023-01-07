package types

type IDayDriveWithHotel struct {
	DayDrive IDayDrive `json:"day_drive"`
	Hotels   []IHotel  `json:"hotels"`
}

// IDayDriveWithHotel builder pattern code
type IDayDriveWithHotelBuilder struct {
	iDayDriveWithHotel *IDayDriveWithHotel
}

func NewIDayDriveWithHotelBuilder() *IDayDriveWithHotelBuilder {
	iDayDriveWithHotel := &IDayDriveWithHotel{}
	b := &IDayDriveWithHotelBuilder{iDayDriveWithHotel: iDayDriveWithHotel}
	return b
}

func (b *IDayDriveWithHotelBuilder) DayDrive(dayDrive IDayDrive) *IDayDriveWithHotelBuilder {
	b.iDayDriveWithHotel.DayDrive = dayDrive
	return b
}

func (b *IDayDriveWithHotelBuilder) Hotels(hotels []IHotel) *IDayDriveWithHotelBuilder {
	b.iDayDriveWithHotel.Hotels = hotels
	return b
}

func (b *IDayDriveWithHotelBuilder) Build() *IDayDriveWithHotel {
	return b.iDayDriveWithHotel
}
