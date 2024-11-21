package controllers

import "net/http"

func HandleFilter(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "ssss", http.StatusNotFound)
}
