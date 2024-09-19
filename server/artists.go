package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

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
		Artists           []Artist
		VisualizationType string
	}{
		Artists:           artists,
		VisualizationType: visualizationType,
	}); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// FetchArtists retrieves artists data from the API and returns it as a slice of Artist structs.
func FetchArtists() ([]Artist, error) {
	var artists []Artist
	err := fetchData(artistsURL, &artists)
	return artists, err
}

// fetchData is a helper function that retrieves data from the specified API URL and unmarshals it into the provided interface.
func fetchData(url string, result interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error: received status code %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body) // Updated to use io.ReadAll
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	return nil
}
