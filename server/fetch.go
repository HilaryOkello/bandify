package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var templates map[string]*template.Template

var artists []Artist

var artistsURL = "https://groupietrackers.herokuapp.com/api/artists"

func init() {
	var err error
	templates, err = loadTemplates()
	if err != nil {
		log.Fatal(err)
	}

	if err := FetchArtists(); err != nil {
		log.Fatal("could not fetch artists: ", err)
	}
}

func loadTemplates() (map[string]*template.Template, error) {
	templates = make(map[string]*template.Template)
	layout := "templates/layout.html"

	pages, err := filepath.Glob("templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to load template files: %w", err)
	}

	for _, page := range pages {
		if page == layout {
			continue
		}
		files := []string{layout, page}
		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, fmt.Errorf("failed to parse template %s: %w", page, err)
		}
		templates[filepath.Base(page)] = tmpl
	}

	return templates, nil
}

func fetchData(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from %s: %w", url, err)
	}

	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body from %s: %w", url, err)
	}

	if err := json.Unmarshal(bytes, target); err != nil {
		return fmt.Errorf("failed to unmarshal data from %s: %w", url, err)
	}

	return nil
}

func FetchArtists() error {
	return fetchData(artistsURL, &artists)
}

func FetchLocations(url string) (Loc, error) {
	var location Loc
	err := fetchData(url, &location)
	return location, err
}

func FetchRelation(url string) (Relation, error) {
	var relation Relation
	err := fetchData(url, &relation)
	return relation, err
}

func FetchDates(url string) (Date, error) {
	var dates Date
	err := fetchData(url, &dates)
	return dates, err
}
