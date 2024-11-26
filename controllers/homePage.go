package controllers

import (
	"net/http"
	"sync"

	"groupietracker/database"
)

var (
	cache    sync.Map
	LocaFltr database.LocaFltr
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == http.MethodGet {
			var artists []database.Artists

			data := &database.Data{}
			var wg sync.WaitGroup
			err := FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

			artists[20].Image = "assets/img/3ib.jpg"
			data.Art = artists
			wg.Add(4)
			go data.FindMinMax(&artists, &wg)
			go HandleLocations(&LocaFltr, data, &wg)
			go StoreDataOncache(&artists, &wg)
			wg.Wait()

			err = RenderTempalte(w, "./templates/index.html", data, http.StatusOK)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}
		} else {
			e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
			RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
			return
		}
	default:
		e := database.ErrorPage{Status: 404, Type: "Page Not Found"}
		RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
		return
	}
}

func StoreDataOncache(a *[]database.Artists, wg *sync.WaitGroup) {
	defer wg.Done()
	cache.Store("artists", a)
}
