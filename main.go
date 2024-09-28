package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/server"
)

func main() {
	http.HandleFunc("/", server.MainPage)
	http.HandleFunc("/artists/", server.InfoAboutArtist)
	http.HandleFunc("/search/", server.SearchHandler)
	fmt.Println("Server running on http://localhost:4949/")
	err := http.ListenAndServe(":4949", nil)
	if err != nil {
		log.Fatal(err)
	}
}
