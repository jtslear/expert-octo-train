package main

import (
	"encoding/json"
	"fmt"
	"github.com/jtslear/expert-octo-train/httpService"
	"github.com/jtslear/expert-octo-train/mapper"
	"io"
	"net/http"
)

// Locations holds the duration of the locaitons
type Locations struct {
	Location string
	Duration string
}

// Choices contains our choices
type Choices struct {
	Potential string
	Locations []Locations
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func getDurations(potentialPlace string, keyPoints []string) Choices {

	var blobs []Locations

	for _, address := range keyPoints {
		duration, err := mapper.GetDuration(address, potentialPlace)
		if err != nil {
			fmt.Println("Unable to get duration for", address, err)
		}
		blobs = append(blobs, Locations{Location: address, Duration: duration})
	}
	stuff := Choices{Potential: potentialPlace, Locations: blobs}
	return stuff
}

func getPlaces(allTheThings Choices) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		a, err := json.Marshal(allTheThings)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(a)
	}
}

func main() {

	seed := []string{
		"New York City, NY",
		"Washington DC",
	}

	potentialPlace := "Raleigh, NC"

	result := getDurations(potentialPlace, seed)

	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/", getPlaces(result))

	err := httpService.StartServer()
	if err != nil {
		panic(err)
	}

}
