package server

import "net/http"

// RelationsHandler handles requests for the relations page.
func RelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := FetchRelations()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Render the relations page with the fetched data
	if err := tmpl.ExecuteTemplate(w, "relations.html", relations); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
	}
}

// FetchRelations retrieves relations data from the API and returns it as a slice of Relation structs.
func FetchRelations() ([]Relation, error) {
	var relations []Relation
	err := fetchData(relationsURL, &relations)
	return relations, err
}
