package controllers

import (
	"net/http"

	"groupietracker/database"
	"groupietracker/utils"
)

// HandleFilter fetches data from an API, filters the artists based on user input, and renders the filtered data on the page.
func HandleFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
		RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
		return
	}
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

	artistsData[20].Image = "assets/img/3ib.jpg"
	artistsFiltred(&artistsData, &artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts, numberOfMembers)
	HandDatafilter(data, &artistsData, &LocaFltr)
	data.Art = artists
	RenderTempalte(w, "templates/index.html", data, http.StatusOK)
}

// artistsFiltred filters the artists based on the provided criteria such as creation date, first album, number of members, and concert locations.
func artistsFiltred(data *[]database.Artists, a *[]database.Artists, minCreationDate, maxCreationDate, firstAlbum1, firstAlbum2, locationsOfConcerts string, numberOfMembers []string) {
	hasDate := false
	hasFirstAlbum := false
	hasMembers := false
	hasLocations := false
	for _, artist := range *data {
		hasDate = utils.GetCreattionDate(&artist, minCreationDate, maxCreationDate)
		hasFirstAlbum = utils.GetFirstAlbum(&artist, firstAlbum1, firstAlbum2)
		hasMembers = utils.NumberOfMembers(&artist, numberOfMembers)
		hasLocations = utils.LocationsOfConcert(&LocaFltr, &artist, locationsOfConcerts)
		if hasDate && hasFirstAlbum && hasMembers && hasLocations {
			*a = append(*a, artist)
		}

	}
}
