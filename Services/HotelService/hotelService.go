package amadeusApi

import (
	"encoding/json"
	"fmt"
	amadeustypes "github/billcui57/tripplanner/Services/HotelService/AmadeusTypes"
	types "github/billcui57/tripplanner/Types"
	utils "github/billcui57/tripplanner/Utils"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const api = "https://test.api.amadeus.com/v1"

func apiUrlBuilder(baseUrl string, subUrl string, searchParam string) string {
	url := baseUrl
	if subUrl != "" {
		url += "/" + subUrl
	}

	if searchParam != "" {
		url += "?" + searchParam
	}
	return url
}

func getAccessToken() {
	if accessToken != "" {
		return
	}

	grant_type := utils.GetEnvVar("GRANT_TYPE")
	client_id := utils.GetEnvVar("CLIENT_ID")
	client_secret := utils.GetEnvVar("CLIENT_SECRET")

	url := apiUrlBuilder(api, "security/oauth2/token", "")

	requestPayload := amadeustypes.IAccessTokenRequest{
		GrantType:    grant_type,
		ClientId:     client_id,
		ClientSecret: client_secret,
	}

	v, _ := query.Values(requestPayload)
	buf := v.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(buf))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := amadeustypes.IAccessTokenResponse{}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Fatal(err)
	}

	accessToken = res.AccessToken
}

func GetHotelsByGeocode(getHotelsByGeocodeRequest amadeustypes.IGetHotelsByGeocodeRequest) amadeustypes.IGetHotelsByGeocodeResponse {
	getAccessToken()
	v, _ := query.Values(getHotelsByGeocodeRequest)
	searchParams := v.Encode()

	url := apiUrlBuilder(api, "reference-data/locations/hotels/by-geocode", searchParams)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", accessToken))

	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	res := amadeustypes.IGetHotelsByGeocodeResponse{}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func FindHotelForDayDrive(dayDrive types.IDayDrive, hotelFindingRadius int) []types.IHotel {
	endLocation := dayDrive.EndLocation
	hotelsByGeoCodeRequest := amadeustypes.IGetHotelsByGeocodeRequest{
		Latitude:  endLocation.Latitude,
		Longitude: endLocation.Longitude,
		Radius:    hotelFindingRadius, RadiusUnit: "KM", HotelSource: "ALL",
	}
	hotelsByGeoCodeResponse := GetHotelsByGeocode(hotelsByGeoCodeRequest)
	var hotels []types.IHotel
	for _, hotel := range hotelsByGeoCodeResponse.Data {
		hotels = append(hotels, types.IHotel{Location: hotel.GeoCode, Name: hotel.Name})
	}
	return hotels
}

var accessToken string
