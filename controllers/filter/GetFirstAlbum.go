package filter

import (
	"sync"

	"groupietracker/database"
)

func GetFirstAlbum(a *[]database.Artists, data *[]database.Artists, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(key) == 0 {
		return
	}
	for _, ele := range *data {
		if ele.FirstAlbum == key {
			*a = append(*a, ele)
		}
	}
}
