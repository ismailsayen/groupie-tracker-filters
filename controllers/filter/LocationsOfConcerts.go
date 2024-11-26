package filter

import (
	"sync"

	"groupietracker/database"
)

func LocationsOfConcert(l *database.LocaFltr, a *[]database.Artists, data *[]database.Artists, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, locations := range l.Index {
		for _, adress := range locations.Locations {
			if adress == key {
				for _, ele := range *data {
					if locations.ID == ele.ID {
						*a = append(*a, ele)
					}
				}
			}
		}
	}
}
