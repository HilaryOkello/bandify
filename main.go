package main

import (
    "log"
    "net/http"
	"groupie-tracker/backend"
)

func main() {
    http.HandleFunc("/", backend.HomeHandler) // Home page
    http.HandleFunc("/artists", backend.ArtistsHandler)
    http.HandleFunc("/locations", backend.LocationsHandler)
    http.HandleFunc("/dates", backend.DatesHandler)
    http.HandleFunc("/relations", backend.RelationsHandler)

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
