package myLogger

import (
	"log"
	"net/http"
	"time"
)

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
