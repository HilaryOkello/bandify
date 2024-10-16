package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

func TestRenderTemplate(t *testing.T) {
	// Mock templates
	templates = map[string]*template.Template{
		"test.html": template.Must(template.New("layout.html").Parse("{{.Title}}")),
	}

	tests := []struct {
		name     string
		tmpl     string
		data     interface{}
		expected int
	}{
		{"Valid template", "test.html", TemplateData{Title: "Test"}, http.StatusOK},
		{"Invalid template", "nonexistent.html", nil, http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			renderTemplate(w, tt.tmpl, tt.data)
			if w.Code != tt.expected {
				t.Errorf("Expected status code %d, got %d", tt.expected, w.Code)
			}
		})
	}
}

func TestCheckMethodAndPath(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		path           string
		expectedMethod string
		expectedPath   string
		expected       bool
	}{
		{"Correct method and path", "GET", "/test", "GET", "/test", true},
		{"Incorrect method", "POST", "/test", "GET", "/test", false},
		{"Incorrect path", "GET", "/wrong", "GET", "/test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.method, tt.path, nil)
			result := checkMethodAndPath(w, r, tt.expectedMethod, tt.expectedPath)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// func TestMainPage(t *testing.T) {
// 	w := httptest.NewRecorder()
// 	r := httptest.NewRequest("GET", "/", nil)
// 	MainPage(w, r)
// 	if w.Code != http.StatusOK {
// 		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
// 	}
// }

// func TestInfoAboutArtist(t *testing.T) {
// 	// Mock artists slice
// 	artists = []Artist{{ID: 1, Name: "Test Artist"}}

// 	tests := []struct {
// 		name     string
// 		id       string
// 		expected int
// 	}{
// 		{"Valid artist", "1", http.StatusOK},
// 		{"Invalid artist ID", "0", http.StatusNotFound},
// 		{"Non-existent artist", "2", http.StatusNotFound},
// 		{"Invalid ID format", "abc", http.StatusNotFound},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			r := httptest.NewRequest("GET", "/artists/?id="+tt.id, nil)
// 			InfoAboutArtist(w, r)
// 			if w.Code != tt.expected {
// 				t.Errorf("Expected status code %d, got %d", tt.expected, w.Code)
// 			}
// 		})
// 	}
// }

// func TestSearchPage(t *testing.T) {
// 	// Mock artists slice
// 	artists = []Artist{
// 		{ID: 1, Name: "Artist One"},
// 		{ID: 2, Name: "Artist Two"},
// 	}

// 	tests := []struct {
// 		name     string
// 		query    string
// 		expected int
// 	}{
// 		{"Valid search", "Artist", http.StatusOK},
// 		{"Empty query", "", http.StatusBadRequest},
// 		{"No results", "NonexistentArtist", http.StatusOK},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			r := httptest.NewRequest("GET", "/search/?q="+tt.query, nil)
// 			SearchPage(w, r)
// 			if w.Code != tt.expected {
// 				t.Errorf("Expected status code %d, got %d", tt.expected, w.Code)
// 			}
// 		})
// 	}
// }

func TestErrorPage(t *testing.T) {
	tests := []struct {
		name     string
		code     int
		expected string
	}{
		{"Not Found", http.StatusNotFound, "404 - Not Found"},
		{"Bad Request", http.StatusBadRequest, "400 - Bad Request"},
		{"Method Not Allowed", http.StatusMethodNotAllowed, "405 - Method Not Allowed"},
		{"Internal Server Error", http.StatusInternalServerError, "500 - Internal Server Error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ErrorPage(w, tt.code)
			if w.Code != tt.code {
				t.Errorf("Expected status code %d, got %d", tt.code, w.Code)
			}
			if !strings.HasPrefix(w.Body.String(), tt.expected) {
				t.Errorf("Expected body %s, got %s,", tt.expected, w.Body.String())
			}
		})
	}
}
