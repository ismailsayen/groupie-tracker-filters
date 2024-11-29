package controllers

import (
	"sync"

	"groupietracker/database"
)

func HandDatafilter(d *database.Data, artistsData *[]database.Artists, l *database.LocaFltr) {
	var wg sync.WaitGroup
	wg.Add(4)
	go d.FindMinMax(artistsData, &wg)
	go HandleLocations(l, d, &wg)
	go d.Albums(artistsData, &wg)
	wg.Wait()
}
