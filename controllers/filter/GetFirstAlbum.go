package filter

import (
	"strconv"
	"strings"
	"sync"

	"groupietracker/database"
)

func GetFirstAlbum(a *[]database.Artists, data *[]database.Artists, firstAlbum1, firstAlbum2 string, wg *sync.WaitGroup) {
	defer wg.Done()
	if len(firstAlbum1) == 0 && len(firstAlbum2) == 0 {
		return
	}
	y1, y2 := GetYear(firstAlbum1, firstAlbum2)
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	for _, ele := range *data {
		minyear, _ := strconv.Atoi(y1)
		maxyear, _ := strconv.Atoi(y2)
		for i := minyear; i < maxyear; i++ {
			if ele.FirstAlbum == strconv.Itoa(i) {
				*a = append(*a, ele)
			}
		}
	}
}

func GetYear(year1, year2 string) (string, string) {
	y1 := strings.Split(year1, "-")
	y2 := strings.Split(year2, "-")
	return y1[2], y2[2]
}
