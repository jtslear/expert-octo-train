package mapper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// DurationAPI accepts the json output from google map api duration
type DurationAPI struct {
	Text string `json:"text"`
}

// LegAPI accepts the json output from google map api legs
type LegAPI struct {
	Duration DurationAPI `json:"duration"`
}

// RouteAPI accepts the json output from google map api routes
type RouteAPI struct {
	Legs [1]LegAPI `json:"legs"`
}

// MapsAPI accepts the json output from google map api
type MapsAPI struct {
	Routes [1]RouteAPI `json:"routes"`
	Status string      `json:"status"`
}

// GetDuration returns entire MapsAPI struct
func GetDuration(origin string, dest string) (string, error) {
	mapsAPIURL := "http://maps.googleapis.com/maps/api/directions/json?"
	values := url.Values{
		"origin":      []string{origin},
		"destination": []string{dest},
	}
	mapsAPIURL += values.Encode()

	response, err := http.Get(mapsAPIURL)
	if err != nil {
		return "", fmt.Errorf("GET Error: %s", err)
	}
	defer response.Body.Close()

	var bar MapsAPI
	decodedJSON := json.NewDecoder(response.Body)
	if err = decodedJSON.Decode(&bar); err != nil {
		return "", fmt.Errorf("JSON Decode Error: %s", err)
	}

	return bar.Routes[0].Legs[0].Duration.Text, nil

}
