package controllers

import (
	"sync"

	"groupietracker/database"
)

// Manipulate artist data to get the min and max creation dates, retrieve all locations, and remove duplicate locations
func HandDatafilter(d *database.Data, artistsData *[]database.Artists, l *database.LocaFltr) {
	var wg sync.WaitGroup
	wg.Add(2)
	go d.FindMinMax(artistsData, &wg)
	go HandleLocations(l, d, &wg)
	wg.Wait()
}
