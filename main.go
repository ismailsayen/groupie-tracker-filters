package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"groupietracker/controllers"
)

const port string = ":8082"

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Please enter only the program name.")
		return
	}

	http.HandleFunc("/assets/", controllers.AssetsHandler)
	http.HandleFunc("/", controllers.IndexHandler)
	http.HandleFunc("/Infos", controllers.InfosHandler)
	http.HandleFunc("/Filter", controllers.HandleFilter)
	fmt.Println("http://localhost" + port + "/")
	log.Fatal(http.ListenAndServe(port, nil))
}
