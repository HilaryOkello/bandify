package server

import (
	"net/http"
)

// ArtistsHandler handles requests for the artists page.
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Title string
		Data  []Artist
	}{
		Title: "Groupie Trackers - Artists",
		Data:  artists,
	}

	renderTemplate(w, "artists.html", data)
}
