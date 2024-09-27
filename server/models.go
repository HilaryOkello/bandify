package server

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
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
