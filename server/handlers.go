package server

import (
	"fmt"
	"net/http"
	"strconv"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, ok := templates[tmpl]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	err := t.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPage(w, http.StatusMethodNotAllowed)
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

func InfoAboutArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artists/" {
		ErrorPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id <= 0 || id > len(artists) || err != nil {
		fmt.Println(err)
		ErrorPage(w, http.StatusNotFound)
		return
	}
	id++
	locations, err := FetchLocations(artists[id].Locations)
	if err != nil {
		fmt.Println(err)
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	dates, err := FetchDates(artists[id].ConcertDates)
	if err != nil {
		fmt.Println(err)
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
	rel, err := FetchRelation(artists[id].Relations)
	if err != nil {
		fmt.Println(err)
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	data := struct {
		Title     string
		Artist    Artist
		Locations Loc
		Dates     Date
		Relations Relation
	}{
		Title:     "Artist Details",
		Artist:    artists[id],
		Locations: locations,
		Dates:     dates,
		Relations: rel,
	}

	renderTemplate(w, "details.html", data)
}

// func SearchHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/search/" {
// 		err := "404 Page not found"
// 		ErrorPage(w, err, http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != http.MethodGet {
// 		err := "405 Method is not allowed"
// 		ErrorPage(w, err, http.StatusMethodNotAllowed)
// 		return
// 	}

// 	artists, err := FetchArtists()
// 	if err != nil {
// 		err := "500 Internal Server Error"
// 		ErrorPage(w, err, http.StatusInternalServerError)
// 		return
// 	}
// 	location, err := FetchLocations()
// 	if err != nil {
// 		err := "500 Internal Server Error"
// 		ErrorPage(w, err, http.StatusInternalServerError)
// 		return
// 	}
// 	allInfo := Everything{artists, location}

// 	searchTerm := r.FormValue("Search")
// 	if strings.Contains(searchTerm, " - ") {
// 		searchTermWithoutGroups := strings.Split(searchTerm, " - ")
// 		searchTerm = searchTermWithoutGroups[0]
// 	}

// 	var searchedArtists Everything
// 	if searchTerm != "" {
// 		searchedArtists, err = Search(allInfo, searchTerm)
// 		if err != nil {
// 			err := "500 Internal Server Error"
// 			ErrorPage(w, err, http.StatusInternalServerError)
// 			return
// 		}

// 	}

// 	err = tmp.ExecuteTemplate(w, "search.html", searchedArtists)
// 	if err != nil {
// 		err := "500 Internal Server Error"
// 		ErrorPage(w, err, http.StatusInternalServerError)
// 		return
// 	}
// }

func ErrorPage(w http.ResponseWriter, code int) {
	var message string
	switch code {
	case http.StatusNotFound:
		message = "404 Not Found"
	case http.StatusBadRequest:
		message = "400 Bad Request"
	case http.StatusMethodNotAllowed:
		message = "405 Method Not Allowed"
	default:
		message = "500 Internal Server Error"
	}
	// Prepare the data to be passed to the template
	data := struct {
		Title   string
		Status  int
		Message string
	}{
		Title:   "Error",
		Status:  code,
		Message: message,
	}
	w.WriteHeader(code)
	// Render the error page using the errors.html template
	renderTemplate(w, "error.html", data)
}
