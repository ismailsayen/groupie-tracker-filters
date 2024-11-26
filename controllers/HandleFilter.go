package controllers

import (
	"net/http"
	"sync"

	"groupietracker/controllers/filter"
	"groupietracker/database"
)

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		var wg sync.WaitGroup
		var artists []database.Artists
		minCreationDate := r.FormValue("minCreationDate")
		maxCreationDate := r.FormValue("maxCreationDate")
		firstAlbum1 := r.FormValue("firstAlbum1")
		firstAlbum2 := r.FormValue("firstAlbum2")
		numberOfMembers := r.Form["numberOfMembers"]
		locationsOfConcerts := r.FormValue("locationsOfConcerts")
		if minCreationDate >= maxCreationDate {
			minCreationDate, maxCreationDate = maxCreationDate, minCreationDate
		}
		if value, ok := cache.Load("artists"); ok {
			if artistsCache, ok := value.(*[]database.Artists); ok {
				wg.Add(4)
				go filter.GetCreattionDate(&artists, artistsCache, minCreationDate, maxCreationDate, &wg)
				go filter.GetFirstAlbum(&artists, artistsCache, firstAlbum1, firstAlbum2, &wg)
				go filter.NumberOfMembers(&artists, artistsCache, numberOfMembers, &wg)
				go filter.LocationsOfConcert(&LocaFltr, &artists, artistsCache, locationsOfConcerts, &wg)
				wg.Wait()
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
