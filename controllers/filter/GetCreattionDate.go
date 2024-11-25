package filter

import (
	"strconv"
	"sync"

	"groupietracker/database"
)

func GetCreattionDate(a *[]database.Artists, data *[]database.Artists, min string, max string, wg *sync.WaitGroup) {
	defer wg.Done()
	minV, _ := strconv.Atoi(min)
	maxV, _ := strconv.Atoi(max)
	for _, ele := range *data {
		if ele.CreationDate >= minV && maxV >= ele.CreationDate {
			*a = append(*a, ele)
		}
	}
}
