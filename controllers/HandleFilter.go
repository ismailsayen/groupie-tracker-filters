package controllers

import (
	"fmt"
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
		firstAlbum := r.FormValue("firstAlbum")
		numberOfMembers := r.Form["numberOfMembers"]
		locationsOfConcerts := r.FormValue("locationsOfConcerts")
		if minCreationDate >= maxCreationDate {
			e := database.ErrorPage{Status: 400, Type: "The minimum date must always be less than the maximum date. Please go to the homepage and try again."}
			RenderTempalte(w, "templates/error.html", e, http.StatusBadRequest)
			return
		}
		if value, ok := cache.Load("artists"); ok {
			if artistsCache, ok := value.(*[]database.Artists); ok {
				wg.Add(4)
				go filter.GetCreattionDate(&artists, artistsCache, minCreationDate, maxCreationDate, &wg)
				go filter.GetFirstAlbum(&artists, artistsCache, firstAlbum, &wg)
				go filter.NumberOfMembers(&artists, artistsCache, numberOfMembers, &wg)
				go filter.LocationsOfConcert(&LocaFltr, &artists, artistsCache, locationsOfConcerts, &wg)
				wg.Wait()
				fmt.Println(artists)
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
