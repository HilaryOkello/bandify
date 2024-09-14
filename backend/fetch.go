package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

type Artist struct {
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Members []string `json:"members"`
}

type Location struct {
	Location string `json:"location"`
}

type Date struct {
	Date string `json:"date"`
}

type Relation struct {
	ArtistID   int `json:"artistId"`
	LocationID int `json:"locationId"`
	DateID     int `json:"dateId"`
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

// FetchArtists retrieves artists data from the API and returns it as a slice of Artist structs.
func FetchArtists() ([]Artist, error) {
	var artists []Artist
	err := fetchData(artistsURL, &artists)
	return artists, err
}

// FetchLocations retrieves locations data from the API and returns it as a slice of Location structs.
func FetchLocations() ([]Location, error) {
	var locations []Location
	err := fetchData(locationsURL, &locations)
	return locations, err
}

// FetchDates retrieves dates data from the API and returns it as a slice of Date structs.
func FetchDates() ([]Date, error) {
	var dates []Date
	err := fetchData(datesURL, &dates)
	return dates, err
}

// FetchRelations retrieves relations data from the API and returns it as a slice of Relation structs.
func FetchRelations() ([]Relation, error) {
	var relations []Relation
	err := fetchData(relationsURL, &relations)
	return relations, err
}


// func smth() {
// 	// Example usage
// 	artists, err := FetchArtists()
// 	if err != nil {
// 		log.Fatalf("Failed to fetch artists: %v", err)
// 	}
// 	fmt.Printf("Artists: %+v\n", artists)

// 	locations, err := FetchLocations()
// 	if err != nil {
// 		log.Fatalf("Failed to fetch locations: %v", err)
// 	}
// 	fmt.Printf("Locations: %+v\n", locations)

// 	dates, err := FetchDates()
// 	if err != nil {
// 		log.Fatalf("Failed to fetch dates: %v", err)
// 	}
// 	fmt.Printf("Dates: %+v\n", dates)

// 	relations, err := FetchRelations()
// 	if err != nil {
// 		log.Fatalf("Failed to fetch relations: %v", err)
// 	}
// 	fmt.Printf("Relations: %+v\n", relations)
// }
