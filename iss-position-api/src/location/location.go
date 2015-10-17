package location

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ipApiResponse struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type GPSCoords struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func GetMyLocation() (*GPSCoords, error) {
	// Find my IP
	resp, err := http.Get("http://ip-api.com/json")
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var ipResp ipApiResponse
	err = json.Unmarshal(contents, &ipResp)
	if err != nil {
		return nil, err
	}
	return &GPSCoords{
		Latitude:  ipResp.Latitude,
		Longitude: ipResp.Longitude,
	}, nil
}
