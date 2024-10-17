package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, ok := templates[tmpl]
	if !ok {
		log.Println(tmpl, "not found")
		ErrorPage(w, http.StatusNotFound)
		return
	}
	err := t.ExecuteTemplate(w, "layout.html", data)
	if err != nil {
		ErrorPage(w, http.StatusInternalServerError)
		return
	}
}

func checkMethodAndPath(w http.ResponseWriter, r *http.Request, method, path string) bool {
	if r.Method != method {
		ErrorPage(w, http.StatusMethodNotAllowed)
		return false
	}
	if r.URL.Path != path {
		ErrorPage(w, http.StatusNotFound)
		return false
	}
	return true
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	if !checkMethodAndPath(w, r, http.MethodGet, "/") {
		return
	}
	data := TemplateData{
		Title: "Groupie Trackers - Artists",
		Data:  artists,
	}
	renderTemplate(w, "index.html", data)
}

func InfoAboutArtist(w http.ResponseWriter, r *http.Request) {
	if !checkMethodAndPath(w, r, http.MethodGet, "/artists/") {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id <= 0 || id > len(artists) || err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusBadRequest)
		return
	}
	id--

	locations, err := FetchLocations(artists[id].Locations)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	dates, err := FetchDates(artists[id].ConcertDates)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	rel, err := FetchRelation(artists[id].Relations)
	if err != nil {
		log.Println(err)
		ErrorPage(w, http.StatusInternalServerError)
		return
	}

	data := TemplateData{
		Title:     "Artist Details",
		Artist:    artists[id],
		Locations: locations,
		Dates:     dates,
		Concerts:  rel,
	}

	renderTemplate(w, "details.html", data)
}

// SearchPage handles the artist search functionality.
func SearchPage(w http.ResponseWriter, r *http.Request) {
	if !checkMethodAndPath(w, r, http.MethodGet, "/search/") {
		return
	}

	query := r.URL.Query().Get("q")
	if query == "" {
		ErrorPage(w, http.StatusBadRequest)
		return
	}

	var results []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}

	data := TemplateData{
		Title:   "Search Results",
		Query:   query,
		Results: results,
	}

	if len(results) == 0 {
		data.Message = "No artists found matching your query."
	}

	renderTemplate(w, "search.html", data)
}

func ErrorPage(w http.ResponseWriter, code int) {
	var message string
	switch code {
	case http.StatusNotFound:
		message = "Not Found"
	case http.StatusBadRequest:
		message = "Bad Request"
	case http.StatusMethodNotAllowed:
		message = "Method Not Allowed"
	case http.StatusForbidden:
		message = "Forbidden"
	default:
		message = "Internal Server Error"
	}
	data := TemplateData{
		Title:   "Error",
		Status:  code,
		Message: message,
	}

	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("%d - %s", code, message), code)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("%d - %s", code, message), code)
	}
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	// Remove the /static/ prefix from the URL path
	filePath := path.Join("static", strings.TrimPrefix(r.URL.Path, "/static/"))

	// Check if the file exists and is not a directory
	info, err := os.Stat(filePath)
	if err != nil || info.IsDir() {
		ErrorPage(w, http.StatusForbidden)
		return
	}

	// Check the file extension
	ext := filepath.Ext(filePath)
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	default:
		ErrorPage(w, http.StatusForbidden)
		return
	}

	// Serve the file
	http.ServeFile(w, r, filePath)
}
