package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

var artists []Artist

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)

	layout := "templates/layout.html"
	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {
		if page != layout {
			files := []string{layout, page}
			templates[filepath.Base(page)] = template.Must(template.ParseFiles(files...))
		}
	}

	if err := fetchData(artistsURL, &artists); err != nil {
		log.Fatal("Could not fetch artists")
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
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	if err := json.Unmarshal(body, result); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	return nil
}
