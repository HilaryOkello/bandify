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
		Title string
		Data  []Artist
	}{
		Title: "Groupie Trackers - Artists",
		Data:  artists,
	}

	renderTemplate(w, "index.html", data)
}
