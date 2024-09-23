package server

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

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

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)

	layout := "templates/layout.html"
	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		if page != layout {
			files := []string{layout, page}
			templates[filepath.Base(page)] = template.Must(template.ParseFiles(files...))
		}
	}
}

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
