package filter

import (
	"strconv"
	"strings"

	"groupietracker/database"
)

func GetFirstAlbum(a *database.Artists, firstAlbum1, firstAlbum2 string) bool {
	if len(firstAlbum1) == 0 && len(firstAlbum2) == 0 {
		return true
	}
	y1 := firstAlbum1
	y2 := firstAlbum2
	if y1 > y2 {
		y1, y2 = y2, y1
	}

	minyear, _ := strconv.Atoi(y1)
	maxyear, _ := strconv.Atoi(y2)
	for i := minyear; i <= maxyear; i++ {
		if strings.HasSuffix(a.FirstAlbum, "-"+strconv.Itoa(i)) {
			return true
		}
	}
	return false
}
