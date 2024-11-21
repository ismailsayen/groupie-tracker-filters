package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"groupietracker/database"
)

func InfosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id := r.URL.Query().Get("id")

		if id != "" {

			ID, err := strconv.Atoi(id)
			if err != nil {
				e := database.ErrorPage{Status: 404, Type: "User not found"}
				RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
				return
			}

			id = strconv.Itoa(ID)

			var artist database.Artists
			err = FetchAPI("https://groupietrackers.herokuapp.com/api/artists/"+id, &artist)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

			if artist.ID == 0 {
				e := database.ErrorPage{Status: 404, Type: "User not found"}
				RenderTempalte(w, "templates/error.html", e, http.StatusNotFound)
				return
			}

			err = GetForeignData(&artist)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}

			if artist.ID == 21 {
				artist.Image = "./assets/img/3ib.jpg"
			}
			fmt.Println("-------------------")
			fmt.Println(artist.CreationDate)
			fmt.Println(artist.FirstAlbum)
			fmt.Println(len(artist.Members))
			fmt.Println(artist.Loca.Locations)
			fmt.Println("-------------------")
			err = RenderTempalte(w, "./templates/infos.html", artist, http.StatusOK)
			if err != nil {
				e := database.ErrorPage{Status: 500, Type: "Server Error"}
				RenderTempalte(w, "templates/error.html", e, http.StatusInternalServerError)
				return
			}
		} else {
			e := database.ErrorPage{Status: 400, Type: "Bad Request"}
			RenderTempalte(w, "templates/error.html", e, http.StatusBadRequest)
			return
		}
	} else {
		e := database.ErrorPage{Status: 405, Type: "Method Not Allowed"}
		RenderTempalte(w, "templates/error.html", e, http.StatusMethodNotAllowed)
		return
	}
}
