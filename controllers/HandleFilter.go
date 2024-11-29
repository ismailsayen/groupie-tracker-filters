package controllers

import (
	"net/http"

	"groupietracker/controllers/filter"
	"groupietracker/database"
)

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		var artists []database.Artists
		var artistsData []database.Artists
		data := &database.Data{}
		err := FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artistsData)
		if err != nil {
			e := database.ErrorPage{Status: 500, Type: "Server Error"}
			RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
			return
		}

		minCreationDate := r.FormValue("minCreationDate")
		maxCreationDate := r.FormValue("maxCreationDate")
		firstAlbum1 := r.FormValue("firstAlbum1")
		firstAlbum2 := r.FormValue("firstAlbum2")
		numberOfMembers := r.Form["numberOfMembers"]
		locationsOfConcerts := r.FormValue("locationsOfConcerts")

		artistsFiltred(&artistsData, &artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts, numberOfMembers)
		HandDatafilter(data, &artistsData, &LocaFltr)
		artistsData[20].Image = "assets/img/3ib.jpg"
		data.Art = artists
		RenderTempalte(w, "templates/index.html", data, http.StatusOK)
		return

	} else {
		e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
		RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
		return
	}
}

func artistsFiltred(data *[]database.Artists, a *[]database.Artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts string, numberOfMembers []string) {
	hasDate := false
	hasFirstAlbum := false
	hasMembers := false
	hasLocations := false
	for _, artist := range *data {
		hasDate = filter.GetCreattionDate(&artist, minCreationDate, maxCreationDate)
		hasFirstAlbum = filter.GetFirstAlbum(&artist, firstAlbum1, firstAlbum2)
		hasMembers = filter.NumberOfMembers(&artist, numberOfMembers)
		hasLocations = filter.LocationsOfConcert(&LocaFltr, &artist, locationsOfConcerts)
		if hasDate && hasFirstAlbum && hasMembers && hasLocations {
			*a = append(*a, artist)
		}

	}
}
