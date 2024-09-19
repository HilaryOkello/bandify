package server

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

// Home handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if err := templates.ExecuteTemplate(w, "home.html", nil); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}
