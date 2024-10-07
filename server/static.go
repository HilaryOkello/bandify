package server

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	// Remove the /static/ prefix from the URL path
	filePath := path.Join("static", strings.TrimPrefix(r.URL.Path, "/static/"))
	// Check if the file path is a directory
	info, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
		ErrorPage(w, http.StatusNotFound)
		return
	}
	if info.IsDir() {
		ErrorPage(w, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, filePath)
}
