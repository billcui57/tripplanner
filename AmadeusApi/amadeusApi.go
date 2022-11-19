package amadeusApi

import (
	"encoding/json"
	"fmt"
	amadeustypes "github/billcui57/tripplanner/AmadeusApi/AmadeusTypes"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/joho/godotenv"
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

func getEnvVar(varName string) string {
	envVar := os.Getenv(varName)
	if envVar == "" {
		log.Fatal("Env variable not set")
	}
	return envVar
}

func getAccessToken() string {
	grant_type := getEnvVar("GRANT_TYPE")
	client_id := getEnvVar("CLIENT_ID")
	client_secret := getEnvVar("CLIENT_SECRET")

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

	return res.AccessToken
}

func GetHotelsByGeocode(getHotelsByGeocodeRequest amadeustypes.IGetHotelsByGeocodeRequest) amadeustypes.IGetHotelsByGeocodeResponse {

	v, _ := query.Values(getHotelsByGeocodeRequest)
	searchParams := v.Encode()

	url := apiUrlBuilder(api, "reference-data/locations/hotels/by-geocode", searchParams)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", AccessToken))

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

var AccessToken string

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AccessToken = getAccessToken()
}
