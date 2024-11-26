package filter

import (
	"strconv"
	"sync"

	"groupietracker/database"
)

func NumberOfMembers(a *[]database.Artists, data *[]database.Artists, key []string, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(key) == 0 {
		return
	}
	for _, ele := range *data {
		for _, e := range key {
			nb, _ := strconv.Atoi(e)
			if len(ele.Members) == nb {
				*a = append(*a, ele)
			}
		}
	}
}
