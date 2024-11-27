package controllers

import (
	"net/http"

	"groupietracker/database"
)

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()

		var artists []database.Artists
		minCreationDate := r.FormValue("minCreationDate")
		maxCreationDate := r.FormValue("maxCreationDate")
		firstAlbum1 := r.FormValue("firstAlbum1")
		firstAlbum2 := r.FormValue("firstAlbum2")
		numberOfMembers := r.Form["numberOfMembers"]
		locationsOfConcerts := r.FormValue("locationsOfConcerts")
		if minCreationDate > maxCreationDate {
			minCreationDate, maxCreationDate = maxCreationDate, minCreationDate
		}
		if value, ok := cache.Load("artists"); ok {
			if artistsCache, ok := value.(*[]database.Artists); ok {
				artistsFiltred(artistsCache, &artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts, numberOfMembers)
				RenderTempalte(w, "templates/filter.html", artists, http.StatusInternalServerError)
				return
			} else {
				e := database.ErrorPage{Status: 500, Type: "Error while getting artist data. Please go back to the home page and try again."}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}
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

func artistsFiltred(data *[]database.Artists, a *[]database.Artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts string, numberOfMembers []string) []database.Artists {
	hasDate:=false
	hasFirstAlbum:=false
	hasMembers:=false
	hasLocations:=false
}
