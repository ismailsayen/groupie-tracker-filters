package controllers

import (
	"sync"

	"groupietracker/database"
)

// manipulate data artists and get min and max dates of creation, get all locations and delete duplicate locations
func HandDatafilter(d *database.Data, artistsData *[]database.Artists, l *database.LocaFltr) {
	var wg sync.WaitGroup
	wg.Add(2)
	go d.FindMinMax(artistsData, &wg)
	go HandleLocations(l, d, &wg)
	wg.Wait()
}
