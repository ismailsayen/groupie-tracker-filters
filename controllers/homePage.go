package controllers

import (
	"net/http"

	"groupietracker/database"
)

var LocaFltr database.LocaFltr

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		if r.Method == http.MethodGet {
			data := &database.Data{}

			var artists []database.Artists
			err := FetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}
			artists[20].Image = "assets/img/3ib.jpg"
			HandDatafilter(data, &artists, &LocaFltr)
			data.Art = artists

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
