package backend

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Home handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if err := templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// ArtistsHandler handles requests for the artists page.
// ArtistsHandler handles requests for the artists page.
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := FetchArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get the selected visualization type from the query parameters
	visualizationType := r.URL.Query().Get("type")

	// Render the artists page with the fetched data and selected visualization type
	if err := templates.ExecuteTemplate(w, "artists.html", struct {
		Artists            []Artist
		VisualizationType  string
	}{
		Artists:            artists,
		VisualizationType:  visualizationType,
	}); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}


// LocationsHandler handles requests for the locations page.
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := FetchLocations()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Render the locations page with the fetched data
	if err := templates.ExecuteTemplate(w, "locations.html", locations); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// DatesHandler handles requests for the dates page.
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := FetchDates()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Render the dates page with the fetched data
	if err := templates.ExecuteTemplate(w, "dates.html", dates); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// RelationsHandler handles requests for the relations page.
func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := FetchRelations()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Render the relations page with the fetched data
	if err := templates.ExecuteTemplate(w, "relations.html", relations); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// ErrorHandler handles various HTTP error codes.
// ErrorHandler handles various HTTP error codes and renders an error page.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	var message string

	switch status {
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
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}

	// Render the error page using the errors.html template
	if err := templates.ExecuteTemplate(w, "errors.html", data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
