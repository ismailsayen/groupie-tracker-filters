package filter

import (
	"strconv"

	"groupietracker/database"
)

func GetCreattionDate(a *database.Artists, min string, max string) bool {
	minV, _ := strconv.Atoi(min)
	maxV, _ := strconv.Atoi(max)
	if minV > maxV {
		minV, maxV = maxV, minV
	}
	if (minV == 1987 && maxV == 1987) || (a.CreationDate >= minV && maxV >= a.CreationDate) {
		return true
	}
	return false
}
