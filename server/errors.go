package server

import "net/http"

func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	var message string

	switch status {
	case http.StatusNotFound:
		message = "404 Not Found"
	case http.StatusBadRequest:
		message = "400 Bad Request"
	case http.StatusMethodNotAllowed:
		message = "405 Method Not Allowed"
	default:
		message = "500 Internal Server Error"
	}

	// Prepare the data to be passed to the template
	data := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}

	// Render the error page using the errors.html template
	if err := templates.ExecuteTemplate(w, "errors.html", data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
