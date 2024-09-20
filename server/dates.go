package server

import "net/http"

// DatesHandler handles requests for the dates page.
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := FetchDates()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Render the dates page with the fetched data
	if err := tmpl.ExecuteTemplate(w, "dates.html", dates); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// FetchDates retrieves dates data from the API and returns it as a slice of Date structs.
func FetchDates() ([]Date, error) {
	var dates []Date
	err := fetchData(datesURL, &dates)
	return dates, err
}
