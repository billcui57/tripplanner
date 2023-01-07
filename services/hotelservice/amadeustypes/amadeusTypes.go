package amadeustypes

import types "github/billcui57/tripplanner/types"

type IAccessTokenRequest struct {
	GrantType    string `url:"grant_type"`
	ClientId     string `url:"client_id"`
	ClientSecret string `url:"client_secret"`
}

type IAccessTokenResponse struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientId        string `json:"client_id"`
	TokenType       string `json:"token_type`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

type IGetHotelsByGeocodeRequest struct {
	Latitude    float64 `url:"latitude"`
	Longitude   float64 `url:"longitude"`
	Radius      int     `url:"radius"`
	RadiusUnit  string  `url:"radiusUnit"`
	HotelSource string  `url:"hotelSource"`
}

type IGetHotelsByGeocodeResponse struct {
	Data []IHotel `json:"data"`
	Meta IMeta    `json:"meta"`
}

type IAddress struct {
	CountryCode string `json:"country_code"`
}

type IUnit string

const (
	NIGHT          IUnit = "night"
	PIXELS               = "pixels"
	KILOGRAMS            = "kilograms"
	POUNDS               = "pounds"
	CENTIMETERS          = "centers"
	INCHES               = "inches"
	BITS_PER_PIXEL       = "bits_per_pixel"
	KILOMETERS           = "kilometers"
	MILES                = "miles"
	BYTES                = "bytes"
	KILOBYTES            = "kilobytes"
)

type IDistance struct {
	Unit         IUnit   `json:"unit`
	Value        float64 `json:"value"`
	DisplayValue string  `json:"displayValue"`
	IsUnlimited  string  `json:"isUnlimited"`
}

type IHotel struct {
	SubType          string         `json:"subType"`
	Name             string         `json:"name"`
	TimeZoneName     string         `json:"timeZoneName"`
	IataCode         string         `json:"iataCode"`
	Address          IAddress       `json:"address"`
	GeoCode          types.IGeoCode `json:"geoCode"`
	GooglePlaceId    string         `json: "googlePlaceId"`
	OpenjetAirportId string         `json: "openjetAirportId"`
	UicCode          string         `json:"uicCode"`
	HotelId          string         `json:"hotelId"`
	ChainCode        string         `json:"chainCode"`
	Distance         IDistance      `json:"distance"`
	LastUpdate       string         `json:"last_update"`
}

type ILinks struct {
	self  string `json:"self"`
	first string `json:"first"`
	prev  string `json:"prev"`
	next  string `json:"next"`
	last  string `json:"last"`
}

type IMeta struct {
	Count int64    `json:"count"`
	Sort  []string `json:"sort"`
	Links ILinks   `json:"links`
}
