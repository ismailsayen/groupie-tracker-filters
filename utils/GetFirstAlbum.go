package utils

import (
	"strconv"
	"strings"

	"groupietracker/database"
)

func GetFirstAlbum(a *database.Artists, y1, y2 string) bool {
	if len(y1) == 0 && len(y2) == 0 {
		return true
	}
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
