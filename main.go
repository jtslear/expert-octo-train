package main

import (
	"encoding/json"
	"github.com/jtslear/expert-octo-train/mapper"
	"io"
	"log"
	"net/http"
	"time"
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

// Log cuz we need to log shit
func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		end := time.Now()
		renderTime := end.Sub(start)
		log.Printf("%s %s %s %13v",
			r.RemoteAddr,
			r.Method,
			r.URL,
			renderTime)
	})
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func getPlaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	seed := []string{
		"New York City, NY",
		"Washington DC",
	}

	potentialPlace := "Raleigh, NC"

	var blobs []Locations

	for _, address := range seed {
		duration, err := mapper.GetDuration(address, potentialPlace)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		blobs = append(blobs, Locations{Location: address, Duration: duration})
	}
	stuff := Choices{Potential: potentialPlace, Locations: blobs}
	a, err := json.Marshal(stuff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(a)
}

func main() {
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/", getPlaces)
	http.ListenAndServe(":8000", Log(http.DefaultServeMux))
}
