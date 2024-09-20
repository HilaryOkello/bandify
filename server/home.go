package server

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

// Initialize templates using ParseGlob
func init() {
    tmpl = template.Must(template.ParseGlob("templates/*.html"))
}


// Home handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data := struct {
        Title string
    }{
        Title: "Groupie Trackers - Home",
    }

	renderTemplate(w, "index.html", data)
}
