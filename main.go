package main

import (
    "log"
    "net/http"
	"groupie-tracker/server"
)

func main() {
    http.HandleFunc("/", server.HomeHandler) // Home page
    http.HandleFunc("/artists", server.ArtistsHandler)
    http.HandleFunc("/locations", server.LocationsHandler)
    http.HandleFunc("/dates", server.DatesHandler)
    http.HandleFunc("/relations", server.RelationsHandler)

    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
