package utils

import (
	"groupietracker/database"
)

func LocationsOfConcert(l *database.LocaFltr, a *database.Artists, key string) bool {
	if key == "" {
		return true
	}
	if key == "seattle-usa" {
		key = "washington-usa"
	}
	for _, locations := range l.Index {
		for _, adress := range locations.Locations {
			if adress == key {
				if locations.ID == a.ID {
					return true
				}
			}
		}
	}
	return false
}
