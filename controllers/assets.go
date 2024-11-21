package controllers

import (
	"net/http"
	"os"

	"groupietracker/database"
)

func AssetsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		file, err := os.Stat(r.URL.Path[1:])
		if err != nil || file.IsDir() {
			e := database.ErrorPage{Status: 404, Type: "Page Not Found"}
			RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
			return
		}

		fs := http.Dir("assets/")
		http.StripPrefix("/assets/", http.FileServer(fs)).ServeHTTP(w, r)
		return

	} else {
		e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
		RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
		return
	}
}
