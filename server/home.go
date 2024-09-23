package server

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}

	data := struct {
		Title   string
		Artists []Artist
	}{
		Title:   "Groupie Trackers - Home",
		Artists: artists,
	}

	renderTemplate(w, "index.html", data)
}
