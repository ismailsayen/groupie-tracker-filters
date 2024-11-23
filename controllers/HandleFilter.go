package controllers

import (
	"net/http"

	"groupietracker/database"
)

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		var artists []database.Artists
		CreationDate := r.FormValue("creationDate")
		firstAlbum := r.FormValue("firstAlbum")
		numberOfMembers := r.Form["numberOfMembers"]
		locationsOfConcerts := r.FormValue("")
		if value, ok := cache.Load("artists"); ok {
			Filter(value, numberOfMembers, CreationDate, firstAlbum, locationsOfConcerts)
		} else {
			e := database.ErrorPage{Status: 500, Type: "Error while getting artist data. Please go back to the home page and try again."}
			RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
			return
		}
	} else {
		e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
		RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
		return
	}
}

func Filter(a any, ,numberOfMembers []string, CreationDate, firstAlbum, locationsOfConcerts string) {
}
