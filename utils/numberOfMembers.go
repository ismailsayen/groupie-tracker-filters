package utils

import (
	"strconv"

	"groupietracker/database"
)

func NumberOfMembers(a *database.Artists, key []string) bool {
	if len(key) == 0 {
		return true
	}

	for _, e := range key {
		nb, _ := strconv.Atoi(e)
		if len(a.Members) == nb {
			return true
		}
	}
	return false
}
