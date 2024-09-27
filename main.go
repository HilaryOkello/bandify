package main

import (
	"log"
	"net/http"

	"groupie-tracker/server"
)

func main() {
	http.HandleFunc("/static/", server.StaticHandler)
	http.HandleFunc("/", server.HomeHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
