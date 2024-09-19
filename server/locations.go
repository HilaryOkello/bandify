package server

import "net/http"

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

// FetchLocations retrieves locations data from the API and returns it as a slice of Location structs.
func FetchLocations() ([]Location, error) {
	var locations []Location
	err := fetchData(locationsURL, &locations)
	return locations, err
}
