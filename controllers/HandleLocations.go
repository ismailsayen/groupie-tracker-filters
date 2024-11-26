package controllers

import (
	"fmt"
	"sync"

	"groupietracker/database"
)

func HandleLocations(c *database.LocaFltr, d *database.Data, wg *sync.WaitGroup) {
	defer wg.Done()
	err := FetchAPI("https://groupietrackers.herokuapp.com/api/locations", &c)
	if err != nil {
		fmt.Println("Error fetching API:", err)
		return
	}
	go d.AllLocations(c, wg)
}
