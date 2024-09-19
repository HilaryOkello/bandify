package server

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